package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonCategoryService interface {
	CreateAddonCategory(ctx context.Context, request dto.AddonCategoryRequest) (dto.AddonCategoryResponse, error)
	GetAddonCategoryById(ctx context.Context, id int64) (dto.AddonCategoryResponse, error)
	GetAllAddonCategory(ctx context.Context, request dto.GetAllAddonCategoryRequest) ([]dto.AddonCategoryResponse, error)
	UpdateAddonCategory(ctx context.Context, id int64, request dto.AddonCategoryRequest) (dto.AddonCategoryResponse, error)
	DeleteAddonCategory(ctx context.Context, id int64) (dto.AddonCategoryResponse, error)
}
