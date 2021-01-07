package services

import (
	"context"
	"fmt"
	"github.com/iambenkay/nairacheck/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitializeDatabaseConnection(uri string) {
	utils.Contextualize(func(ctx context.Context) {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		Bean.DatabaseClient = client

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(err)
		}
		fmt.Println("Database connection was successful")
	})
}
