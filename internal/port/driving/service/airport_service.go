package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirportService interface {
	CreateAirport(ctx context.Context, request dto.AirportRequest) (dto.AirportResponse, error)
	GetAirportById(ctx context.Context, id int64) (dto.AirportResponse, error)
	GetAllAirport(ctx context.Context, request dto.GetAllAirportRequest) ([]dto.AirportResponse, error)
	UpdateAirport(ctx context.Context, id int64, request dto.AirportRequest) (dto.AirportResponse, error)
	DeleteAirport(ctx context.Context, id int64) (dto.AirportResponse, error)
}
