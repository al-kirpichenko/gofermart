package models

import "time"

type Order struct {
	ID       int
	Number   string
	Status   string
	Accrual  int
	Uploaded time.Time
}
