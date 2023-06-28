package main

import (
	"log"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/app"
	"github.com/al-kirpichenko/gofermart/internal/router"
)

func main() {

	cfg := config.NewConfig()
	application := app.NewApp(cfg)

	r := router.Router(application)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("dont start it!")
		return
	}
}
