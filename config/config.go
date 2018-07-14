package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Configuration = struct {
	PgConnection    string `env:"DATABASE_URL"`
	RedisConnection string `env:"REDIS_URL"`
	MigrationsPath  string
}

func Load(configuration *Configuration) {
	configor.Load(configuration, "./config/config.json")
	fmt.Printf("config: %#v", configuration)
}
