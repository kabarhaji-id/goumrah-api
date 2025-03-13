package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ItineraryMapper struct {
	imageMapper        ImageMapper
	itineraryDayMapper ItineraryDayMapper
}

func NewItineraryMapper(
	imageMapper ImageMapper,
	itineraryDayMapper ItineraryDayMapper,
) ItineraryMapper {
	return ItineraryMapper{imageMapper, itineraryDayMapper}
}

func (m ItineraryMapper) MapEntityToResponse(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	hotelRepository repository.HotelRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryEntity entity.Itinerary,
) (dto.ItineraryResponse, error) {
	imageEntities, err := itineraryRepository.FindImages(ctx, itineraryEntity.Id)
	if err != nil {
		return dto.ItineraryResponse{}, err
	}
	imageResponses := m.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

	itineraryDayEntity, err := itineraryDayRepository.FindById(ctx, itineraryEntity.DayId)
	if err != nil {
		return dto.ItineraryResponse{}, err
	}
	itineraryDayEntities := []entity.ItineraryDay{itineraryDayEntity}
	for itineraryDayEntity.NextId.Valid {
		itineraryDayEntity, err = itineraryDayRepository.FindById(ctx, itineraryDayEntity.NextId.Int64)
		if err != nil {
			return dto.ItineraryResponse{}, err
		}
		itineraryDayEntities = append(itineraryDayEntities, itineraryDayEntity)
	}
	itineraryDayResponses, err := m.itineraryDayMapper.MapEntitiesToResponses(
		ctx,
		imageRepository,
		hotelRepository,
		itineraryWidgetRepository,
		itineraryWidgetActivityRepository,
		itineraryWidgetHotelRepository,
		itineraryWidgetInformationRepository,
		itineraryWidgetTransportRepository,
		itineraryWidgetRecommendationRepository,
		itineraryDayEntities,
	)
	if err != nil {
		return dto.ItineraryResponse{}, err
	}

	return dto.ItineraryResponse{
		Id:        itineraryEntity.Id,
		City:      itineraryEntity.City,
		Images:    imageResponses,
		Days:      itineraryDayResponses,
		CreatedAt: itineraryEntity.CreatedAt,
		UpdatedAt: itineraryEntity.UpdatedAt,
		DeletedAt: itineraryEntity.DeletedAt,
	}, nil
}

func (m ItineraryMapper) MapEntitiesToResponses(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	hotelRepository repository.HotelRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryEntities []entity.Itinerary,
) ([]dto.ItineraryResponse, error) {
	itineraryResponses := make([]dto.ItineraryResponse, len(itineraryEntities))
	var err error

	for i, itineraryEntity := range itineraryEntities {
		itineraryResponses[i], err = m.MapEntityToResponse(
			ctx,
			imageRepository,
			hotelRepository,
			itineraryRepository,
			itineraryDayRepository,
			itineraryWidgetRepository,
			itineraryWidgetActivityRepository,
			itineraryWidgetHotelRepository,
			itineraryWidgetInformationRepository,
			itineraryWidgetTransportRepository,
			itineraryWidgetRecommendationRepository,
			itineraryEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return itineraryResponses, nil
}
