package service

import (
	"context"

	"github.com/vier21/pc-01-network-be/pkg/user/domain"
)

type IService interface {
	Register(context.Context, domain.NewUser) (*DataUserAuthenticated, error)
	Login(context.Context, LoginRequest) (*SecurityAuthenticatedUser, error)
}
