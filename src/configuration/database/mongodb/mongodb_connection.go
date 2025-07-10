package mongodb

import (
	"context"
	"os"

	"github.com/tyerry/crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_CRUDINIT = "MONGODB_CRUDINIT"
)

func NewMongoDbConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	database_name := os.Getenv(MONGODB_CRUDINIT)

	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}
	
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Mongodb database connected successfully")

	return client.Database(database_name), nil
}