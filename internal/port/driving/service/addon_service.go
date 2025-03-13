package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonService interface {
	CreateAddon(ctx context.Context, request dto.AddonRequest) (dto.AddonResponse, error)
	GetAddonById(ctx context.Context, id int64) (dto.AddonResponse, error)
	GetAllAddon(ctx context.Context, request dto.GetAllAddonRequest) ([]dto.AddonResponse, error)
	UpdateAddon(ctx context.Context, id int64, request dto.AddonRequest) (dto.AddonResponse, error)
	DeleteAddon(ctx context.Context, id int64) (dto.AddonResponse, error)
}
