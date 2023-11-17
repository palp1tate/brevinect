package initialize

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/brevinect/service/third/global"
)

func InitRedis() {
	r := global.ServerConfig.Redis
	host := r.Host
	port := r.Port
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       r.Database,
	})
}
