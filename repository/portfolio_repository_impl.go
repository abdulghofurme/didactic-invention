package repository

import (
	"context"
	"database/sql"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/domain"
)

func NewPortfolioRepository() PortfolioRepository {
	return &PortfolioRepositoryImpl{}
}

type PortfolioRepositoryImpl struct{}

func (repository *PortfolioRepositoryImpl) Create(
	ctx context.Context,
	tx *sql.Tx,
	portfolio domain.Portfolio,
) domain.Portfolio {
	SQL := `insert into 
		portfolios(id, name, description, balance, nominal)
		values(?, ?, ?, ?, ?)`
	_, err := tx.ExecContext(
		ctx,
		SQL,
		portfolio.ID,
		portfolio.Name,
		portfolio.Description,
		portfolio.Balance,
		portfolio.Nominal,
	)
	helper.PanicIfError(err)

	portfolio = repository.FindByID(ctx, tx, portfolio.ID)
	return portfolio
}

func (repository *PortfolioRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	portfolio domain.Portfolio,
) domain.Portfolio {
	SQL := `update portfolios set
		name=?, description=?, balance=?, nominal=?, updated_at=?
		where id=?`
	_, err := tx.ExecContext(
		ctx,
		SQL,
		portfolio.Name,
		portfolio.Description,
		portfolio.Balance,
		portfolio.Nominal,
		portfolio.UpdatedAt,
		portfolio.ID,
	)
	helper.PanicIfError(err)

	portfolio = repository.FindByID(ctx, tx, portfolio.ID)
	return portfolio
}

func (repository *PortfolioRepositoryImpl) Delete(
	ctx context.Context,
	tx *sql.Tx, portfolio domain.Portfolio,
) {
	SQL := `update portfolios set deleted_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, portfolio.DeletedAt.Time, portfolio.ID)
	helper.PanicIfError(err)
}

func (repository *PortfolioRepositoryImpl) FindByID(
	ctx context.Context,
	tx *sql.Tx,
	portfolioId string,
) domain.Portfolio {
	SQL := `select 
	id, name, description, balance, nominal, created_at, updated_at, deleted_at 
	from portfolios where id=? and deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL, portfolioId)
	helper.PanicIfError(err)
	defer rows.Close()

	portfolio := domain.Portfolio{}

	if rows.Next() {
		err = rows.Scan(
			&portfolio.ID,
			&portfolio.Name,
			&portfolio.Description,
			&portfolio.Balance,
			&portfolio.Nominal,
			&portfolio.CreatedAt,
			&portfolio.UpdatedAt,
			&portfolio.DeletedAt,
		)
		helper.PanicIfError(err)
	}
	return portfolio

}
func (repository *PortfolioRepositoryImpl) FindByName(
	ctx context.Context,
	tx *sql.Tx,
	portfolioName string,
) domain.Portfolio {
	SQL := `select 
	id, name, description, balance, nominal, created_at, updated_at, deleted_at 
	from portfolios where name=?`
	rows, err := tx.QueryContext(ctx, SQL, portfolioName)
	helper.PanicIfError(err)
	defer rows.Close()

	portfolio := domain.Portfolio{}

	if rows.Next() {
		err = rows.Scan(
			&portfolio.ID,
			&portfolio.Name,
			&portfolio.Description,
			&portfolio.Balance,
			&portfolio.Nominal,
			&portfolio.CreatedAt,
			&portfolio.UpdatedAt,
			&portfolio.DeletedAt,
		)
		helper.PanicIfError(err)
	}
	return portfolio
}

func (repository *PortfolioRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Portfolio {
	SQL := `select 
	id, name, description, balance, nominal, created_at, updated_at, deleted_at 
	from portfolios where deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var portfolios []domain.Portfolio
	for rows.Next() {
		portfolio := domain.Portfolio{}
		err = rows.Scan(
			&portfolio.ID,
			&portfolio.Name,
			&portfolio.Description,
			&portfolio.Balance,
			&portfolio.Nominal,
			&portfolio.CreatedAt,
			&portfolio.UpdatedAt,
			&portfolio.DeletedAt,
		)
		helper.PanicIfError(err)
		portfolios = append(portfolios, portfolio)
	}
	return portfolios
}
