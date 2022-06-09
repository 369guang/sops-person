package user

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/apis"
)

func ApiRouter(app fiber.Router) {
	// User
	app.Get("/user", apis.Query)
	app.Post("/user", apis.Create)
	app.Put("/user/:id", apis.Update)
	app.Patch("/user/:id", apis.Update)
	app.Delete("/user/:id", apis.Destroy)
}
