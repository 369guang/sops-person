package auth

import (
	"errors"
)

var ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")

//func ParseRequest(c *gin.Context) (*Context, error) {
//	header := c.Request.Header.Get("Authorization")
//	secret := core.VIPER.GetString("jwt.secret_key")
//
//	if len(header) == 0 {
//		return &Context{}, ErrMissingHeader
//	}
//
//	var t string
//
//	fmt.Sscanf(header, "BR %s", &t)
//	return Parse(t, secret)
//}
