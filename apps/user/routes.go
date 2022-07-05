package user

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/apis"
)

func ApiRouter(app fiber.Router) {
	// User
	app.Get("/user/", apis.List)
	app.Get("/user/:id", apis.Retrieve)
	app.Post("/user/", apis.Create)
	app.Put("/user/:id", apis.Update)
	app.Patch("/user/:id", apis.Update)
	app.Delete("/user/:id", apis.Destroy)

	// user info
	app.Get("/user/info/", apis.Info)
	// user auth
	app.Get("/user/authtotp/", apis.AuthToTp)
	// settings
	app.Post("/user/settings/", apis.UpdateSettings)

}

func AuthRouter(app fiber.Router) {
	app.Post("/login", apis.Login)
}
