package apis

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"person/apps/user/models"
	"person/core"
	"person/core/errno"
)

func Query(c *fiber.Ctx) error {
	type User struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile"`
	}

	user := new(User)
	if err := c.QueryParser(user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	fmt.Println(user)

	return core.Response(c, nil, nil)
}

func Create(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}
	fmt.Println(user)
	return core.Response(c, nil, nil)
}

func Update(c *fiber.Ctx) error { // default: PUT
	id, err := c.ParamsInt("id")
	if err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	fmt.Println(id)

	if c.Method() == "PATCH" {
		fmt.Println("PATCH")
	}

	return core.Response(c, nil, nil)
}

func Destroy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return core.Response(c, nil, errno.ErrBind)
	}

	fmt.Println(id)
	return core.Response(c, nil, nil)
}
