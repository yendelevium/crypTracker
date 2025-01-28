package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// To load .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file: %s", err)
	}
}
