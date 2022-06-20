package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"person/core"
	"person/core/errno"
)

func AuthToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(core.VIPER.GetString("jwt.secret_key")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return core.Response(c, nil, errno.ErrTokenInvalid)
}
