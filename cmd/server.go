package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var timeout = 5 * time.Second

func main() {
	port, ok := os.LookupEnv("MONGODB_PORT_NUMBER")
	if !ok {
		panic(errors.New("port environment not set"))
	}

	if err := pingUri(fmt.Sprintf("mongodb://localhost:%v", port), timeout); err != nil {
		log.Printf("failed to connect to mongo: %+v", err)
		panic(errors.New("failed to connect to mongo"))
	}
}

func pingUri(uri string, timeout time.Duration) error {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(uri),
		&options.ClientOptions{
			Timeout: &timeout,
		})
	if err != nil {
		log.Panicf("failed to create client: %v", err)
	}

	db := client.Database("admin")
	return db.RunCommand(context.Background(), bson.M{"ping": 1}).Err()
}
