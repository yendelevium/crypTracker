package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// To load .env file
func LoadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Couldn't load .env file: %s", err)
	}
}
