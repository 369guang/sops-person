package apis

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/serializers"
	"person/core"
	"person/core/auth"
)

func AuthToTp(c *fiber.Ctx) error {
	ctx := auth.ParseRequest(c)
	data, err := serializers.Auth(ctx)
	if err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, data, nil)
}
