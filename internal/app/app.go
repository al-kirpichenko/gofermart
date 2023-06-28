package app

import "github.com/al-kirpichenko/gofermart/cmd/gophermart/config"

type Application struct {
	config *config.Config
}

func NewApp(config *config.Config) *Application {

	return &Application{
		config: config,
	}
}
