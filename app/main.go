package main

import (
	"gobitly/datastore"
	"gobitly/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	ip := os.Getenv("POSTGRES_IP")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	datastore.Setup(ip, username, password, dbname, port, sslmode)
	server.SetupAndListen()
}
