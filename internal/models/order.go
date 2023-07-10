package models

import (
	"time"
)

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Number    string    `gorm:"not null;unique" json:"number"`
	Status    string    `json:"status"`
	Accrual   int       `json:"accrual"`
	CreatedAt time.Time `json:"uploaded_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
	UserID    uint      `json:"-"`
}
