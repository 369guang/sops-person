package settings

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"person/apps/user/models"
	"person/core"
	"person/core/auth"
)

func CreateSuperUser(db *gorm.DB, username, password string) {
	user := models.User{
		Username: username,
		Password: password,
	}
	err := db.Create(&user).Error
	if err != nil {
		core.LOGGER.Error("创建超级管理员失败", zap.Any("err", err))
		os.Exit(0)
	}

	// 创建用户角色
	key, url := auth.CreateKey(username)
	authTotp := models.AuthToTp{
		UserID: user.ID,
		Keys:   key,
		Urls:   url,
	}
	core.DATABASE.Create(&authTotp)

	core.LOGGER.Info("创建超级管理员 " + username + " 成功 !!!")
}
