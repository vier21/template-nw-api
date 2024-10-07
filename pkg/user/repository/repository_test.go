package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vier21/pc-01-network-be/config"
	"github.com/vier21/pc-01-network-be/database/dbuser"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	config.InitConfig(".")
}

func createUser(repo IRepository) (*domain.User, error) {

	user := &domain.User{
		Username:     "Abijat",
		Email:        "abijat@gmail.com",
		FirstName:    "abijat",
		LastName:     "sigh",
		Role:         "admin",
		HashPassword: "sdad123",
	}

	result, err := repo.CreateUser(context.TODO(), user)

	return result, err
}

func testRepo() (IRepository, *logrus.Logger, *mongo.Client, context.CancelFunc) {

	logger := logrus.New()
	cli, cancel, _, err := dbuser.NewMongoConnection(logger)
	if err != nil {
		logger.Error("test error")
		return nil, logger, cli, cancel
	}
	repo := NewUserRepository(cli, logger)

	return repo, logger, cli, cancel
}

func TestCreateUser(t *testing.T) {
	repo, logger, cli, cancel := testRepo()
	defer dbuser.CloseConnection(context.TODO(), logger, cli, cancel)
	user := &domain.User{
		Username:     "Abijat",
		Email:        "abijat@gmail.com",
		FirstName:    "abijat",
		LastName:     "sigh",
		Role:         "admin",
		HashPassword: "sdad123",
	}

	result, err := repo.CreateUser(context.TODO(), user)
	fmt.Println(result)
	assert.NotNil(t, result)
	assert.Equal(t, nil, err)

	del, err := repo.DeleteUser(context.TODO(), result.Id)
	assert.NotNil(t, del)
	assert.Equal(t, nil, err)

}

func TestGetOneByUsername(t *testing.T) {
	repo, logger, cli, cancel := testRepo()
	defer dbuser.CloseConnection(context.TODO(), logger, cli, cancel)

	user1, err := createUser(repo)
	assert.NotNil(t, user1)
	assert.Nil(t, err)

	user, err := repo.GetOneByUsername(context.TODO(), user1.Username)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	_, err = repo.DeleteUser(context.TODO(), user1.Id)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	repo, logger, cli, cancel := testRepo()
	defer dbuser.CloseConnection(context.TODO(), logger, cli, cancel)

	user1, err := createUser(repo)
	assert.NotNil(t, user1)
	assert.Nil(t, err)
	user2, err := createUser(repo)

	assert.NotNil(t, user2)
	assert.Nil(t, err)

	users, err := repo.GetAll(context.TODO())

	assert.NotNil(t, users)
	assert.Nil(t, err)

	del, err := repo.DeleteUser(context.TODO(), user1.Id, user2.Id)

	assert.NotNil(t, del)
	assert.Nil(t, err)

}
