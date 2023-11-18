package service

import (
	"context"

	"github.com/abdulghofurme/go-mkr/model/web"
)

type PortfolioHouseService interface {
	Create(ctx context.Context, portfolioHouseRequest web.PortfolioHouseCreateRequest) web.PortfolioHouseResponse
	Update(ctx context.Context, portfolioHouseRequest web.PortfolioHouseUpdateRequest) web.PortfolioHouseResponse
	Delete(ctx context.Context, portfolioHouseId string) web.PortfolioHouseResponse
	FindByID(ctx context.Context, portfolioHouseId string) web.PortfolioHouseResponse
	FindAll(ctx context.Context) []web.PortfolioHouseResponse
}
