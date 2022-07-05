package models

import "person/core/database"

type AuthToTp struct {
	database.BaseModel
	UserID int    `gorm:"comment('用户ID')" json:"user_id"`
	Keys   string `gorm:"comment('密钥')" json:"keys"`
	Urls   string `gorm:"comment('映射地址')" json:"urls"`
}

func (AuthToTp) TableName() string {
	return "system_auth_totp"
}
