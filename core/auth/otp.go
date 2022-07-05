package auth

import (
	"github.com/pquerna/otp/totp"
)

// CreateKey 创建密钥
// return key, url， err
func CreateKey(username string) (string, string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "sops.cloud",
		AccountName: username,
	})
	if err != nil {
		return "", ""
	}

	return key.Secret(), key.URL()
}

// Validate 验证
// return bool
func Validate(passcode, key string) bool {
	valid := totp.Validate(passcode, key)
	if valid {
		return true
	} else {
		return false
	}
}
