package service

import (
	"context"

	"github.com/abdulghofurme/go-mkr/model/web"
)

type PortfolioService interface {
	Create(ctx context.Context, portfolioRequest web.PortfolioCreateRequest) web.PortfolioResponse
	Update(ctx context.Context, portfolioRequest web.PortfolioUpdateRequest) web.PortfolioResponse
	Delete(ctx context.Context, portfolioId string) web.PortfolioResponse
	FindByID(ctx context.Context, portfolioId string) web.PortfolioResponse
	FindAll(ctx context.Context) []web.PortfolioResponse
}
