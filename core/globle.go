package core

import (
	//"github.com/RichardKnop/machinery/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DATABASE *gorm.DB
	VIPER    *viper.Viper
	//TASKS    *machinery.Server
)
