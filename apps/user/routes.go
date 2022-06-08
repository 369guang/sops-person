package user

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/apis"
)

func ApiRouter(app fiber.Router) {
	// 获取用户信息
	app.Post("/user", apis.Create)
}
