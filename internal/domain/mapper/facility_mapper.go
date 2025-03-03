package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FacilityMapper struct {
}

func NewFacilityMapper() FacilityMapper {
	return FacilityMapper{}
}

func (FacilityMapper) MapRequestToEntity(ctx context.Context, request dto.FacilityRequest) entity.Facility {
	return entity.Facility{
		Name: request.Name,
		Icon: request.Icon,
	}
}

func (m FacilityMapper) MapEntityToResponse(ctx context.Context, facilityEntity entity.Facility) dto.FacilityResponse {
	return dto.FacilityResponse{
		Id:        facilityEntity.Id,
		Name:      facilityEntity.Name,
		Icon:      facilityEntity.Icon,
		CreatedAt: facilityEntity.CreatedAt,
		UpdatedAt: facilityEntity.UpdatedAt,
		DeletedAt: facilityEntity.DeletedAt,
	}
}

func (m FacilityMapper) MapEntitiesToResponses(ctx context.Context, facilityEntities []entity.Facility) []dto.FacilityResponse {
	facilityResponses := make([]dto.FacilityResponse, len(facilityEntities))

	for i, facilityEntity := range facilityEntities {
		facilityResponses[i] = m.MapEntityToResponse(ctx, facilityEntity)
	}

	return facilityResponses
}
