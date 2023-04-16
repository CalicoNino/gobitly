package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")
	ip := os.Getenv("MONGO_DB_IP")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?retryWrites=true&w=majority", username, password, ip)
	return uri
}
