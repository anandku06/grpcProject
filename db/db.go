package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollections *mongo.Collection

func InitMango() error {
	// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
	// Canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // if connection not happens, then after 10 seconds
	defer cancel() // this is called if connection does'nt happen
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	MONGO_URL := os.Getenv("MONGODB_URI")
	// Set client options

	clientOptions := options.Client().ApplyURI(MONGO_URL) // ApplyURI sets the URI to connect to the MongoDB server
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println("Error in MongoDB Connection!", err)
		return err
	}

	UserCollections = client.Database("grpc").Collection("users")
	// Collection gets a handle for a collection with the given name configured with the given CollectionOptions.
	// Database returns a handle for a database with the given name configured with the given DatabaseOptions.

	return nil
}
