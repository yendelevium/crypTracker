package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/yendelevium/crypTracker/internal/database"
	"github.com/yendelevium/crypTracker/internal/initializers"
	"github.com/yendelevium/crypTracker/internal/routes"
	"github.com/yendelevium/crypTracker/internal/websockets"
	"github.com/yendelevium/crypTracker/models"
)

// This runs BEFORE main
func init() {
	initializers.LoadEnv()
}

func main() {
	fmt.Print("Whatup crypTracker?!")
	dbClient, err := database.ConnectPostgres()
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

	// Start to fetch API data every 10 seconds
	go dbClient.StartScraping()

	// Creating a fiber App
	app := fiber.New()
	// WS upgrade
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws", websockets.WSRouter())

	// Mounting a sub-router, which has the paths for /crypto
	// I want a path for /crypto/coins for now, but we can expand it later with /crytpo/nfts
	app.Mount("/crypto", routes.CoinRouter(dbClient))
	app.Mount("/users", routes.UserRouter(dbClient))

	// PRODUCTION
	if os.Getenv("RUNTIME_ENV") == "production" {
		log.Println("Serving PROD")
		pd, _ := os.Getwd()
		fmt.Println("Current working directory:", pd)

		// Ensure /assets/ paths are correctly served, ie the .js and .css static files in vite
		// If you don't add this, the server is misconfigured and serves index.html instead of actual JS files.
		// Otherwise, you get an error like this:
		// Loading module from “http://localhost:8080/assets/index-B8OBafhz.js” was blocked because of a disallowed MIME type (“text/html”).
		app.Use(func(c *fiber.Ctx) error {
			path := c.Path()
			// Checking if it has /assets/ in it's path, (which is true for .js and .css files)
			// If it is, serve THOSE files instead of sending index.html (default for app.Static())
			if strings.HasPrefix(path, "/assets/") {
				filePath := "./web/dist" + path
				if _, err := os.Stat(filePath); err == nil {
					return c.SendFile(filePath)
				}
				return c.SendStatus(404) // Return 404 if the file doesn't exist
			}
			return c.SendFile("./web/dist/index.html") // Fallback for React Router
		})

		// Serving the static files, which by default, is index.html
		// But you can configure it to be wtv u want using fiber.Static{} config
		app.Static("/", "./web/dist")
	}
	log.Fatal(app.Listen(":8080"))
}
