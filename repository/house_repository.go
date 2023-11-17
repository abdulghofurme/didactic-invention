package repository

import (
	"context"
	"database/sql"

	"github.com/abdulghofurme/go-mkr/model/domain"
)

type HouseRepository interface {
	Create(ctx context.Context, tx *sql.Tx, house domain.House) domain.House
	Update(ctx context.Context, tx *sql.Tx, house domain.House) domain.House
	Delete(ctx context.Context, tx *sql.Tx, house domain.House)
	FindByID(ctx context.Context, tx *sql.Tx, houseId string) (domain.House, error)
	FindByName(ctx context.Context, tx *sql.Tx, houseName string) (domain.House, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.House
}
