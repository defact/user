package config

import (
	"github.com/jinzhu/configor"
)

type Configuration = struct {
	Port            int    `env:"PORT"`
	PgConnection    string `env:"DATABASE_URL"`
	RedisConnection string `env:"REDIS_URL"`
	MigrationsPath  string
	Secret          string `env:"TOKEN_SECRET"`
}

func Load(configuration *Configuration) {
	configor.Load(configuration, "config/config.json")
}
