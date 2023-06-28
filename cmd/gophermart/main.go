package main

import (
	config2 "github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	app "github.com/al-kirpichenko/gofermart/internal/app"
	router "github.com/al-kirpichenko/gofermart/internal/router"
)

func main() {

	config := config2.NewConfig()
	application := app.NewApp(config)

	router := router.Router(application)
}
