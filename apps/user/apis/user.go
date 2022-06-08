package apis

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/models"
	"person/core"
	"person/core/errno"
)

func Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	return core.Response(c, nil, nil)
}
