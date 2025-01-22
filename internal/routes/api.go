package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Api() *fiber.App {
	// Route handlers have the follwing function signature
	// func (*fiber.Ctx) error {}
	api := fiber.New()
	api.Get("/mount", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Testing mount")
	})
	return api
}
