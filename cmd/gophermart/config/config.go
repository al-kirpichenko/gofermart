package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	ServiceHost    string `env:"RUN_ADDRESS"`
	DatabaseURI    string `env:"DATABASE_URI"`
	ServiceAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewConfig() *Config {

	config := Config{}

	flag.StringVar(&config.ServiceHost, "a", "localhost:8080", "It's a Host")

	flag.StringVar(&config.DatabaseURI, "d", "", "it's conn string")
	flag.StringVar(&config.ServiceAddress, "r", "", "It's a accrual system address")

	flag.Parse()
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
