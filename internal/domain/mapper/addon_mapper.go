package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonMapper struct {
	addonCategoryRepository repository.AddonCategoryRepository
	addonCategoryMapper     AddonCategoryMapper
}

func NewAddonMapper(
	addonCategoryRepository repository.AddonCategoryRepository,
	addonCategoryMapper AddonCategoryMapper,
) AddonMapper {
	return AddonMapper{
		addonCategoryRepository, addonCategoryMapper,
	}
}

func (AddonMapper) MapRequestToEntity(ctx context.Context, request dto.AddonRequest) entity.Addon {
	return entity.Addon{
		CategoryId: request.Category,
		Name:       request.Name,
		Price:      request.Price,
	}
}

func (m AddonMapper) MapEntityToResponse(ctx context.Context, addonEntity entity.Addon) (dto.AddonResponse, error) {
	categoryEntity, err := m.addonCategoryRepository.FindById(ctx, addonEntity.CategoryId)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	categoryResponse := m.addonCategoryMapper.MapEntityToResponse(ctx, categoryEntity)

	return dto.AddonResponse{
		Id:        addonEntity.Id,
		Category:  categoryResponse,
		Name:      addonEntity.Name,
		Price:     addonEntity.Price,
		CreatedAt: addonEntity.CreatedAt,
		UpdatedAt: addonEntity.UpdatedAt,
		DeletedAt: addonEntity.DeletedAt,
	}, nil
}

func (m AddonMapper) MapEntitiesToResponses(ctx context.Context, addonEntities []entity.Addon) ([]dto.AddonResponse, error) {
	addonResponses := make([]dto.AddonResponse, len(addonEntities))
	var err error

	for i, addonEntity := range addonEntities {
		addonResponses[i], err = m.MapEntityToResponse(ctx, addonEntity)
		if err != nil {
			return nil, err
		}
	}

	return addonResponses, nil
}
