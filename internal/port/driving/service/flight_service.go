package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FlightService interface {
	CreateFlight(ctx context.Context, request dto.FlightRequest) (dto.FlightResponse, error)
	GetFlightById(ctx context.Context, id int64) (dto.FlightResponse, error)
	GetAllFlight(ctx context.Context, request dto.GetAllFlightRequest) ([]dto.FlightResponse, error)
	UpdateFlight(ctx context.Context, id int64, request dto.FlightRequest) (dto.FlightResponse, error)
	DeleteFlight(ctx context.Context, id int64) (dto.FlightResponse, error)
}
