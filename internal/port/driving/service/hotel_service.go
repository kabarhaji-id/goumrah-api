package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type HotelService interface {
	CreateHotel(ctx context.Context, request dto.HotelRequest) (dto.HotelResponse, error)
	GetHotelById(ctx context.Context, id int64) (dto.HotelResponse, error)
	GetAllHotel(ctx context.Context, request dto.GetAllHotelRequest) ([]dto.HotelResponse, error)
	UpdateHotel(ctx context.Context, id int64, request dto.HotelRequest) (dto.HotelResponse, error)
	DeleteHotel(ctx context.Context, id int64) (dto.HotelResponse, error)
}
