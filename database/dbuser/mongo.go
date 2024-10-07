package dbuser

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoConnection(logger *logrus.Logger) (*mongo.Client, context.CancelFunc, context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	optCred := options.Credential{
		Username: config.GetConfig().MongoDBUsername,
		Password: config.GetConfig().MongoDBPassword,
	}
	opts := []*options.ClientOptions{
		options.Client().ApplyURI(config.GetConfig().MongoDBURL),
		options.Client().SetAuth(optCred),
		options.Client().SetMaxPoolSize(20),
	}

	client, err := mongo.Connect(ctx, opts...)
	if err != nil {
		logger.Fatalf("mongodb connection error: %s\n", err)
		os.Exit(1)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatalf("mongodb ping issue: %s \n", err)
		os.Exit(1)
	}

	logger.Infof("Successfully connected to database: %v", config.GetConfig().MongoDBURL)

	return client, cancel, ctx, err
}

func CloseConnection(ctx context.Context, logger *logrus.Logger, client *mongo.Client, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(ctx); err != nil {
			logger.Fatalf("cannot disconnect to db: %s", err)
		}
		logger.Infof("Successfully Disconnected to DB: %s", config.GetConfig().MongoDBURL)
	}()
}
