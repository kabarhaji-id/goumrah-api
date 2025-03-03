package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
)

type EmbarkationMapper struct {
}

func NewEmbarkationMapper() EmbarkationMapper {
	return EmbarkationMapper{}
}

func (EmbarkationMapper) MapRequestToEntity(ctx context.Context, request dto.EmbarkationRequest) entity.Embarkation {
	return entity.Embarkation{
		Name:      request.Name,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Slug:      sluger.Slug(request.Name),
	}
}

func (m EmbarkationMapper) MapEntityToResponse(ctx context.Context, embarkationEntity entity.Embarkation) dto.EmbarkationResponse {
	return dto.EmbarkationResponse{
		Id:        embarkationEntity.Id,
		Name:      embarkationEntity.Name,
		Latitude:  embarkationEntity.Latitude,
		Longitude: embarkationEntity.Longitude,
		Slug:      embarkationEntity.Slug,
		CreatedAt: embarkationEntity.CreatedAt,
		UpdatedAt: embarkationEntity.UpdatedAt,
		DeletedAt: embarkationEntity.DeletedAt,
	}
}

func (m EmbarkationMapper) MapEntitiesToResponses(ctx context.Context, embarkationEntities []entity.Embarkation) []dto.EmbarkationResponse {
	embarkationResponses := make([]dto.EmbarkationResponse, len(embarkationEntities))

	for i, embarkationEntity := range embarkationEntities {
		embarkationResponses[i] = m.MapEntityToResponse(ctx, embarkationEntity)
	}

	return embarkationResponses
}
