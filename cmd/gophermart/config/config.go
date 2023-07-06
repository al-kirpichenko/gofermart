package config

import (
	"flag"
	"fmt"
)

type Config struct {
	ServiceHost    string `env:"RUN_ADDRESS"`
	DatabaseURI    string `env:"DATABASE_URI"`
	ServiceAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

const (
	DBhost     = "localhost"
	DBuser     = "postgres"
	DBpassword = "123"
	DBname     = "postgres"
)

func NewConfig() *Config {

	config := Config{}

	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBhost, DBuser, DBpassword, DBname)

	flag.StringVar(&config.ServiceHost, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&config.DatabaseURI, "d", ps, "it's conn string")
	//flag.StringVar(&config.DatabaseURI, "d", "", "it's conn string")
	flag.StringVar(&config.ServiceAddress, "r", "localhost:8088", "It's a accrual system address")

	return &config
}
