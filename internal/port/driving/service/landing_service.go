package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type LandingService interface {
	CreateLanding(ctx context.Context, request dto.LandingRequest) (dto.LandingResponse, error)
	GetLanding(ctx context.Context) (dto.LandingResponse, error)
	UpdateLanding(ctx context.Context, request dto.LandingRequest) (dto.LandingResponse, error)
}
