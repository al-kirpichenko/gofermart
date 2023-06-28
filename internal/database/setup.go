package database

import (
	"github.com/jinzhu/gorm"

	"github.com/al-kirpichenko/gofermart/internal/models"
)

func InitDB(conf string) *gorm.DB {
	db, err := gorm.Open("pgx", conf)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Orders{})

	return db
}
