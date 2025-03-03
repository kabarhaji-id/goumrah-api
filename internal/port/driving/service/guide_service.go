package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type GuideService interface {
	CreateGuide(ctx context.Context, request dto.GuideRequest) (dto.GuideResponse, error)
	GetGuideById(ctx context.Context, id int64) (dto.GuideResponse, error)
	GetAllGuide(ctx context.Context, request dto.GetAllGuideRequest) ([]dto.GuideResponse, error)
	UpdateGuide(ctx context.Context, id int64, request dto.GuideRequest) (dto.GuideResponse, error)
	DeleteGuide(ctx context.Context, id int64) (dto.GuideResponse, error)
}
