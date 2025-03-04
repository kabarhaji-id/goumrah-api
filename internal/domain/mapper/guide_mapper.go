package mapper

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type GuideMapper struct {
	imageMapper ImageMapper
}

func NewGuideMapper(imageMapper ImageMapper) GuideMapper {
	return GuideMapper{
		imageMapper,
	}
}

func (GuideMapper) MapRequestToEntity(ctx context.Context, request dto.GuideRequest) entity.Guide {
	return entity.Guide{
		AvatarId:    request.Avatar,
		Name:        request.Name,
		Type:        request.Type,
		Description: request.Description,
	}
}

func (m GuideMapper) MapEntityToResponse(ctx context.Context, imageRepository repository.ImageRepository, guideEntity entity.Guide) (dto.GuideResponse, error) {
	avatarResponse := null.NewValue(dto.ImageResponse{}, false)
	if guideEntity.AvatarId.Valid {
		avatarEntity, err := imageRepository.FindById(ctx, guideEntity.AvatarId.Int64)
		if err != nil {
			return dto.GuideResponse{}, err
		}

		avatarResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, avatarEntity))
	}

	return dto.GuideResponse{
		Id:          guideEntity.Id,
		Avatar:      avatarResponse,
		Name:        guideEntity.Name,
		Type:        guideEntity.Type,
		Description: guideEntity.Description,
		CreatedAt:   guideEntity.CreatedAt,
		UpdatedAt:   guideEntity.UpdatedAt,
		DeletedAt:   guideEntity.DeletedAt,
	}, nil
}

func (m GuideMapper) MapEntitiesToResponses(ctx context.Context, imageRepository repository.ImageRepository, guideEntities []entity.Guide) ([]dto.GuideResponse, error) {
	guideResponses := make([]dto.GuideResponse, len(guideEntities))
	var err error

	for i, guideEntity := range guideEntities {
		guideResponses[i], err = m.MapEntityToResponse(ctx, imageRepository, guideEntity)
		if err != nil {
			return nil, err
		}
	}

	return guideResponses, nil
}
