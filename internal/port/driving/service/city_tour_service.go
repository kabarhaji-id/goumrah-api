package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type CityTourService interface {
	CreateCityTour(ctx context.Context, request dto.CityTourRequest) (dto.CityTourResponse, error)
	GetCityTourById(ctx context.Context, id int64) (dto.CityTourResponse, error)
	GetAllCityTour(ctx context.Context, request dto.GetAllCityTourRequest) ([]dto.CityTourResponse, error)
	UpdateCityTour(ctx context.Context, id int64, request dto.CityTourRequest) (dto.CityTourResponse, error)
	DeleteCityTour(ctx context.Context, id int64) (dto.CityTourResponse, error)
}
