package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	ServiceHost    string `env:"RUN_ADDRESS"`
	DatabaseURI    string `env:"DATABASE_URI"`
	ServiceAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

const (
	DBhost           = "localhost"
	DBuser           = "postgres"
	DBpassword       = "123"
	DBname           = "postgres"
	UpdateDuration   = time.Second * 10
	ClientTimeout    = time.Second * 10
	OrdersBatchNum   = 20
	StatusNew        = "NEW"
	StatusProcessing = "PROCESSING"
	StatusRegistered = "REGISTERED"
)

func NewConfig() *Config {

	config := Config{}

	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBhost, DBuser, DBpassword, DBname)

	flag.StringVar(&config.ServiceHost, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&config.DatabaseURI, "d", ps, "it's conn string")
	//flag.StringVar(&config.DatabaseURI, "d", "", "it's conn string")
	flag.StringVar(&config.ServiceAddress, "r", "", "It's a accrual system address")

	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
