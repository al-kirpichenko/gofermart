package database

import (
	"log"

	"github.com/al-kirpichenko/gofermart/internal/models"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func InitDB(conf string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conf), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Withdraw{})

	return db
}
