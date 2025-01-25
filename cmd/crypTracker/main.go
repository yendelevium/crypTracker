package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yendelevium/crypTracker/internal/routes"
)

func main() {
	fmt.Print("Whatup crypTracker?!")

	// Creating a fiber App
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Whatup crypTracker?!")
	})

	// Mounting a sub-router, which has the paths for /crypto
	// I want a path for /crypto/coins for now, but we can expand it later with /crytpo/nfts
	app.Mount("/crypto", routes.CoinRouter())
	app.Mount("/users", routes.UserRouter())

	log.Fatal(app.Listen(":8080"))
}
