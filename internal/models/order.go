package models

import (
	"time"
)

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time `json:"-"`
	Number    string    `gorm:"not null;unique" json:"number"`
	Status    string    `json:"status"`
	Accrual   int       `json:"accrual"`
	UpdatedAt time.Time `json:"uploaded_at"`
	DeletedAt time.Time `json:"-"`
	UserID    uint      `json:"-"`
}
