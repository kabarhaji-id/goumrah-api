package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FacilityService interface {
	CreateFacility(ctx context.Context, request dto.FacilityRequest) (dto.FacilityResponse, error)
	GetFacilityById(ctx context.Context, id int64) (dto.FacilityResponse, error)
	GetAllFacility(ctx context.Context, request dto.GetAllFacilityRequest) ([]dto.FacilityResponse, error)
	UpdateFacility(ctx context.Context, id int64, request dto.FacilityRequest) (dto.FacilityResponse, error)
	DeleteFacility(ctx context.Context, id int64) (dto.FacilityResponse, error)
}
