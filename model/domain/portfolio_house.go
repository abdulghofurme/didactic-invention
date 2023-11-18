package domain

import (
	"database/sql"
	"time"
)

type PortfolioHouse struct {
	ID          string
	HouseID     string
	House       House
	PortfolioID string
	Portfolio   Portfolio
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
