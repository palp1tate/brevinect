package global

import (
	"github.com/palp1tate/brevinect/service/meeting/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	Debug        bool
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
