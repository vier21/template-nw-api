package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRepository interface {
	GetAll(context.Context) ([]domain.User, error)
	GetOneByUsername(context.Context, string) (*domain.User, error)
	CreateUser(context.Context, *domain.User) (*domain.User, error)
	DeleteUser(context.Context, ...primitive.ObjectID) ([]string, error)

	GetLog() *logrus.Logger
}
