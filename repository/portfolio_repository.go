package repository

import (
	"context"
	"database/sql"

	"github.com/abdulghofurme/go-mkr/model/domain"
)

type PortfolioRepository interface {
	Create(ctx context.Context, tx *sql.Tx, portfolio domain.Portfolio) domain.Portfolio
	Update(ctx context.Context, tx *sql.Tx, portfolio domain.Portfolio) domain.Portfolio
	Delete(ctx context.Context, tx *sql.Tx, portfolio domain.Portfolio)
	FindByID(ctx context.Context, tx *sql.Tx, portfolioId string) domain.Portfolio
	FindByName(ctx context.Context, tx *sql.Tx, portfolioName string) domain.Portfolio
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Portfolio
}
