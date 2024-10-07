package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetAll(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.User), args.Error(1)
}
func (m *UserRepositoryMock) GetOneByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	usr := args.Get(0).(*domain.User)
	return usr , args.Error(1)
}
func (m *UserRepositoryMock) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *UserRepositoryMock) DeleteUser(ctx context.Context, ids ...primitive.ObjectID) ([]string, error) {
	args := m.Called(ctx, ids)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]string), args.Error(1)
}

func (m *UserRepositoryMock) GetLog() *logrus.Logger {
	return logrus.New()
}
