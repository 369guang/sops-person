package serializers

import (
	"person/apps/user/models"
	"person/core"
	"person/core/auth"
)

func Auth(c *auth.Context) (interface{}, error) { // 获取ToTp认证

	totp := new(models.AuthToTp)
	if err := core.DATABASE.Where("user_id = ?", c.ID).First(&totp).Error; err != nil {
		return nil, err
	}

	totp.Keys = ""

	return totp, nil
}
