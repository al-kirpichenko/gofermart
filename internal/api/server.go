package api

import (
	"github.com/jinzhu/gorm"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/database"
)

type Server struct {
	config *config.Config
	DB     *gorm.DB
	logger *zap.Logger
}

func NewServer(config *config.Config) *Server {

	l, _ := zap.NewDevelopment()

	return &Server{
		config: config,
		DB:     database.InitDB(config.DatabaseURI),
		logger: l,
	}
}
