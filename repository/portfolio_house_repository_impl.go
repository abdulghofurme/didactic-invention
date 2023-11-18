package repository

import (
	"context"
	"database/sql"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/domain"
)

func NewPortfolioHouseRepository() PortfolioHouseRepository {
	return &PortfolioHouseRepositoryImpl{}
}

type PortfolioHouseRepositoryImpl struct{}

func (repository *PortfolioHouseRepositoryImpl) Create(
	ctx context.Context,
	tx *sql.Tx,
	portfolioHouse domain.PortfolioHouse,
) domain.PortfolioHouse {
	SQL := `insert into 
		portfolio_houses(id, house_id, portfolio_id)
		values(?, ?, ?)`
	_, err := tx.ExecContext(
		ctx,
		SQL,
		portfolioHouse.ID,
		portfolioHouse.HouseID,
		portfolioHouse.PortfolioID,
	)
	helper.PanicIfError(err)

	portfolioHouse = repository.FindByID(ctx, tx, portfolioHouse.ID)
	return portfolioHouse
}

func (repository *PortfolioHouseRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	portfolioHouse domain.PortfolioHouse,
) domain.PortfolioHouse {
	SQL := `update portfolio_houses set
		house_id=?, portfolio_id=?
		where id=?`
	_, err := tx.ExecContext(
		ctx,
		SQL,
		portfolioHouse.HouseID,
		portfolioHouse.PortfolioID,
		portfolioHouse.ID,
	)
	helper.PanicIfError(err)

	portfolioHouse = repository.FindByID(ctx, tx, portfolioHouse.ID)
	return portfolioHouse
}

func (repository *PortfolioHouseRepositoryImpl) Delete(
	ctx context.Context,
	tx *sql.Tx,
	portfolioHouse domain.PortfolioHouse,
) {
	SQL := `update portfolio_houses set deleted_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, portfolioHouse.DeletedAt.Time, portfolioHouse.ID)
	helper.PanicIfError(err)
}

func (repository *PortfolioHouseRepositoryImpl) FindByID(
	ctx context.Context,
	tx *sql.Tx,
	portfolioHouseId string,
) domain.PortfolioHouse {
	SQL := `select 
	id, created_at, updated_at, deleted_at 
	from portfolio_houses where id=? and deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL, portfolioHouseId)
	helper.PanicIfError(err)
	defer rows.Close()

	portfolioHouse := domain.PortfolioHouse{}

	if rows.Next() {
		err = rows.Scan(
			&portfolioHouse.ID,
			&portfolioHouse.CreatedAt,
			&portfolioHouse.UpdatedAt,
			&portfolioHouse.DeletedAt,
		)
		helper.PanicIfError(err)
	}
	return portfolioHouse

}
func (repository *PortfolioHouseRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.PortfolioHouse {
	SQL := `select 
	id, created_at, updated_at, deleted_at 
	from portfolio_houses where deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var portfolio_houses []domain.PortfolioHouse
	for rows.Next() {
		portfolio_house := domain.PortfolioHouse{}
		err = rows.Scan(
			&portfolio_house.ID,
			&portfolio_house.CreatedAt,
			&portfolio_house.UpdatedAt,
			&portfolio_house.DeletedAt,
		)
		helper.PanicIfError(err)
		portfolio_houses = append(portfolio_houses, portfolio_house)
	}
	return portfolio_houses
}
