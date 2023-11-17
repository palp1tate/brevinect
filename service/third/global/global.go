package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/brevinect/service/third/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	Debug        bool
	RedisClient  *redis.Client
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
