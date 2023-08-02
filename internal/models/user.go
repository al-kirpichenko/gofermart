package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login     string  `gorm:"size:255;not null;unique" json:"login"`
	Password  string  `gorm:"size:255;not null;unique" json:"password"`
	Balance   float64 `json:"-"`
	Withdrawn float64 `json:"-"`
}
