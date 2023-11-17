package domain

import (
	"database/sql"
	"time"
)

type House struct {
	ID          string
	BlockName   string
	BlockNumber int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
