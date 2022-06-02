package settings

import (
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}