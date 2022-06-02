package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"os"
	"person/core"
)

var (
	DB *gorm.DB
)

func OpenDatabase(username, password, addr, port, name string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", addr, username, password, name, port)
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   logger,
		DisableAutomaticPing:                     true, // 初始化后，ping数据库是否能连上
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: core.VIPER.GetString("prefix"), // 表前缀
		},
		QueryFields: true,
	})
	if err != nil {
		core.LOGGER.Error("postgres 连接失败！！！", zap.Any("error:", err))
		os.Exit(2)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(core.VIPER.GetInt("db.max_pool")) //设置连接池的空闲数大小
	sqlDB.SetMaxIdleConns(core.VIPER.GetInt("db.max_idle")) //设置最大打开连接数
	core.LOGGER.Info(fmt.Sprintf("postgres 连接成功！！！\n 连接到 %v", dsn))
	return db
}
