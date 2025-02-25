package main

import (
	"fmt"

	"github.com/abhinandpn/MongoDB-Golang/database"
)

func main() {
	// Connect to MongoDB
	database.ConnectMongoDB()

	// Run Database Migration
	database.MigrateDatabase()

	fmt.Println("🚀 Server is ready!")
}

