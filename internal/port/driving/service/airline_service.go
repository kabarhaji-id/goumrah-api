package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirlineService interface {
	CreateAirline(ctx context.Context, request dto.AirlineRequest) (dto.AirlineResponse, error)
	GetAirlineById(ctx context.Context, id int64) (dto.AirlineResponse, error)
	GetAllAirline(ctx context.Context, request dto.GetAllAirlineRequest) ([]dto.AirlineResponse, error)
	UpdateAirline(ctx context.Context, id int64, request dto.AirlineRequest) (dto.AirlineResponse, error)
	DeleteAirline(ctx context.Context, id int64) (dto.AirlineResponse, error)
}
