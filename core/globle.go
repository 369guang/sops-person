package core

import (
	"github.com/gofiber/fiber/v2"
	//"github.com/RichardKnop/machinery/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DATABASE *gorm.DB
	VIPER    *viper.Viper
	LOGGER   *zap.Logger
	//TASKS    *machinery.Server
)

type Router interface {
	InstallRouter(app *fiber.App)
}
