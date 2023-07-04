package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Number   int `gorm:"not null;unique"`
	UserID   uint
	Status   string
	Accrual  int
	Uploaded time.Time
}
