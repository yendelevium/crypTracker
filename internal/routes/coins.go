package routes

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yendelevium/crypTracker/internal/database"
	"github.com/yendelevium/crypTracker/models"
)

// TODO:
// Update all errors to be of type models.Error instead of just sending a message
func CoinRouter(dbClient *database.DBClient) *fiber.App {
	// Route handlers have the follwing function signature
	// func (*fiber.Ctx) error {}
	coinRouter := fiber.New()
	coinRouter.Get("/test", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Testing mount")
	})

	coinRouter.Get("/coins", func(c *fiber.Ctx) error {
		// Change this to fetch the data from the DB
		allCoins := []models.Coin{}
		result := dbClient.Client.Find(&allCoins)
		if result.Error != nil {
			log.Printf("Error fetching coin data : %s", result.Error)
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: result.Error.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		log.Printf("Fetched %d records", result.RowsAffected)
		c.Status(http.StatusOK)
		return c.JSON(allCoins)
	})

	coinRouter.Get("/coins/:coinId", func(c *fiber.Ctx) error {
		// Logic for single coin details
		// The stuff we display in the /coins, + Short desc, graph, etc
		var coin models.Coin
		// SELECT * FROM "coins" WHERE coin_gecko_id = 'coinId' ORDER BY "coins"."coin_gecko_id" LIMIT 1
		result := dbClient.Client.First(&coin, "coin_gecko_id = ?", c.Params("coinId"))
		if result.Error != nil {
			log.Printf("Couldn't find coin :%s", result.Error)
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: "CoinID doesn't exist",
				Status:  http.StatusBadRequest,
			})
		}
		c.Status(http.StatusOK)
		return c.JSON(coin)
	})

	return coinRouter
}
