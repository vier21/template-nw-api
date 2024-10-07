package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"github.com/vier21/pc-01-network-be/pkg/user/service"
)

type handler struct {
	svc service.IService
}

func writeHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

}

func NewHTTPHandler(svc service.IService) *http.ServeMux {
	handlers := &handler{
		svc: svc,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)
	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)
	return mux
}

func methodValidator(method string, w http.ResponseWriter, r *http.Request) error {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New("wrong method")
	}
	return nil
}

type GeneralResponse struct {
	Data any `json:"data"`
}

func resEncoder(w http.ResponseWriter, data any) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	writeHeader(w)
	if err := methodValidator(http.MethodPost, w, r); err != nil {
		return
	}

	var body domain.NewUser

	json.NewDecoder(r.Body).Decode(&body)

	user, err := h.svc.Register(r.Context(), body)
	if err != nil {
		resEncoder(w, GeneralResponse{
			Data: nil,
		})
	}

	resEncoder(w, GeneralResponse{
		Data: user,
	})

}
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	writeHeader(w)
	if err := methodValidator(http.MethodPost, w, r); err != nil {
		return
	}

	var body service.LoginRequest

	json.NewDecoder(r.Body).Decode(&body)

	user, err := h.svc.Login(r.Context(), body)
	if err != nil {
		resEncoder(w, GeneralResponse{
			Data: nil,
		})
	}

	resEncoder(w, GeneralResponse{
		Data: user,
	})

}

func (h *handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	if err := methodValidator(http.MethodGet, w, r); err != nil {
		return
	}

	writeHeader(w)
	resEncoder(w, GeneralResponse{
		Data: "hello",
	})
}
