package repository

import (
	"context"
	"database/sql"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/domain"
)

func NewHouseRepository() HouseRepository {
	return &HouseRepositoryImpl{}
}

type HouseRepositoryImpl struct{}

func (repository *HouseRepositoryImpl) Create(
	ctx context.Context,
	tx *sql.Tx,
	house domain.House,
) domain.House {
	SQL := `insert into houses(id, block_name, block_number) values(?, ?, ?)`
	_, err := tx.ExecContext(ctx, SQL, house.ID, house.BlockName, house.BlockNumber)
	helper.PanicIfError(err)

	house = repository.FindByID(ctx, tx, house.ID)
	return house
}

func (repository *HouseRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	house domain.House,
) domain.House {
	SQL := `update houses set block_name=?, block_number=?, updated_at=? where id=?`

	_, err := tx.ExecContext(
		ctx,
		SQL,
		house.BlockName,
		house.BlockNumber,
		house.UpdatedAt,
		house.ID,
	)
	helper.PanicIfError(err)

	house = repository.FindByID(ctx, tx, house.ID)
	return house
}

func (repository *HouseRepositoryImpl) Delete(
	ctx context.Context,
	tx *sql.Tx,
	house domain.House,
) {
	SQL := `update houses set deleted_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, house.DeletedAt.Time, house.ID)
	helper.PanicIfError(err)
}

func (repository *HouseRepositoryImpl) FindByID(
	ctx context.Context,
	tx *sql.Tx,
	houseId string,
) domain.House {
	SQL := `select 
	id, block_name, block_number, created_at, updated_at, deleted_at 
	from houses where id=? and deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL, houseId)
	helper.PanicIfError(err)
	defer rows.Close()

	house := domain.House{}
	if rows.Next() {
		err := rows.Scan(
			&house.ID,
			&house.BlockName,
			&house.BlockNumber,
			&house.CreatedAt,
			&house.UpdatedAt,
			&house.DeletedAt,
		)
		helper.PanicIfError(err)
	}

	return house
}

func (repository *HouseRepositoryImpl) FindByBlockNumber(
	ctx context.Context,
	tx *sql.Tx,
	houseBlock string,
	houseNumber int,
) domain.House {
	SQL := `select 
	id, block_name, block_number, created_at, updated_at, deleted_at 
	from houses where block_name=? and block_number=?`
	rows, err := tx.QueryContext(ctx, SQL, houseBlock, houseNumber)
	helper.PanicIfError(err)
	defer rows.Close()

	house := domain.House{}
	if rows.Next() {
		err := rows.Scan(
			&house.ID,
			&house.BlockName,
			&house.BlockNumber,
			&house.CreatedAt,
			&house.UpdatedAt,
			&house.DeletedAt,
		)
		helper.PanicIfError(err)
	}

	return house
}

func (repository *HouseRepositoryImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
) []domain.House {
	SQL := `select 
	id, block_name, block_number, created_at, updated_at, deleted_at 
	from houses where deleted_at is null 
	order by block_name, block_number`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var houses []domain.House

	for rows.Next() {
		house := domain.House{}
		err := rows.Scan(
			&house.ID,
			&house.BlockName,
			&house.BlockNumber,
			&house.CreatedAt,
			&house.UpdatedAt,
			&house.DeletedAt,
		)
		helper.PanicIfError(err)

		houses = append(houses, house)
	}

	return houses
}
