package redis

import (
	"fmt"
	"logger/config"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func init() {
	cnfg := config.GetConfig()
	connectionString := fmt.Sprintf("%s:%s", cnfg.Redis.Host, cnfg.Redis.Port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: cnfg.Redis.Password,
		DB:       cnfg.Redis.DB,
	})

	if Rdb == nil {
		panic("failed to connect to redis client")
	}

	fmt.Println("successfully connected to redis client")
}

func GetRedis() *redis.Client {
	return Rdb
}
