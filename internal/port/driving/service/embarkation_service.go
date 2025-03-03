package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type EmbarkationService interface {
	CreateEmbarkation(ctx context.Context, request dto.EmbarkationRequest) (dto.EmbarkationResponse, error)
	GetEmbarkationById(ctx context.Context, id int64) (dto.EmbarkationResponse, error)
	GetAllEmbarkation(ctx context.Context, request dto.GetAllEmbarkationRequest) ([]dto.EmbarkationResponse, error)
	UpdateEmbarkation(ctx context.Context, id int64, request dto.EmbarkationRequest) (dto.EmbarkationResponse, error)
	DeleteEmbarkation(ctx context.Context, id int64) (dto.EmbarkationResponse, error)
}
