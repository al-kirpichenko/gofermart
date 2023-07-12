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

	errUsers := db.AutoMigrate(&models.User{})
	if errUsers != nil {
		log.Fatal(err)
		return nil
	}
	errOrders := db.AutoMigrate(&models.Order{})
	if errOrders != nil {
		log.Fatal(err)
		return nil
	}
	errWithdraws := db.AutoMigrate(&models.Withdraw{})
	if errWithdraws != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
