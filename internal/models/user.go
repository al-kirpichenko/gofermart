package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login     string  `gorm:"size:255;not null;unique" json:"login"`
	Password  string  `gorm:"size:255;not null;unique" json:"password"`
	Balance   float32 `json:"-"`
	Withdrawn float32 `json:"-"`
}
