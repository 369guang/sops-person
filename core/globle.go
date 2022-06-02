package core

import (
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
