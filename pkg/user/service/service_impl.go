package service

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"github.com/vier21/pc-01-network-be/pkg/user/repository"
	"github.com/vier21/pc-01-network-be/utils"
)

type userService struct {
	repo       repository.IRepository
	tokenMaker IToken
}

func NewUserService(repo repository.IRepository, tmaker IToken) IService {
	return &userService{
		repo:       repo,
		tokenMaker: tmaker,
	}
}

func (s *userService) Register(ctx context.Context, nwuser domain.NewUser) (*DataUserAuthenticated, error) {
	user := toUserFromNewUser(nwuser)
	meth := "Register"

	users, _ := s.repo.GetOneByUsername(ctx, user.Username)

	if users != nil {
		s.repoSvcLog(meth).Error("user already exist")
		return nil, errors.New("user exist")
	}

	if nwuser.Password != utils.GenerateHashPassword(nwuser.Password) {
		user.HashPassword = utils.GenerateHashPassword(user.HashPassword)
	}

	insertedUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		s.repoSvcLog(meth).Errorf("Error Creating User: %s", err.Error())
		return nil, err
	}
	return setAuthenticatedUser(insertedUser), nil
}
func (s *userService) Login(ctx context.Context, req LoginRequest) (*SecurityAuthenticatedUser, error) {
	meth := "Login"
	checkUser, err := s.repo.GetOneByUsername(ctx, req.Username)

	if err != nil {
		s.repoSvcLog(meth).Errorf("error while find document: %s", err.Error())
		return nil, err
	}

	if err := utils.CompareHashPassword(checkUser.HashPassword, req.Password); err != nil {
		s.repoSvcLog(meth).Error("Wrong Password")
		return nil, err
	}

	ss, token, err := s.tokenMaker.GenerateToken(ctx, *checkUser)
	if err != nil {
		s.repoSvcLog(meth).Errorf("Error when generate token: %s", err.Error())
		return nil, err
	}

	exp, err := token.Claims.GetExpirationTime()

	if err != nil {
		s.repoSvcLog(meth).Errorf("Error set expiration time: %s", err.Error())
		return nil, err
	}

	return setToDataSecAuth(*setAuthenticatedUser(checkUser), DataSecurityAuthenticated{
		JWTAccessToken:            ss,
		JWTRefreshToken:           ss,
		ExpirationAccessDateTime:  exp.Time,
		ExpirationRefreshDateTime: exp.Time,
	}), nil
}

func (s *userService) repoSvcLog(meth string) *logrus.Entry {

	return s.repo.GetLog().WithFields(logrus.Fields{
		"layer":  "service",
		"method": meth,
	})
}
