package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vier21/pc-01-network-be/config"
	"github.com/vier21/pc-01-network-be/config/keys"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"github.com/vier21/pc-01-network-be/pkg/user/repository"
	"github.com/vier21/pc-01-network-be/utils"
)

func init() {
	config.InitConfig(".")
}

func TestRegister(t *testing.T) {
	repomock := &repository.UserRepositoryMock{}
	newService := NewUserService(repomock, NewTokenRSA(keys.LoadPrivateKey(), keys.LoadPublicKey()))
	userr := &domain.User{
		Username:     "test",
		Email:        "test",
		FirstName:    "test",
		LastName:     "test",
		Role:         "test",
		HashPassword: "test",
	}

	newUser := &domain.NewUser{
		Username:  "test",
		Email:     "test",
		FirstName: "test",
		LastName:  "test",
		Role:      "test",
		Password:  "test",
	}
	repomock.On("GetOneByUsername", context.TODO(), userr.Username).Return(nil, errors.New("err"))
	repomock.On("CreateUser", context.TODO(), userr).Return(userr, nil)

	usr, err := newService.Register(context.TODO(), *newUser)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestLogin(t *testing.T) {
	req := &LoginRequest{
		Username: "lorem",
		Password: "ipsum",
	}

	user := &domain.User{
		Username:     "lorem",
		Email:        "test@gmail.com",
		FirstName:    "test",
		LastName:     "test",
		Role:         "test",
		HashPassword: utils.GenerateHashPassword(req.Password),
	}
	repoMock := &repository.UserRepositoryMock{}
	svc := NewUserService(repoMock, NewTokenRSA(keys.LoadPrivateKey(), keys.LoadPublicKey()))
	repoMock.On("GetOneByUsername", context.TODO(), req.Username).Return(user, nil)

	res, err := svc.Login(context.TODO(), *req)
	fmt.Println(res)

	assert.NotNil(t, res)
	assert.Nil(t, err)

}
