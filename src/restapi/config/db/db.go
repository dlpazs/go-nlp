package db

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)



var client *mongo.Client

func GetDBCollection() (*mongo.Collection, error) {
	URI, exists := os.LookupEnv("MONGODB")
	if exists {
		fmt.Println("Mongo Connection string exists ")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")
	collection := client.Database("userdb").Collection("users")
	return collection, nil
}