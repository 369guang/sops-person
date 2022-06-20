package auth

import (
	"github.com/golang-jwt/jwt"
	"person/core"
	"time"
)

type Context struct {
	ID       int
	Username string
}

// Sign 签名
func Sign(c Context, secret string) (string, error) {

	if secret == "" {
		secret = core.VIPER.GetString("jwt.secret_key")
	}
	thirty, _ := time.ParseDuration(core.VIPER.GetString("jwt.expiration_time"))
	expTime := time.Now().Add(thirty)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"iss":      "sops.cloud",
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"exp":      expTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
