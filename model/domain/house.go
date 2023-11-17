package domain

import (
	"database/sql"
	"time"
)

type House struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
