package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirportMapper struct {
}

func NewAirportMapper() AirportMapper {
	return AirportMapper{}
}

func (AirportMapper) MapRequestToEntity(ctx context.Context, request dto.AirportRequest) entity.Airport {
	return entity.Airport{
		City: request.City,
		Name: request.Name,
		Code: request.Code,
	}
}

func (m AirportMapper) MapEntityToResponse(ctx context.Context, airportEntity entity.Airport) dto.AirportResponse {
	return dto.AirportResponse{
		Id:        airportEntity.Id,
		City:      airportEntity.City,
		Name:      airportEntity.Name,
		Code:      airportEntity.Code,
		CreatedAt: airportEntity.CreatedAt,
		UpdatedAt: airportEntity.UpdatedAt,
		DeletedAt: airportEntity.DeletedAt,
	}
}

func (m AirportMapper) MapEntitiesToResponses(ctx context.Context, airportEntities []entity.Airport) []dto.AirportResponse {
	airportResponses := make([]dto.AirportResponse, len(airportEntities))

	for i, airportEntity := range airportEntities {
		airportResponses[i] = m.MapEntityToResponse(ctx, airportEntity)
	}

	return airportResponses
}
