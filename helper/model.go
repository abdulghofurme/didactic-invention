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
