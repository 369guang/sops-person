package settings

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
	"person/core"
	"person/core/database"
)

func InitGorm() *gorm.DB {
	return database.OpenDatabase(core.VIPER.GetString("db.username"),
		core.VIPER.GetString("db.password"),
		core.VIPER.GetString("db.host"),
		core.VIPER.GetString("db.port"),
		core.VIPER.GetString("db.name"))
}

func LoadMiddleware(app *fiber.App) {
	// middleware
	app.Use(logger.New())                           // api日志
	app.Use(recover.New())                          // recover
	app.Use(cors.New())                             // 跨域
	app.Use(pprof.New())                            // 开启pprof分析
	app.Get("/metrics", monitor.New(monitor.Config{ // 监控
		Title: "Service Metrics Page"}))
}
