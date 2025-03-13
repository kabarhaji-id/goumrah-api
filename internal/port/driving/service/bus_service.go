package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type BusService interface {
	CreateBus(ctx context.Context, request dto.BusRequest) (dto.BusResponse, error)
	GetBusById(ctx context.Context, id int64) (dto.BusResponse, error)
	GetAllBus(ctx context.Context, request dto.GetAllBusRequest) ([]dto.BusResponse, error)
	UpdateBus(ctx context.Context, id int64, request dto.BusRequest) (dto.BusResponse, error)
	DeleteBus(ctx context.Context, id int64) (dto.BusResponse, error)
}
