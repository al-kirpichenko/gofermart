package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Number   string
	Status   string
	Accrual  int
	Uploaded time.Time
}
