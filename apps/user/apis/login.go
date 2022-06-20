package apis

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/serializers"
	"person/core"
	"person/core/errno"
)

func Login(c *fiber.Ctx) error {
	user := new(serializers.LoginFields)
	if err := c.BodyParser(&user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	data, err := serializers.Login(user)
	if err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, data, nil)
}
