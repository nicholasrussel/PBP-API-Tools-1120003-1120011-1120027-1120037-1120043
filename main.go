package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Hello World")
	//TestGomail()
	TestGoRedis()
}

func LoadEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
