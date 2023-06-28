package repo

import (
	"github.com/jinzhu/gorm"

	"github.com/al-kirpichenko/gofermart/internal/models"
)

type OrderRepo struct {
	Store *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Store: db,
	}
}

func (r *UserRepo) CreateOrder(user *models.User) {
	r.Store.Create(&user)
}
