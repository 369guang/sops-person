package settings

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"person/core"
)

func MigrateTable(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		core.LOGGER.Error("数据库迁移失败", zap.Any("err", err))
		os.Exit(0)
	}

	core.LOGGER.Info("数据库迁移成功")
}
