package models

import "time"

type Withdraw struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	Order       string    `json:"order"`
	Sum         float64   `json:"sum"`
	UserID      uint      `json:"-"`
	ProcessedAt time.Time `json:"processed_at"`
}
