package global

import (
	"GoLangStudy/config"
	"gorm.io/gorm"
)

// 全局变量
var (
	GvaMysqlClient *gorm.DB
	GvaConfig      config.ServerConfig
)
