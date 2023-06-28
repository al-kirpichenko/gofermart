package repo

import (
	"github.com/jinzhu/gorm"

	"github.com/al-kirpichenko/gofermart/internal/models"
)

type UserRepo struct {
	Store *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Store: db,
	}
}

func (r *UserRepo) CreateUser(user *models.User) {
	r.Store.Create(&user)
}
