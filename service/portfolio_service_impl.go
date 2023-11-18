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

func NewPortfolioService(
	portfolioRepository repository.PortfolioRepository,
	db *sql.DB,
	validate *validator.Validate,
) PortfolioService {
	return &PortfolioServiceImpl{
		PortfolioRepository: portfolioRepository,
		DB:                  db,
		Validate:            validate,
	}
}

type PortfolioServiceImpl struct {
	PortfolioRepository repository.PortfolioRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func (service *PortfolioServiceImpl) Create(
	ctx context.Context,
	portfolioRequest web.PortfolioCreateRequest,
) web.PortfolioResponse {
	err := service.Validate.Struct(portfolioRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolio := service.PortfolioRepository.FindByName(
		ctx,
		tx,
		portfolioRequest.Name,
	)
	portfolio.ID = uuid.NewString()
	portfolio.Name = portfolioRequest.Name
	portfolio.Description = portfolioRequest.Description
	portfolio.Balance = portfolioRequest.Balance
	portfolio.Nominal = portfolioRequest.Nominal

	portfolio = service.PortfolioRepository.Create(
		ctx,
		tx,
		portfolio,
	)

	return helper.ToPortfolioResponse(&portfolio)
}

func (service *PortfolioServiceImpl) Update(
	ctx context.Context,
	portfolioRequest web.PortfolioUpdateRequest,
) web.PortfolioResponse {
	err := service.Validate.Struct(portfolioRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolio := service.PortfolioRepository.FindByName(
		ctx,
		tx,
		portfolioRequest.Name,
	)
	if portfolio.ID != portfolioRequest.ID {
		panic(fmt.Sprintf("%v sudah digunakan", portfolioRequest.Name))
	}
	if portfolio.DeletedAt.Valid {
		panic(fmt.Sprintf("%v sudah tidak lagi aktif", portfolioRequest.Name))
	}
	portfolio.ID = portfolioRequest.ID
	portfolio.Name = portfolioRequest.Name
	portfolio.Description = portfolioRequest.Description
	portfolio.Balance = portfolioRequest.Balance
	portfolio.Nominal = portfolioRequest.Nominal
	portfolio.UpdatedAt = time.Now()

	portfolio = service.PortfolioRepository.Update(ctx, tx, portfolio)
	return helper.ToPortfolioResponse(&portfolio)
}

func (service *PortfolioServiceImpl) Delete(
	ctx context.Context,
	portfolioId string,
) web.PortfolioResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolio := service.PortfolioRepository.FindByID(ctx, tx, portfolioId)
	if portfolio.ID == "" {
		return helper.ToPortfolioResponse(&portfolio)
	}
	portfolio.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	service.PortfolioRepository.Delete(ctx, tx, portfolio)

	return helper.ToPortfolioResponse(&portfolio)
}

func (service *PortfolioServiceImpl) FindByID(
	ctx context.Context,
	portfolioId string,
) web.PortfolioResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolio := service.PortfolioRepository.FindByID(ctx, tx, portfolioId)
	return helper.ToPortfolioResponse(&portfolio)
}

func (service *PortfolioServiceImpl) FindAll(ctx context.Context) []web.PortfolioResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	portfolios := service.PortfolioRepository.FindAll(ctx, tx)

	var portfoliosResponse []web.PortfolioResponse
	for _, portfolio := range portfolios {
		portfoliosResponse = append(portfoliosResponse, helper.ToPortfolioResponse(&portfolio))
	}
	return portfoliosResponse
}
