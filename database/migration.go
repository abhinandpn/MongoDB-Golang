package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MigrateDatabase ensures collections and indexes exist
func MigrateDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get Collections
	userCollection := GetCollection("users")
	postCollection := GetCollection("posts")

	// Create Unique Index on Email in Users Collection
	_, err := userCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // 1 means ascending order
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Println("⚠️  Could not create index on users collection: ", err)
	}

	// Ensure posts collection exists (Index on Title for fast search)
	_, err = postCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"title": 1},
	})
	if err != nil {
		log.Println("⚠️  Could not create index on posts collection: ", err)
	}

	fmt.Println("✅ Database migration completed successfully!")
}
