package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Database {
	mongo_uri := os.Getenv("MONGO_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_uri).SetServerAPIOptions(serverAPI)

	// create new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)

	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	farmforumDatabase := client.Database(os.Getenv("FARM_FORUM_DB_NAME"))
	return farmforumDatabase
}
