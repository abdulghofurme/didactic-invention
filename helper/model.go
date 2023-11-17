package helper

import (
	"github.com/abdulghofurme/go-mkr/model/domain"
	"github.com/abdulghofurme/go-mkr/model/web"
)

func ToHouseResponse(house *domain.House) web.HouseResponse {
	return web.HouseResponse{
		ID:        house.ID,
		Name:      house.Name,
		CreatedAt: house.CreatedAt.String(),
		UpdatedAt: house.UpdatedAt.String(),
		DeletedAt: SQLTimeNullToString(&house.DeletedAt),
	}
}
