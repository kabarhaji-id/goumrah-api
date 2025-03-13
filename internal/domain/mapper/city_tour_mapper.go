package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type CityTourMapper struct {
}

func NewCityTourMapper() CityTourMapper {
	return CityTourMapper{}
}

func (CityTourMapper) MapRequestToEntity(ctx context.Context, request dto.CityTourRequest) entity.CityTour {
	return entity.CityTour{
		Name:        request.Name,
		City:        request.City,
		Description: request.Description,
	}
}

func (m CityTourMapper) MapEntityToResponse(ctx context.Context, cityTourEntity entity.CityTour) dto.CityTourResponse {
	return dto.CityTourResponse{
		Id:          cityTourEntity.Id,
		Name:        cityTourEntity.Name,
		City:        cityTourEntity.City,
		Description: cityTourEntity.Description,
		CreatedAt:   cityTourEntity.CreatedAt,
		UpdatedAt:   cityTourEntity.UpdatedAt,
		DeletedAt:   cityTourEntity.DeletedAt,
	}
}

func (m CityTourMapper) MapEntitiesToResponses(ctx context.Context, cityTourEntities []entity.CityTour) []dto.CityTourResponse {
	cityTourResponses := make([]dto.CityTourResponse, len(cityTourEntities))

	for i, cityTourEntity := range cityTourEntities {
		cityTourResponses[i] = m.MapEntityToResponse(ctx, cityTourEntity)
	}

	return cityTourResponses
}
