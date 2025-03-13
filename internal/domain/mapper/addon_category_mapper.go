package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonCategoryMapper struct {
}

func NewAddonCategoryMapper() AddonCategoryMapper {
	return AddonCategoryMapper{}
}

func (AddonCategoryMapper) MapRequestToEntity(ctx context.Context, request dto.AddonCategoryRequest) entity.AddonCategory {
	return entity.AddonCategory{
		Name: request.Name,
	}
}

func (m AddonCategoryMapper) MapEntityToResponse(ctx context.Context, addonCategoryEntity entity.AddonCategory) dto.AddonCategoryResponse {
	return dto.AddonCategoryResponse{
		Id:        addonCategoryEntity.Id,
		Name:      addonCategoryEntity.Name,
		CreatedAt: addonCategoryEntity.CreatedAt,
		UpdatedAt: addonCategoryEntity.UpdatedAt,
		DeletedAt: addonCategoryEntity.DeletedAt,
	}
}

func (m AddonCategoryMapper) MapEntitiesToResponses(ctx context.Context, addonCategoryEntities []entity.AddonCategory) []dto.AddonCategoryResponse {
	addonCategoryResponses := make([]dto.AddonCategoryResponse, len(addonCategoryEntities))

	for i, addonCategoryEntity := range addonCategoryEntities {
		addonCategoryResponses[i] = m.MapEntityToResponse(ctx, addonCategoryEntity)
	}

	return addonCategoryResponses
}
