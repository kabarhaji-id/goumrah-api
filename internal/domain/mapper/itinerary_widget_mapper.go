package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ItineraryWidgetMapper struct {
	imageMapper ImageMapper
	hotelMapper HotelMapper
}

func NewItineraryWidgetMapper(
	imageMapper ImageMapper,
	hotelMapper HotelMapper,
) ItineraryWidgetMapper {
	return ItineraryWidgetMapper{imageMapper, hotelMapper}
}

func (m ItineraryWidgetMapper) MapEntityToResponse(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	hotelRepository repository.HotelRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryWidgetEntity entity.ItineraryWidget,
) (dto.ItineraryWidgetResponse, error) {
	if itineraryWidgetEntity.ActivityId.Valid {
		itineraryWidgetActivity, err := itineraryWidgetActivityRepository.FindById(ctx, itineraryWidgetEntity.ActivityId.Int64)
		if err != nil {
			return nil, err
		}

		imageEntities, err := itineraryWidgetActivityRepository.FindImages(ctx, itineraryWidgetActivity.Id)
		if err != nil {
			return nil, err
		}
		imageResponses := m.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

		return dto.ItineraryWidgetActivityResponse{
			Title:       itineraryWidgetActivity.Title,
			Description: itineraryWidgetActivity.Description,
			Images:      imageResponses,
		}, nil
	}

	if itineraryWidgetEntity.HotelId.Valid {
		itineraryWidgetHotel, err := itineraryWidgetHotelRepository.FindById(ctx, itineraryWidgetEntity.HotelId.Int64)
		if err != nil {
			return nil, err
		}

		hotelEntity, err := hotelRepository.FindById(ctx, itineraryWidgetHotel.HotelId)
		if err != nil {
			return nil, err
		}
		hotelResponse, err := m.hotelMapper.MapEntityToResponse(ctx, hotelRepository, hotelEntity)
		if err != nil {
			return nil, err
		}

		return dto.ItineraryWidgetHotelResponse{
			Hotel: hotelResponse,
		}, nil
	}

	if itineraryWidgetEntity.InformationId.Valid {
		itineraryWidgetInformation, err := itineraryWidgetInformationRepository.FindById(ctx, itineraryWidgetEntity.InformationId.Int64)
		if err != nil {
			return nil, err
		}

		return dto.ItineraryWidgetInformationResponse{
			Description: itineraryWidgetInformation.Description,
		}, nil
	}

	if itineraryWidgetEntity.TransportId.Valid {
		itineraryWidgetTransport, err := itineraryWidgetTransportRepository.FindById(ctx, itineraryWidgetEntity.TransportId.Int64)
		if err != nil {
			return nil, err
		}

		return dto.ItineraryWidgetTransportResponse{
			Transportation: itineraryWidgetTransport.Transportation,
			From:           itineraryWidgetTransport.From,
			To:             itineraryWidgetTransport.To,
		}, nil
	}

	if itineraryWidgetEntity.RecommendationId.Valid {
		itineraryWidgetRecommendation, err := itineraryWidgetRecommendationRepository.FindById(ctx, itineraryWidgetEntity.RecommendationId.Int64)
		if err != nil {
			return nil, err
		}

		imageEntities, err := itineraryWidgetRecommendationRepository.FindImages(ctx, itineraryWidgetRecommendation.Id)
		if err != nil {
			return nil, err
		}
		imageResponses := m.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

		return dto.ItineraryWidgetRecommendationResponse{
			Description: itineraryWidgetRecommendation.Description,
			Images:      imageResponses,
		}, nil
	}

	return nil, nil
}

func (m ItineraryWidgetMapper) MapEntitiesToResponses(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	hotelRepository repository.HotelRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryWidgetEntities []entity.ItineraryWidget,
) ([]dto.ItineraryWidgetResponse, error) {
	itineraryWidgetResponses := make([]dto.ItineraryWidgetResponse, len(itineraryWidgetEntities))
	var err error

	for i, itineraryWidgetEntity := range itineraryWidgetEntities {
		itineraryWidgetResponses[i], err = m.MapEntityToResponse(
			ctx,
			imageRepository,
			hotelRepository,
			itineraryWidgetActivityRepository,
			itineraryWidgetHotelRepository,
			itineraryWidgetInformationRepository,
			itineraryWidgetTransportRepository,
			itineraryWidgetRecommendationRepository,
			itineraryWidgetEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return itineraryWidgetResponses, nil
}
