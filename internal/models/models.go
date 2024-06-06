package models

import "time"

type Task struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Date        time.Time `json:"date" db:"date"`
	Completed   bool      `json:"completed" db:"completed"`
}
