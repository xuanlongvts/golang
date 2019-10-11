package driver

import (
	"fmt"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	connStr := fmt.Sprintf("mongodb://localhost:27017")
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect success")
	Mongo.Client = client
	return Mongo
}