package api

import (
	"gorm.io/gorm"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/gofermart/cmd/gofermart/config"
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
