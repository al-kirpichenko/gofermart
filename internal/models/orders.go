package models

import "time"

type Orders struct {
	ID       int
	Number   string
	Status   string
	Accrual  int
	Uploaded time.Time
}
