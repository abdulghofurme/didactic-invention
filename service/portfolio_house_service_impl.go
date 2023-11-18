package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/domain"
	"github.com/abdulghofurme/go-mkr/model/web"
	"github.com/abdulghofurme/go-mkr/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewPortfolioHouseService(
	portfolioHouseRepository repository.PortfolioHouseRepository,
	db *sql.DB,
	validate *validator.Validate,
) PortfolioHouseService {
	return &PortfolioHouseServiceImpl{
		PortfolioHouseRepository: portfolioHouseRepository,
		DB:                       db,
		Validate:                 validate,
	}
}

type PortfolioHouseServiceImpl struct {
	PortfolioHouseRepository repository.PortfolioHouseRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func (service *PortfolioHouseServiceImpl) Create(
	ctx context.Context,
	portfolioHouseRequest web.PortfolioHouseCreateRequest,
) web.PortfolioHouseResponse {
	err := service.Validate.Struct(portfolioHouseRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolioHouse := domain.PortfolioHouse{
		ID:          uuid.NewString(),
		HouseID:     portfolioHouseRequest.HouseID,
		PortfolioID: portfolioHouseRequest.PortfolioID,
	}

	portfolioHouse = service.PortfolioHouseRepository.Create(
		ctx,
		tx,
		portfolioHouse,
	)

	return helper.ToPortfolioHouseResponse(&portfolioHouse)
}

func (service *PortfolioHouseServiceImpl) Update(
	ctx context.Context,
	portfolioHouseRequest web.PortfolioHouseUpdateRequest,
) web.PortfolioHouseResponse {
	err := service.Validate.Struct(portfolioHouseRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolioHouse := service.PortfolioHouseRepository.FindByID(ctx, tx, portfolioHouseRequest.ID)
	if portfolioHouse.DeletedAt.Valid {
		panic(fmt.Sprintf("relasi portfolio house sudah tidak lagi aktif"))
	}
	portfolioHouse.ID = portfolioHouseRequest.ID
	portfolioHouse.HouseID = portfolioHouseRequest.HouseID
	portfolioHouse.PortfolioID = portfolioHouseRequest.PortfolioID
	portfolioHouse.UpdatedAt = time.Now()

	portfolioHouse = service.PortfolioHouseRepository.Update(ctx, tx, portfolioHouse)
	return helper.ToPortfolioHouseResponse(&portfolioHouse)
}

func (service *PortfolioHouseServiceImpl) Delete(
	ctx context.Context,
	portfolioHouseId string,
) web.PortfolioHouseResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolioHouse := service.PortfolioHouseRepository.FindByID(ctx, tx, portfolioHouseId)
	if portfolioHouse.ID == "" {
		return helper.ToPortfolioHouseResponse(&portfolioHouse)
	}
	portfolioHouse.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	service.PortfolioHouseRepository.Delete(ctx, tx, portfolioHouse)

	return helper.ToPortfolioHouseResponse(&portfolioHouse)
}

func (service *PortfolioHouseServiceImpl) FindByID(
	ctx context.Context,
	portfolioHouseId string,
) web.PortfolioHouseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolioHouse := service.PortfolioHouseRepository.FindByID(ctx, tx, portfolioHouseId)
	return helper.ToPortfolioHouseResponse(&portfolioHouse)
}

func (service *PortfolioHouseServiceImpl) FindAll(ctx context.Context) []web.PortfolioHouseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolioHouses := service.PortfolioHouseRepository.FindAll(ctx, tx)

	portfolioHousesResponse := []web.PortfolioHouseResponse{}
	for _, portfolioHouse := range portfolioHouses {
		portfolioHousesResponse = append(portfolioHousesResponse, helper.ToPortfolioHouseResponse(&portfolioHouse))
	}
	return portfolioHousesResponse
}
