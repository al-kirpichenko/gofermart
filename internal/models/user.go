package models

type User struct {
	ID        int
	Login     string
	Password  string
	Balance   int
	Withdrawn int
}
