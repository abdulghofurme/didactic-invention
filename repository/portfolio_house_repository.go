package repository

import (
	"context"
	"database/sql"

	"github.com/abdulghofurme/go-mkr/model/domain"
)

type PortfolioHouseRepository interface {
	Create(
		ctx context.Context,
		tx *sql.Tx,
		portfolioPortfolioHouse domain.PortfolioHouse,
	) domain.PortfolioHouse
	Update(
		ctx context.Context,
		tx *sql.Tx,
		portfolioPortfolioHouse domain.PortfolioHouse,
	) domain.PortfolioHouse
	Delete(
		ctx context.Context,
		tx *sql.Tx,
		portfolioPortfolioHouse domain.PortfolioHouse,
	)
	FindByID(
		ctx context.Context,
		tx *sql.Tx,
		portfolioPortfolioHouseId string,
	) domain.PortfolioHouse
	FindAll(ctx context.Context, tx *sql.Tx) []domain.PortfolioHouse
}
