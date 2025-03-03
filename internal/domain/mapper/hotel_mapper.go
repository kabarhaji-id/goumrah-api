package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
)

type HotelMapper struct {
}

func NewHotelMapper() HotelMapper {
	return HotelMapper{}
}

func (HotelMapper) MapRequestToEntity(ctx context.Context, request dto.HotelRequest) entity.Hotel {
	return entity.Hotel{
		Name:        request.Name,
		Rating:      request.Rating,
		Map:         request.Map,
		Address:     request.Address,
		Distance:    request.Distance,
		Review:      request.Review,
		Description: request.Description,
		Location:    request.Location,
		Slug:        sluger.Slug(request.Name),
	}
}

func (m HotelMapper) MapEntityToResponse(ctx context.Context, hotelEntity entity.Hotel) dto.HotelResponse {
	return dto.HotelResponse{
		Id:          hotelEntity.Id,
		Name:        hotelEntity.Name,
		Rating:      hotelEntity.Rating,
		Map:         hotelEntity.Map,
		Address:     hotelEntity.Address,
		Distance:    hotelEntity.Distance,
		Review:      hotelEntity.Review,
		Description: hotelEntity.Description,
		Location:    hotelEntity.Location,
		Slug:        hotelEntity.Slug,
		CreatedAt:   hotelEntity.CreatedAt,
		UpdatedAt:   hotelEntity.UpdatedAt,
		DeletedAt:   hotelEntity.DeletedAt,
	}
}

func (m HotelMapper) MapEntitiesToResponses(ctx context.Context, hotelEntities []entity.Hotel) []dto.HotelResponse {
	hotelResponses := make([]dto.HotelResponse, len(hotelEntities))

	for i, hotelEntity := range hotelEntities {
		hotelResponses[i] = m.MapEntityToResponse(ctx, hotelEntity)
	}

	return hotelResponses
}
