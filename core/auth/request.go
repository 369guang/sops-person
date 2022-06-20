package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

//var ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")

func ParseRequest(c *fiber.Ctx) *Context {
	ctx := &Context{}
	token := c.Locals("user").(*jwt.Token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = int(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx
	}
	return ctx
}
