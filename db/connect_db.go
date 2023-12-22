package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGO_URI = "mongodb+srv://gaurav:gaurav123@cluster0.ffhnw.mongodb.net/test1?retryWrites=true&w=majority"

func InitDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))

	if err != nil {
		panic(err)
	}
	return client
}
