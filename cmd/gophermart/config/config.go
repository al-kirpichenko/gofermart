package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Host        string `env:"RUN_ADDRESS"`
	DatabaseURI string `env:"DATABASE_URI"`
	SysAddress  string `ACCRUAL_SYSTEM_ADDRESS"`
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

	flag.StringVar(&config.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&config.DatabaseURI, "d", ps, "it's conn string")
	flag.StringVar(&config.SysAddress, "r", "", "It's a FilePATH")

	return &config
}
