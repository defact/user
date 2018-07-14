package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Configuration = struct {
	PgConnection    string `env:"DATABASE_URL"`
	RedisConnection string `env:"REDIS_URL"`
	MigrationsPath  string
}

func Load(configuration *Configuration) {
	var path, _ = os.Getwd()
	configor.Load(configuration, path+"/config/config.json")
	fmt.Printf("config: %#v", configuration)
}
