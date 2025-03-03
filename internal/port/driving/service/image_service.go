package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ImageService interface {
	CreateImage(ctx context.Context, request dto.ImageRequest) (dto.ImageResponse, error)
	GetImageById(ctx context.Context, id int64) (dto.ImageResponse, error)
	GetAllImage(ctx context.Context, request dto.GetAllImageRequest) ([]dto.ImageResponse, error)
	UpdateImage(ctx context.Context, id int64, request dto.ImageRequest) (dto.ImageResponse, error)
	DeleteImage(ctx context.Context, id int64) (dto.ImageResponse, error)
}
