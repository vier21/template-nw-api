package domain

type JWTPayload struct {
	Username  string   `json:"username,omitempty"`
	Email     string   `json:"email,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	ExpiredAt int64    `json:"exp,omitempty"`
}
