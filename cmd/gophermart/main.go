package main

import (
	"log"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/api"
	"github.com/al-kirpichenko/gofermart/internal/router"
)

func main() {

	cfg := config.NewConfig()
	server := api.NewServer(cfg)

	defer server.DB.Close()

	r := router.Router(server)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("dont start it!")
		return
	}
}
