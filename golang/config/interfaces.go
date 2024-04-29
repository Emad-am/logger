package config

type DBConfig struct {
	Host string
	Port string
	Name string
	UserName string
	Password string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Config struct {
	Db    DBConfig
	Redis RedisConfig
}
