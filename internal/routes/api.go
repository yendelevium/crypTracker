package routes

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yendelevium/crypTracker/internal/api"
)

func Api() *fiber.App {
	// Route handlers have the follwing function signature
	// func (*fiber.Ctx) error {}
	apiRouter := fiber.New()
	apiRouter.Get("/test", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Testing mount")
	})

	apiRouter.Get("/allcoins", func(c *fiber.Ctx) error {
		coinData, err := api.FetchCoinData()
		if err != nil {
			log.Fatalf("Error fetching coin data: %s", err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(nil)
		}
		c.Status(http.StatusOK)
		return c.JSON(coinData)
	})
	return apiRouter
}
