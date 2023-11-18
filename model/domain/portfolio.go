package domain

import (
	"database/sql"
	"time"
)

type Portfolio struct {
	ID          string
	Name        string
	Description string
	Balance     int
	Nominal     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
