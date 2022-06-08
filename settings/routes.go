package settings

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("/api")
	user.ApiRouter(api)
}
