package service

import (
	"context"

	"github.com/abdulghofurme/go-mkr/model/web"
)

type HouseService interface {
	Create(ctx context.Context, houseRequest web.HouseCreateRequest) web.HouseResponse
	Update(ctx context.Context, houseRequest web.HouseUpdateRequest) web.HouseResponse
	Delete(ctx context.Context, houseId string) web.HouseResponse
	FindByID(ctx context.Context, houseId string) web.HouseResponse
	FindAll(ctx context.Context) []web.HouseResponse
}
