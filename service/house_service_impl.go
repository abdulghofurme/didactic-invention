package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/web"
	"github.com/abdulghofurme/go-mkr/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewHouseService(
	houseRepository repository.HouseRepository,
	db *sql.DB,
	validate *validator.Validate,
) HouseService {
	return &HouseServiceImpl{
		HouseRepository: houseRepository,
		DB:              db,
		Validate:        validate,
	}
}

type HouseServiceImpl struct {
	HouseRepository repository.HouseRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func (service *HouseServiceImpl) Create(
	ctx context.Context,
	houseRequest web.HouseCreateRequest,
) web.HouseResponse {
	err := service.Validate.Struct(houseRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	house := service.HouseRepository.FindByBlockNumber(
		ctx,
		tx,
		houseRequest.BlockName,
		houseRequest.BlockNumber,
	)
	house.ID = uuid.NewString()
	house.BlockName = houseRequest.BlockName
	house.BlockNumber = houseRequest.BlockNumber

	house = service.HouseRepository.Create(ctx, tx, house)
	return helper.ToHouseResponse(&house)
}

func (service *HouseServiceImpl) Update(
	ctx context.Context,
	houseRequest web.HouseUpdateRequest,
) web.HouseResponse {
	err := service.Validate.Struct(houseRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	house := service.HouseRepository.FindByBlockNumber(
		ctx,
		tx,
		houseRequest.BlockName,
		houseRequest.BlockNumber,
	)
	if house.ID != houseRequest.ID {
		panic(fmt.Sprintf("%v%v sudah digunakan", houseRequest.BlockName, houseRequest.BlockNumber))
	}
	if house.DeletedAt.Valid {
		panic(fmt.Sprintf("%v%v sudah tidak lagi aktif", houseRequest.BlockName, houseRequest.BlockNumber))
	}
	house.ID = houseRequest.ID
	house.BlockName = houseRequest.BlockName
	house.BlockNumber = houseRequest.BlockNumber
	house.UpdatedAt = time.Now()

	house = service.HouseRepository.Update(ctx, tx, house)
	return helper.ToHouseResponse(&house)
}

func (service *HouseServiceImpl) Delete(
	ctx context.Context,
	houseId string,
) web.HouseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	house := service.HouseRepository.FindByID(ctx, tx, houseId)
	if house.ID == "" {
		return helper.ToHouseResponse(&house)
	}
	house.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	service.HouseRepository.Delete(ctx, tx, house)

	return helper.ToHouseResponse(&house)
}

func (service *HouseServiceImpl) FindByID(
	ctx context.Context,
	houseId string,
) web.HouseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	house := service.HouseRepository.FindByID(ctx, tx, houseId)
	return helper.ToHouseResponse(&house)
}

func (service *HouseServiceImpl) FindAll(ctx context.Context) []web.HouseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	houses := service.HouseRepository.FindAll(ctx, tx)

	housesResponse := []web.HouseResponse{}
	for _, house := range houses {
		housesResponse = append(housesResponse, helper.ToHouseResponse(&house))
	}
	return housesResponse
}
