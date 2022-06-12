package apis

import (
	"github.com/gofiber/fiber/v2"
	"person/apps/user/models"
	"person/apps/user/serializers"
	"person/core"
	"person/core/errno"
)

func List(c *fiber.Ctx) error {

	user := new(serializers.User)
	if err := c.QueryParser(user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	data, total, err := serializers.List(user)
	if err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, core.ListRequest{
		Data:  data,
		Total: total,
	}, nil)
}

func Retrieve(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	data, err := serializers.Retrieve(id)
	if err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, data, nil)
}

func Create(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	if err := serializers.Create(user); err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, nil, nil)
}

func Update(c *fiber.Ctx) error { // default: PUT
	id, err := c.ParamsInt("id")
	if err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	data, err := serializers.Update(id, user, c.Method())
	if err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, data, nil)
}

func Destroy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	if err := serializers.Destroy(id); err != nil {
		return core.Response(c, nil, err)
	}

	return core.Response(c, nil, nil)
}
