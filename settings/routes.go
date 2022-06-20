package settings

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user"
	"person/core/middleware"
)

func LoadRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	user.AuthRouter(auth)

	api := app.Group("/api", middleware.AuthToken())
	user.ApiRouter(api)
}
