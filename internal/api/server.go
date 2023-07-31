package api

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/database"
)

type Server struct {
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewServer(config *config.Config) *Server {

	return &Server{
		Config: config,
		DB:     database.InitDB(config.DatabaseURI),
	}
}
