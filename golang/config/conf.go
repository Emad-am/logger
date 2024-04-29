package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yml"
	configPath = "./"
)

func init() {

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		fmt.Println("successfully loaded the config file")
	}
}
func GetConfig() *Config {
	return &Config{
		Db: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Name:     viper.GetString("db.name"),
			UserName: viper.GetString("db.userName"),
			Password: viper.GetString("db.password"),
		},
		Redis: RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		},
	}
}
