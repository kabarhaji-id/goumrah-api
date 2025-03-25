package mapper

import (
	"context"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionMapper struct {
	embarkationMapper EmbarkationMapper
	guideMapper       GuideMapper
	flightMapper      FlightMapper
	busMapper         BusMapper
	itineraryMapper   ItineraryMapper
}

func NewPackageSessionMapper(embarkationMapper EmbarkationMapper, guideMapper GuideMapper, flightMapper FlightMapper, busMapper BusMapper, itineraryMapper ItineraryMapper) PackageSessionMapper {
	return PackageSessionMapper{
		embarkationMapper, guideMapper, flightMapper, busMapper, itineraryMapper,
	}
}

func (PackageSessionMapper) MapRequestToEntity(ctx context.Context, request dto.PackageSessionRequest) entity.PackageSession {
	departureDate, _ := time.Parse("02/01/2006", request.DepartureDate)

	return entity.PackageSession{
		EmbarkationId:    request.Embarkation,
		DepartureDate:    departureDate,
		Quota:            request.Quota,
		DoublePrice:      request.DoublePrice,
		DoubleFinalPrice: request.QuadFinalPrice,
		TriplePrice:      request.TriplePrice,
		TripleFinalPrice: request.TripleFinalPrice,
		QuadPrice:        request.QuadPrice,
		QuadFinalPrice:   request.QuadFinalPrice,
		InfantPrice:      request.InfantPrice,
		InfantFinalPrice: request.InfantFinalPrice,
		BusId:            request.Bus,
	}
}

func (m PackageSessionMapper) MapEntityToResponse(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	imageRepository repository.ImageRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	airportRepository repository.AirportRepository,
	busRepository repository.BusRepository,
	hotelRepository repository.HotelRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	packageSessionEntity entity.PackageSession,
) (dto.PackageSessionResponse, error) {
	embarkationEntity, err := embarkationRepository.FindById(ctx, packageSessionEntity.EmbarkationId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	embarkationResponse := m.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	departureFlightRouteEntity, err := flightRouteRepository.FindById(ctx, packageSessionEntity.DepartureFlightRouteId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	departureFlightEntities := []entity.Flight{}
	departureFlightEntity, err := flightRepository.FindById(ctx, departureFlightRouteEntity.FlightId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	departureFlightEntities = append(departureFlightEntities, departureFlightEntity)
	for departureFlightRouteEntity.NextId.Valid {
		departureFlightRouteEntity, err = flightRouteRepository.FindById(ctx, departureFlightRouteEntity.NextId.Int64)
		if err != nil {
			return dto.PackageSessionResponse{}, err
		}
		departureFlightEntity, err = flightRepository.FindById(ctx, departureFlightRouteEntity.FlightId)
		if err != nil {
			return dto.PackageSessionResponse{}, err
		}
		departureFlightEntities = append(departureFlightEntities, departureFlightEntity)
	}
	departureFlightResponses, err := m.flightMapper.MapEntitiesToResponses(
		ctx,
		imageRepository,
		airlineRepository,
		airportRepository,
		departureFlightEntities,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	returnFlightRouteEntity, err := flightRouteRepository.FindById(ctx, packageSessionEntity.ReturnFlightRouteId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	returnFlightEntities := []entity.Flight{}
	returnFlightEntity, err := flightRepository.FindById(ctx, returnFlightRouteEntity.FlightId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	returnFlightEntities = append(returnFlightEntities, returnFlightEntity)
	for returnFlightRouteEntity.NextId.Valid {
		returnFlightRouteEntity, err = flightRouteRepository.FindById(ctx, returnFlightRouteEntity.NextId.Int64)
		if err != nil {
			return dto.PackageSessionResponse{}, err
		}
		returnFlightEntity, err = flightRepository.FindById(ctx, returnFlightRouteEntity.FlightId)
		if err != nil {
			return dto.PackageSessionResponse{}, err
		}
		returnFlightEntities = append(returnFlightEntities, returnFlightEntity)
	}
	returnFlightResponses, err := m.flightMapper.MapEntitiesToResponses(
		ctx,
		imageRepository,
		airlineRepository,
		airportRepository,
		returnFlightEntities,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	guideEntities, err := packageSessionRepository.FindGuides(ctx, packageSessionEntity.Id)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	guideResponses, err := m.guideMapper.MapEntitiesToResponses(ctx, imageRepository, guideEntities)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	busEntity, err := busRepository.FindById(ctx, packageSessionEntity.BusId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	busResponse := m.busMapper.MapEntityToResponse(ctx, busEntity)

	itineraryEntity, err := itineraryRepository.FindById(ctx, packageSessionEntity.ItineraryId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	itineraryEntities := []entity.Itinerary{itineraryEntity}
	for itineraryEntity.NextId.Valid {
		itineraryEntity, err = itineraryRepository.FindById(ctx, itineraryEntity.NextId.Int64)
		if err != nil {
			return dto.PackageSessionResponse{}, err
		}
		itineraryEntities = append(itineraryEntities, itineraryEntity)
	}
	itineraryResponses, err := m.itineraryMapper.MapEntitiesToResponses(
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
		itineraryEntities,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return dto.PackageSessionResponse{
		Id:               packageSessionEntity.Id,
		Package:          packageSessionEntity.PackageId,
		Embarkation:      embarkationResponse,
		DepartureDate:    packageSessionEntity.DepartureDate,
		Quota:            packageSessionEntity.Quota,
		DoublePrice:      packageSessionEntity.DoublePrice,
		DoubleFinalPrice: packageSessionEntity.DoubleFinalPrice,
		TriplePrice:      packageSessionEntity.TriplePrice,
		TripleFinalPrice: packageSessionEntity.TripleFinalPrice,
		QuadPrice:        packageSessionEntity.QuadPrice,
		QuadFinalPrice:   packageSessionEntity.QuadFinalPrice,
		InfantPrice:      packageSessionEntity.InfantPrice,
		InfantFinalPrice: packageSessionEntity.InfantFinalPrice,
		DepartureFlights: departureFlightResponses,
		ReturnFlights:    returnFlightResponses,
		Guides:           guideResponses,
		Bus:              busResponse,
		Itineraries:      itineraryResponses,
		CreatedAt:        packageSessionEntity.CreatedAt,
		UpdatedAt:        packageSessionEntity.UpdatedAt,
		DeletedAt:        packageSessionEntity.DeletedAt,
	}, nil
}

func (m PackageSessionMapper) MapEntityToListResponse(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	flightRouteRepository repository.FlightRouteRepository,
	packageSessionEntity entity.PackageSession,
) (dto.PackageSessionListResponse, error) {
	embarkationEntity, err := embarkationRepository.FindById(ctx, packageSessionEntity.EmbarkationId)
	if err != nil {
		return dto.PackageSessionListResponse{}, err
	}
	embarkationResponse := m.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	departureFlightRouteEntity, err := flightRouteRepository.FindById(ctx, packageSessionEntity.DepartureFlightRouteId)
	if err != nil {
		return dto.PackageSessionListResponse{}, err
	}
	departureFlightIds := []int64{}
	departureFlightIds = append(departureFlightIds, departureFlightRouteEntity.FlightId)
	for departureFlightRouteEntity.NextId.Valid {
		departureFlightRouteEntity, err = flightRouteRepository.FindById(ctx, departureFlightRouteEntity.NextId.Int64)
		if err != nil {
			return dto.PackageSessionListResponse{}, err
		}
		departureFlightIds = append(departureFlightIds, departureFlightRouteEntity.FlightId)
	}

	returnFlightRouteEntity, err := flightRouteRepository.FindById(ctx, packageSessionEntity.ReturnFlightRouteId)
	if err != nil {
		return dto.PackageSessionListResponse{}, err
	}
	returnFlightIds := []int64{}
	returnFlightIds = append(returnFlightIds, returnFlightRouteEntity.FlightId)
	for returnFlightRouteEntity.NextId.Valid {
		returnFlightRouteEntity, err = flightRouteRepository.FindById(ctx, returnFlightRouteEntity.NextId.Int64)
		if err != nil {
			return dto.PackageSessionListResponse{}, err
		}
		returnFlightIds = append(returnFlightIds, returnFlightRouteEntity.FlightId)
	}

	guideIds, err := packageSessionRepository.FindGuideIds(ctx, packageSessionEntity.Id)
	if err != nil {
		return dto.PackageSessionListResponse{}, err
	}

	return dto.PackageSessionListResponse{
		Id:               packageSessionEntity.Id,
		Package:          packageSessionEntity.PackageId,
		Embarkation:      embarkationResponse,
		DepartureDate:    packageSessionEntity.DepartureDate,
		DepartureFlights: departureFlightIds,
		ReturnFlights:    returnFlightIds,
		Guides:           guideIds,
		Bus:              packageSessionEntity.BusId,
		CreatedAt:        packageSessionEntity.CreatedAt,
		UpdatedAt:        packageSessionEntity.UpdatedAt,
		DeletedAt:        packageSessionEntity.DeletedAt,
	}, nil
}

func (m PackageSessionMapper) MapEntitiesToResponses(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	imageRepository repository.ImageRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	airportRepository repository.AirportRepository,
	busRepository repository.BusRepository,
	hotelRepository repository.HotelRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryDayRepository repository.ItineraryDayRepository,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	packageSessionEntities []entity.PackageSession,
) ([]dto.PackageSessionResponse, error) {
	packageSessionResponses := make([]dto.PackageSessionResponse, len(packageSessionEntities))
	var err error

	for i, packageSessionEntity := range packageSessionEntities {
		packageSessionResponses[i], err = m.MapEntityToResponse(
			ctx,
			embarkationRepository,
			packageSessionRepository,
			imageRepository,
			flightRouteRepository,
			flightRepository,
			airlineRepository,
			airportRepository,
			busRepository,
			hotelRepository,
			itineraryRepository,
			itineraryDayRepository,
			itineraryWidgetRepository,
			itineraryWidgetActivityRepository,
			itineraryWidgetHotelRepository,
			itineraryWidgetInformationRepository,
			itineraryWidgetTransportRepository,
			itineraryWidgetRecommendationRepository,
			packageSessionEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return packageSessionResponses, nil
}

func (m PackageSessionMapper) MapEntitiesToListResponses(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	flightRouteRepository repository.FlightRouteRepository,
	packageSessionEntities []entity.PackageSession,
) ([]dto.PackageSessionListResponse, error) {
	packageSessionListResponses := make([]dto.PackageSessionListResponse, len(packageSessionEntities))
	var err error

	for i, packageSessionEntity := range packageSessionEntities {
		packageSessionListResponses[i], err = m.MapEntityToListResponse(
			ctx,
			embarkationRepository,
			packageSessionRepository,
			flightRouteRepository,
			packageSessionEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return packageSessionListResponses, nil
}
