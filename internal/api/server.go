package api

import (
	"github.com/jinzhu/gorm"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/database"
)

type Server struct {
	config *config.Config
	DB     *gorm.DB
}

func NewServer(config *config.Config) *Server {

	return &Server{
		config: config,
		DB:     database.InitDB(config.DatabaseURI),
	}
}
