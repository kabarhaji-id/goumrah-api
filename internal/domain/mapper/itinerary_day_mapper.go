package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ItineraryDayMapper struct {
	imageMapper           ImageMapper
	itineraryWidgetMapper ItineraryWidgetMapper
}

func NewItineraryDayMapper(
	imageMapper ImageMapper,
	itineraryWidgetMapper ItineraryWidgetMapper,
) ItineraryDayMapper {
	return ItineraryDayMapper{imageMapper, itineraryWidgetMapper}
}

func (m ItineraryDayMapper) MapEntityToResponse(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	hotelRepository repository.HotelRepository,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryDayEntity entity.ItineraryDay,
) (dto.ItineraryDayResponse, error) {
	itineraryWidgetEntities := []entity.ItineraryWidget{}
	if itineraryDayEntity.WidgetId.Valid {
		itineraryWidgetEntity, err := itineraryWidgetRepository.FindById(ctx, itineraryDayEntity.WidgetId.Int64)
		if err != nil {
			return dto.ItineraryDayResponse{}, err
		}
		itineraryWidgetEntities = append(itineraryWidgetEntities, itineraryWidgetEntity)
		for itineraryWidgetEntity.NextId.Valid {
			itineraryWidgetEntity, err = itineraryWidgetRepository.FindById(ctx, itineraryWidgetEntity.NextId.Int64)
			if err != nil {
				return dto.ItineraryDayResponse{}, err
			}
			itineraryWidgetEntities = append(itineraryWidgetEntities, itineraryWidgetEntity)
		}
	}

	itineraryWidgetResponses, err := m.itineraryWidgetMapper.MapEntitiesToResponses(
		ctx,
		imageRepository,
		hotelRepository,
		itineraryWidgetActivityRepository,
		itineraryWidgetHotelRepository,
		itineraryWidgetInformationRepository,
		itineraryWidgetTransportRepository,
		itineraryWidgetRecommendationRepository,
		itineraryWidgetEntities,
	)
	if err != nil {
		return dto.ItineraryDayResponse{}, err
	}

	return dto.ItineraryDayResponse{
		Title:       itineraryDayEntity.Title,
		Description: itineraryDayEntity.Description,
		Widgets:     itineraryWidgetResponses,
	}, nil
}

func (m ItineraryDayMapper) MapEntitiesToResponses(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	hotelRepository repository.HotelRepository,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryDayEntities []entity.ItineraryDay,
) ([]dto.ItineraryDayResponse, error) {
	itineraryDayResponses := make([]dto.ItineraryDayResponse, len(itineraryDayEntities))
	var err error

	for i, itineraryDayEntity := range itineraryDayEntities {
		itineraryDayResponses[i], err = m.MapEntityToResponse(
			ctx,
			imageRepository,
			hotelRepository,
			itineraryWidgetRepository,
			itineraryWidgetActivityRepository,
			itineraryWidgetHotelRepository,
			itineraryWidgetInformationRepository,
			itineraryWidgetTransportRepository,
			itineraryWidgetRecommendationRepository,
			itineraryDayEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return itineraryDayResponses, nil
}
