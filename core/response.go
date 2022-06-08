package core

import (
	"github.com/gofiber/fiber/v2"
	"person/core/errno"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(c *fiber.Ctx, data interface{}, err error) error {
	code, msg := errno.DecodeErr(err)
	if code != 0 {
		LOGGER.Error(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
