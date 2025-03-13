package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type BusMapper struct {
}

func NewBusMapper() BusMapper {
	return BusMapper{}
}

func (BusMapper) MapRequestToEntity(ctx context.Context, request dto.BusRequest) entity.Bus {
	return entity.Bus{
		Name: request.Name,
		Seat: request.Seat,
	}
}

func (m BusMapper) MapEntityToResponse(ctx context.Context, busEntity entity.Bus) dto.BusResponse {
	return dto.BusResponse{
		Id:        busEntity.Id,
		Name:      busEntity.Name,
		Seat:      busEntity.Seat,
		CreatedAt: busEntity.CreatedAt,
		UpdatedAt: busEntity.UpdatedAt,
		DeletedAt: busEntity.DeletedAt,
	}
}

func (m BusMapper) MapEntitiesToResponses(ctx context.Context, busEntities []entity.Bus) []dto.BusResponse {
	busResponses := make([]dto.BusResponse, len(busEntities))

	for i, busEntity := range busEntities {
		busResponses[i] = m.MapEntityToResponse(ctx, busEntity)
	}

	return busResponses
}
