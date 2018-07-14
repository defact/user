package config

import (
	"github.com/jinzhu/configor"
)

type Configuration = struct {
	Port            int
	PgConnection    string `env:"DATABASE_URL"`
	RedisConnection string `env:"REDIS_URL"`
	MigrationsPath  string
}

func Load(configuration *Configuration) {
	configor.Load(configuration, "config/config.json")
}
