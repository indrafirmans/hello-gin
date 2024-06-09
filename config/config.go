package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)
	envFilePath := filepath.Join(cwd, "../../.env")
	log.Printf("Loading .env file from: %s", envFilePath)
	errLoad := godotenv.Load(envFilePath)
	if errLoad != nil {
		log.Fatalf("Error loading .env file")
	}

	value, exists := os.LookupEnv("CONSTANT_USER")
	if !exists {
		log.Fatalf("Environment variable %s not set", "CONSTANT_USER")
	}
	log.Printf("value %s", value)
}
