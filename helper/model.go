package helper

import (
	"fmt"

	"github.com/abdulghofurme/go-mkr/model/domain"
	"github.com/abdulghofurme/go-mkr/model/web"
)

func ToHouseResponse(house *domain.House) web.HouseResponse {
	return web.HouseResponse{
		ID:          house.ID,
		Name:        fmt.Sprintf("%v%v", house.BlockName, house.BlockNumber),
		BlockName:   house.BlockName,
		BlockNumber: house.BlockNumber,
		CreatedAt:   house.CreatedAt.String(),
		UpdatedAt:   house.UpdatedAt.String(),
		DeletedAt:   SQLTimeNullToString(&house.DeletedAt),
	}
}

func ToPortfolioResponse(portfolio *domain.Portfolio) web.PortfolioResponse {
	return web.PortfolioResponse{
		ID:          portfolio.ID,
		Name:        portfolio.Name,
		Description: portfolio.Description,
		Balance:     portfolio.Balance,
		Nominal:     portfolio.Nominal,
		CreatedAt:   portfolio.CreatedAt.String(),
		UpdatedAt:   portfolio.UpdatedAt.String(),
		DeletedAt:   SQLTimeNullToString(&portfolio.DeletedAt),
	}
}

func ToPortfolioHouseResponse(portfolioHouse *domain.PortfolioHouse) web.PortfolioHouseResponse {
	return web.PortfolioHouseResponse{
		ID:        portfolioHouse.ID,
		House:     ToHouseResponse(&portfolioHouse.House),
		Portfolio: ToPortfolioResponse(&portfolioHouse.Portfolio),
		CreatedAt: portfolioHouse.CreatedAt.String(),
		UpdatedAt: portfolioHouse.UpdatedAt.String(),
		DeletedAt: SQLTimeNullToString(&portfolioHouse.DeletedAt),
	}
}
