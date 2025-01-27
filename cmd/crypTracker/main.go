package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/yendelevium/crypTracker/internal/database"
	"github.com/yendelevium/crypTracker/internal/routes"
	"github.com/yendelevium/crypTracker/models"
)

func main() {
	fmt.Print("Whatup crypTracker?!")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file :%s", err)
	}
	dbURL := os.Getenv("DB_URL")
	dbClient, err := database.ConnectPostgres(dbURL)
	if err != nil {
		log.Fatalf("Connection to postgres failed :%s", err)
	}

	// Migration
	// The purple lines you see in the console are NOT ERRORS, but migration logs
	// You can turn them off in database.ConnectPostgres()
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Silent), // Set to Silent to suppress query logs
	// })
	err = dbClient.Client.AutoMigrate(&models.User{}, &models.Coin{}, &models.Watchlist{})
	if err != nil {
		log.Fatalf("Couldn't perform DB migrations : %s", err)
	}

	// UNCOMMENT WHEN YOU NEED TO SEED THE Database
	// Find a better way to do this, maybe commandline arguments?!
	// Ig even if u uncomment, it won't reseed as primary keys are same
	// err = dbClient.Seed()
	// if err != nil {
	// 	log.Fatalf("Couldn't seed Database: %s", err)
	// }

	// Creating a fiber App
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Whatup crypTracker?!")
	})

	// Mounting a sub-router, which has the paths for /crypto
	// I want a path for /crypto/coins for now, but we can expand it later with /crytpo/nfts
	app.Mount("/crypto", routes.CoinRouter(dbClient))
	app.Mount("/users", routes.UserRouter(dbClient))

	log.Fatal(app.Listen(":8080"))
}
