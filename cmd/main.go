package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/config"
	"github.com/vier21/pc-01-network-be/config/keys"
	"github.com/vier21/pc-01-network-be/database/dbuser"
	"github.com/vier21/pc-01-network-be/pkg/user/handler"
	"github.com/vier21/pc-01-network-be/pkg/user/repository"
	"github.com/vier21/pc-01-network-be/pkg/user/service"
)

func main() {
	config.InitConfig(".")
	logger := logrus.New()
	cli, cancel, ctx, err := dbuser.NewMongoConnection(logger)
	defer dbuser.CloseConnection(ctx, logger, cli, cancel)

	if err != nil {
		logger.Errorf("error:%s", err)
		return
	}

	repo := repository.NewUserRepository(cli, logger)
	token := service.NewTokenRSA(keys.LoadPrivateKey(), keys.LoadPublicKey())
	svc := service.NewUserService(repo, token)
	mux := handler.NewHTTPHandler(svc)

	server := &http.Server{
		Handler: mux,
		Addr:    ":3030",
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Errorf("Error: %s", err.Error())
	}
}
