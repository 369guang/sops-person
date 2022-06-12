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

// Sign HMAC签名
func Sign(c Context, secret string) (string, error) {

	if secret == "" {
		secret = core.VIPER.GetString("jwt.secret_key")
	}

	thirty, _ := time.ParseDuration(core.VIPER.GetString("jwt.expiration_time"))
	expTime := time.Now().Add(thirty)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"iss":      "52bug.me",
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"exp":      expTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func secreFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// Parse 解开签名
func Parse(tokenString, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secreFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = int(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}
