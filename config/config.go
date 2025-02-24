package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("Environment variables loaded successfully")
}
