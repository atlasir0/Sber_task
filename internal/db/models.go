// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int32
	Title       string
	Description sql.NullString
	Date        time.Time
	Completed   bool
}
