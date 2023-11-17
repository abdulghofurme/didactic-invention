package repository

import (
	"context"
	"database/sql"
	"errors"

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
	SQL := `insert into houses(id, name) values(?, ?)`
	_, err := tx.ExecContext(ctx, SQL, house.ID, house.Name)
	helper.PanicIfError(err)

	house, _ = repository.FindByID(ctx, tx, house.ID)
	return house
}

func (repository *HouseRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	house domain.House,
) domain.House {
	SQL := `update houses set name=?, updated_at=?`
	_, err := tx.ExecContext(ctx, SQL, house.Name, house.UpdatedAt)
	helper.PanicIfError(err)

	house, _ = repository.FindByID(ctx, tx, house.ID)
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
) (domain.House, error) {
	SQL := `select id, name, created_at, updated_at, deleted_at from houses where id=?`
	rows, err := tx.QueryContext(ctx, SQL, houseId)
	helper.PanicIfError(err)
	defer rows.Close()

	house := domain.House{}
	if rows.Next() {
		err := rows.Scan(
			&house.ID,
			&house.Name,
			&house.CreatedAt,
			&house.UpdatedAt,
			&house.DeletedAt,
		)
		helper.PanicIfError(err)

		return house, nil
	}

	return house, errors.New("rumah tidak ditemukan")
}

func (repository *HouseRepositoryImpl) FindByName(
	ctx context.Context,
	tx *sql.Tx, houseName string,
) (domain.House, error) {
	SQL := `select id, name, created_at, updated_at, deleted_at from houses where name=?`
	rows, err := tx.QueryContext(ctx, SQL, houseName)
	helper.PanicIfError(err)
	defer rows.Close()

	house := domain.House{}
	if rows.Next() {
		err := rows.Scan(
			&house.ID,
			&house.Name,
			&house.CreatedAt,
			&house.UpdatedAt,
			&house.DeletedAt,
		)
		helper.PanicIfError(err)

		return house, nil
	}

	return house, errors.New("rumah tidak ditemukan")
}

func (repository *HouseRepositoryImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
) []domain.House {
	SQL := `select id, name, created_at, updated_at, deleted_at from houses where deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var houses []domain.House

	for rows.Next() {
		house := domain.House{}
		err := rows.Scan(
			&house.ID,
			&house.Name,
			&house.CreatedAt,
			&house.UpdatedAt,
			&house.DeletedAt,
		)
		helper.PanicIfError(err)

		houses = append(houses, house)
	}

	return houses
}
