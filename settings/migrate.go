package settings

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"person/apps/user/models"
	"person/core"
)

func MigrateTable(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},       // 用户模块
		models.ActionLogs{}, // 用户日志
		models.LoginLogs{},  // 登陆日志
	)
	if err != nil {
		core.LOGGER.Error("数据库迁移失败", zap.Any("err", err))
		os.Exit(0)
	}

	core.LOGGER.Info("数据库迁移成功")
}
