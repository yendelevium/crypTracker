package routes

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yendelevium/crypTracker/internal/api"
)

// TODO:
// Update all errors to be of type models.Error instead of just sending a message
func CoinRouter() *fiber.App {
	// Route handlers have the follwing function signature
	// func (*fiber.Ctx) error {}
	coinRouter := fiber.New()
	coinRouter.Get("/test", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Testing mount")
	})

	coinRouter.Get("/coins", func(c *fiber.Ctx) error {
		// Change this to fetch the data from the DB
		coinData, err := api.FetchCoinData()
		if err != nil {
			log.Fatalf("Error fetching coin data: %s", err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(nil)
		}
		c.Status(http.StatusOK)
		return c.JSON(coinData)
	})

	coinRouter.Get("/coins/:coinId", func(c *fiber.Ctx) error {
		// Logic for single coin details
		// The stuff we display in the /coins, + Short desc, graph, etc
		// DB requried
		c.Status(http.StatusOK)
		return c.JSON("WIP")
	})

	return coinRouter
}
