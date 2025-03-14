package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
)

type HotelMapper struct {
	imageMapper ImageMapper
}

func NewHotelMapper(imageMapper ImageMapper) HotelMapper {
	return HotelMapper{imageMapper}
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

func (m HotelMapper) MapEntityToResponse(
	ctx context.Context,
	hotelRepository repository.HotelRepository,
	hotelEntity entity.Hotel,
) (dto.HotelResponse, error) {
	imageEntities, err := hotelRepository.FindImages(ctx, hotelEntity.Id)
	if err != nil {
		return dto.HotelResponse{}, err
	}
	imageResponses := m.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

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
		Images:      imageResponses,
		CreatedAt:   hotelEntity.CreatedAt,
		UpdatedAt:   hotelEntity.UpdatedAt,
		DeletedAt:   hotelEntity.DeletedAt,
	}, nil
}

func (m HotelMapper) MapEntitiesToResponses(
	ctx context.Context,
	hotelRepository repository.HotelRepository,
	hotelEntities []entity.Hotel,
) ([]dto.HotelResponse, error) {
	hotelResponses := make([]dto.HotelResponse, len(hotelEntities))
	var err error

	for i, hotelEntity := range hotelEntities {
		hotelResponses[i], err = m.MapEntityToResponse(ctx, hotelRepository, hotelEntity)
		if err != nil {
			return nil, err
		}
	}

	return hotelResponses, nil
}
