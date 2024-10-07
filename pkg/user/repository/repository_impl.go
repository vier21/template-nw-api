package repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/config"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db     *mongo.Database
	logger *logrus.Logger
}

func NewUserRepository(cli *mongo.Client, logger *logrus.Logger) IRepository {
	db := cli.Database(config.GetConfig().MongoDBMain)
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var results []domain.User

	meth := "GetAll"
	cur, err := repo.userCollection().Find(ctx, bson.D{{}})

	if err != nil {
		repo.repoLogFunc(meth).Errorf("Error when get data from db: %s", err.Error())
		return nil, err
	}

	if err = cur.All(ctx, &results); err != nil {
		repo.repoLogFunc(meth).Errorf("Error when bind data to result: %s", err.Error())
		return nil, err
	}

	repo.repoLogFunc(meth).Info("success fetch data from db")

	return results, err
}

func (repo *userRepository) GetOneByUsername(ctx context.Context, username string) (*domain.User, error) {
	var result *domain.User

	meth := "GetOneByUsername"

	curr := repo.userCollection().FindOne(ctx, bson.D{{
		Key:   "username",
		Value: username,
	}})

	if err := curr.Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			repo.repoLogFunc(meth).Error("document not found")
			return nil, err
		}
		repo.repoLogFunc(meth).Errorf("Error to retrieve data from DB: %s", err.Error())
		return nil, err

	}

	repo.repoLogFunc(meth).Info("Successfully retrieve data from db")

	return result, nil
}

func (repo *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	meth := "CreateUser"

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := repo.userCollection().InsertOne(ctx, *user)

	if err != nil {
		repo.repoLogFunc(meth).Errorf("Error when inserting data: %s", err.Error())
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (repo *userRepository) DeleteUser(ctx context.Context, id ...primitive.ObjectID) ([]string, error) {
	meth := "DeleteUser"

	result, err := repo.userCollection().DeleteMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: id}}}})
	if err != nil {
		repo.repoLogFunc(meth).Errorf("Error when delete record: %s", err.Error())
		return nil, err
	}

	if result.DeletedCount < 1 {
		repo.repoLogFunc(meth).Errorf("No record ")
		return nil, err
	}

	arrd := make([]string, len(id))
	for i, objectId := range id {
		arrd[i] = objectId.Hex() // Convert ObjectID to its hex string representation
	}

	return arrd, nil
}

func (repo *userRepository) repoLogFunc(meth string) *logrus.Entry {
	return repo.logger.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": meth,
	})
}

func (repo *userRepository) userCollection() *mongo.Collection {
	return repo.db.Collection("users")
}

func (repo *userRepository) GetLog() *logrus.Logger {
	return repo.logger
}
