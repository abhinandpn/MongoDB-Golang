package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDB initializes a connection to MongoDB and returns a database instance
func ConnectMongoDB() *mongo.Database {
	// Load environment variables from .env file
	fmt.Println("start-----")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve MongoDB URI and Database Name from environment variables
	uri := os.Getenv("DATABASE_URI")
	dbName := os.Getenv("DATABASE_NAME")

	// Check if URI is empty
	if uri == "" || dbName == "" {
		log.Fatal("Missing MongoDB connection details in .env file")
	}

	// Setup MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	// Establish a connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")
	return client.Database(dbName)
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return ConnectMongoDB().Collection(collectionName)
}
