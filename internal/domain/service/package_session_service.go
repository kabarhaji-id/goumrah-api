package service

import (
	"context"
	"slices"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	serviceport "github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type packageSessionServiceImpl struct {
	packageSessionRepository repository.PackageSessionRepository
	packageSessionValidator  validator.PackageSessionValidator
	packageSessionMapper     mapper.PackageSessionMapper

	embarkationRepository repository.EmbarkationRepository

	imageRepository repository.ImageRepository

	flightRouteRepository repository.FlightRouteRepository

	flightRepository repository.FlightRepository

	airlineRepository repository.AirlineRepository

	airportRepository repository.AirportRepository

	busRepository repository.BusRepository

	hotelRepository repository.HotelRepository

	itineraryRepository repository.ItineraryRepository
	itineraryMapper     mapper.ItineraryMapper

	itineraryDayRepository repository.ItineraryDayRepository
	itineraryDayMapper     mapper.ItineraryDayMapper

	itineraryWidgetRepository               repository.ItineraryWidgetRepository
	itineraryWidgetActivityRepository       repository.ItineraryWidgetActivityRepository
	itineraryWidgetHotelRepository          repository.ItineraryWidgetHotelRepository
	itineraryWidgetInformationRepository    repository.ItineraryWidgetInformationRepository
	itineraryWidgetTransportRepository      repository.ItineraryWidgetTransportRepository
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository
	itineraryWidgetMapper                   mapper.ItineraryWidgetMapper

	unitOfWork repository.UnitOfWork
}

func NewPackageSessionService(
	packageSessionRepository repository.PackageSessionRepository,
	packageSessionValidator validator.PackageSessionValidator,
	packageSessionMapper mapper.PackageSessionMapper,
	embarkationRepository repository.EmbarkationRepository,
	imageRepository repository.ImageRepository,
	flightRouteRepository repository.FlightRouteRepository,
	flightRepository repository.FlightRepository,
	airlineRepository repository.AirlineRepository,
	airportRepository repository.AirportRepository,
	busRepository repository.BusRepository,
	hotelRepository repository.HotelRepository,
	itineraryRepository repository.ItineraryRepository,
	itineraryMapper mapper.ItineraryMapper,
	itineraryDayRepository repository.ItineraryDayRepository,
	itineraryDayMapper mapper.ItineraryDayMapper,
	itineraryWidgetRepository repository.ItineraryWidgetRepository,
	itineraryWidgetActivityRepository repository.ItineraryWidgetActivityRepository,
	itineraryWidgetHotelRepository repository.ItineraryWidgetHotelRepository,
	itineraryWidgetInformationRepository repository.ItineraryWidgetInformationRepository,
	itineraryWidgetTransportRepository repository.ItineraryWidgetTransportRepository,
	itineraryWidgetRecommendationRepository repository.ItineraryWidgetRecommendationRepository,
	itineraryWidgetMapper mapper.ItineraryWidgetMapper,
	unitOfWork repository.UnitOfWork,
) serviceport.PackageSessionService {
	return packageSessionServiceImpl{
		packageSessionRepository,
		packageSessionValidator,
		packageSessionMapper,
		embarkationRepository,
		imageRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		airportRepository,
		busRepository,
		hotelRepository,
		itineraryRepository,
		itineraryMapper,
		itineraryDayRepository,
		itineraryDayMapper,
		itineraryWidgetRepository,
		itineraryWidgetActivityRepository,
		itineraryWidgetHotelRepository,
		itineraryWidgetInformationRepository,
		itineraryWidgetTransportRepository,
		itineraryWidgetRecommendationRepository,
		itineraryWidgetMapper,
		unitOfWork,
	}
}

func (s packageSessionServiceImpl) CreatePackageSession(ctx context.Context, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error) {
	// Validate request
	if err := s.packageSessionValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map request into entity
	packageSessionEntity := s.packageSessionMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.PackageSessionResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create package session repository
		packageSessionRepository := factory.NewPackageSessionRepository()

		// Create embarkation repository
		embarkationRepository := factory.NewEmbarkationRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create flight route repository
		flightRouteRepository := factory.NewFlightRouteRepository()

		// Create flight repository
		flightRepository := factory.NewFlightRepository()

		// Create airline repository
		airlineRepository := factory.NewAirlineRepository()

		// Create airport repository
		airportRepository := factory.NewAirportRepository()

		// Create bus repository
		busRepository := factory.NewBusRepository()

		// Create hotel repository
		hotelRepository := factory.NewHotelRepository()

		// Create itinerary repository
		itineraryRepository := factory.NewItineraryRepository()

		// Create itinerary day repository
		itineraryDayRepository := factory.NewItineraryDayRepository()

		// Create itinerary widget repository
		itineraryWidgetRepository := factory.NewItineraryWidgetRepository()

		// Create itinerary widget activity repository
		itineraryWidgetActivityRepository := factory.NewItineraryWidgetActivityRepository()

		// Create itinerary widget hotel repository
		itineraryWidgetHotelRepository := factory.NewItineraryWidgetHotelRepository()

		// Create itinerary widget information repository
		itineraryWidgetInformationRepository := factory.NewItineraryWidgetInformationRepository()

		// Create itinerary widget transport repository
		itineraryWidgetTransportRepository := factory.NewItineraryWidgetTransportRepository()

		// Create itinerary widget recommendation repository
		itineraryWidgetRecommendationRepository := factory.NewItineraryWidgetRecommendationRepository()

		// Create departure flight route
		for _, departureFlight := range slices.Backward(request.DepartureFlights) {
			flightRoute, err := flightRouteRepository.Create(
				ctx,
				entity.FlightRoute{
					FlightId: departureFlight,
					NextId:   null.NewInt(packageSessionEntity.DepartureFlightRouteId, packageSessionEntity.DepartureFlightRouteId != 0),
				},
			)
			if err != nil {
				return err
			}

			packageSessionEntity.DepartureFlightRouteId = flightRoute.Id
		}

		// Create return flight route
		for _, returnFlight := range slices.Backward(request.ReturnFlights) {
			flightRoute, err := flightRouteRepository.Create(
				ctx,
				entity.FlightRoute{
					FlightId: returnFlight,
					NextId:   null.NewInt(packageSessionEntity.ReturnFlightRouteId, packageSessionEntity.ReturnFlightRouteId != 0),
				},
			)
			if err != nil {
				return err
			}

			packageSessionEntity.ReturnFlightRouteId = flightRoute.Id
		}

		// Create itinerary with repository
		itineraryEntity := null.NewValue(entity.Itinerary{}, false)
		for _, itineraryRequest := range request.Itineraries {
			var err error

			itineraryEntity = null.ValueFrom(entity.Itinerary{
				City: itineraryRequest.City,
				NextId: null.NewInt(
					itineraryEntity.V.Id,
					itineraryEntity.Valid,
				),
			})

			itineraryDayEntity := null.NewValue(entity.ItineraryDay{}, false)
			for _, itineraryDayRequest := range itineraryRequest.Days {
				itineraryDayEntity = null.ValueFrom(entity.ItineraryDay{
					Title:       itineraryDayRequest.Title,
					Description: itineraryDayRequest.Description,
					NextId: null.NewInt(
						itineraryDayEntity.V.Id,
						itineraryDayEntity.Valid,
					),
				})

				itineraryWidgetEntity := null.NewValue(entity.ItineraryWidget{}, false)
				for _, itineraryWidgetRequest := range itineraryDayRequest.Widgets {
					itineraryWidgetEntity = null.ValueFrom(entity.ItineraryWidget{
						NextId: null.NewInt(
							itineraryWidgetEntity.V.Id,
							itineraryWidgetEntity.Valid,
						),
					})

					switch itineraryWidgetRequest.Type() {
					case "Activity":
						itineraryWidgetActivityRequest := itineraryWidgetRequest.(dto.ItineraryWidgetActivityRequest)
						itineraryWidgetActivityEntity, err := itineraryWidgetActivityRepository.Create(
							ctx,
							entity.ItineraryWidgetActivity{
								Title:       itineraryWidgetActivityRequest.Title,
								Description: itineraryWidgetActivityRequest.Description,
							},
						)
						if err != nil {
							return err
						}

						if _, err = itineraryWidgetActivityRepository.AttachImages(
							ctx,
							itineraryWidgetActivityEntity.Id,
							itineraryWidgetActivityRequest.Images,
						); err != nil {
							return err
						}

						itineraryWidgetEntity.V.ActivityId = null.IntFrom(itineraryWidgetActivityEntity.Id)
					case "Hotel":
						itineraryWidgetHotelRequest := itineraryWidgetRequest.(dto.ItineraryWidgetHotelRequest)
						itineraryWidgetHotelEntity, err := itineraryWidgetHotelRepository.Create(
							ctx,
							entity.ItineraryWidgetHotel{
								HotelId: itineraryWidgetHotelRequest.Hotel,
							},
						)
						if err != nil {
							return err
						}

						itineraryWidgetEntity.V.HotelId = null.IntFrom(itineraryWidgetHotelEntity.Id)
					case "Information":
						itineraryWidgetInformationRequest := itineraryWidgetRequest.(dto.ItineraryWidgetInformationRequest)
						itineraryWidgetInformationEntity, err := itineraryWidgetInformationRepository.Create(
							ctx,
							entity.ItineraryWidgetInformation{
								Description: itineraryWidgetInformationRequest.Description,
							},
						)
						if err != nil {
							return err
						}

						itineraryWidgetEntity.V.InformationId = null.IntFrom(itineraryWidgetInformationEntity.Id)
					case "Transport":
						itineraryWidgetTransportRequest := itineraryWidgetRequest.(dto.ItineraryWidgetTransportRequest)
						itineraryWidgetTransportEntity, err := itineraryWidgetTransportRepository.Create(
							ctx,
							entity.ItineraryWidgetTransport{
								Transportation: itineraryWidgetTransportRequest.Transportation,
								From:           itineraryWidgetTransportRequest.From,
								To:             itineraryWidgetTransportRequest.To,
							},
						)
						if err != nil {
							return err
						}

						itineraryWidgetEntity.V.TransportId = null.IntFrom(itineraryWidgetTransportEntity.Id)
					case "Recommendation":
						itineraryWidgetRecommendationRequest := itineraryWidgetRequest.(dto.ItineraryWidgetRecommendationRequest)
						itineraryWidgetRecommendationEntity, err := itineraryWidgetRecommendationRepository.Create(
							ctx,
							entity.ItineraryWidgetRecommendation{
								Description: itineraryWidgetRecommendationRequest.Description,
							},
						)
						if err != nil {
							return err
						}

						if _, err = itineraryWidgetRecommendationRepository.AttachImages(
							ctx,
							itineraryWidgetRecommendationEntity.Id,
							itineraryWidgetRecommendationRequest.Images,
						); err != nil {
							return err
						}

						itineraryWidgetEntity.V.RecommendationId = null.IntFrom(itineraryWidgetRecommendationEntity.Id)
					}

					itineraryWidgetEntity.V, err = itineraryWidgetRepository.Create(ctx, itineraryWidgetEntity.V)
					if err != nil {
						return err
					}
				}
				itineraryDayEntity.V.WidgetId = null.NewInt(itineraryWidgetEntity.V.Id, itineraryWidgetEntity.Valid)

				itineraryDayEntity.V, err = itineraryDayRepository.Create(ctx, itineraryDayEntity.V)
				if err != nil {
					return err
				}
			}
			itineraryEntity.V.DayId = itineraryDayEntity.V.Id

			itineraryEntity.V, err = itineraryRepository.Create(ctx, itineraryEntity.V)
			if err != nil {
				return err
			}

			if _, err = itineraryRepository.AttachImages(
				ctx,
				itineraryEntity.V.Id,
				itineraryRequest.Images,
			); err != nil {
				return err
			}

			packageSessionEntity.ItineraryId = itineraryEntity.V.Id
		}

		// Create entity with repository
		packageSessionEntity, err := packageSessionRepository.Create(ctx, packageSessionEntity)
		if err != nil {
			return err
		}

		// Attach guides with repository
		if _, err := packageSessionRepository.AttachGuides(ctx, packageSessionEntity.Id, request.Guides); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.packageSessionMapper.MapEntityToResponse(
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

		return err
	})

	return response, err
}

func (s packageSessionServiceImpl) GetPackageSessionById(ctx context.Context, id int64) (dto.PackageSessionResponse, error) {
	// Validate id
	if err := s.packageSessionValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Find entity by id with repository
	packageSessionEntity, err := s.packageSessionRepository.FindById(ctx, id)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map entity into response
	response, err := s.packageSessionMapper.MapEntityToResponse(
		ctx,
		s.embarkationRepository,
		s.packageSessionRepository,
		s.imageRepository,
		s.flightRouteRepository,
		s.flightRepository,
		s.airlineRepository,
		s.airportRepository,
		s.busRepository,
		s.hotelRepository,
		s.itineraryRepository,
		s.itineraryDayRepository,
		s.itineraryWidgetRepository,
		s.itineraryWidgetActivityRepository,
		s.itineraryWidgetHotelRepository,
		s.itineraryWidgetInformationRepository,
		s.itineraryWidgetTransportRepository,
		s.itineraryWidgetRecommendationRepository,
		packageSessionEntity,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, nil
}

func (s packageSessionServiceImpl) GetAllPackageSession(ctx context.Context, request dto.GetAllPackageSessionRequest) ([]dto.PackageSessionListResponse, error) {
	// Validate request
	if err := s.packageSessionValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Create where find all options
	where := map[string]any{}
	if request.Package.Valid {
		where["package_id"] = request.Package.Int64
	}

	// Find all entities with repository
	packageSessionEntities, err := s.packageSessionRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
		Where:  where,
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.packageSessionMapper.MapEntitiesToListResponses(
		ctx,
		s.embarkationRepository,
		s.packageSessionRepository,
		s.flightRouteRepository,
		packageSessionEntities,
	)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s packageSessionServiceImpl) UpdatePackageSession(ctx context.Context, id int64, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error) {
	// Validate id
	if err := s.packageSessionValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Validate request
	if err := s.packageSessionValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map request into entity
	packageSessionEntity := s.packageSessionMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.PackageSessionResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create package session repository
		packageSessionRepository := factory.NewPackageSessionRepository()

		// Create embarkation repository
		embarkationRepository := factory.NewEmbarkationRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create flight route repository
		flightRouteRepository := factory.NewFlightRouteRepository()

		// Create flight repository
		flightRepository := factory.NewFlightRepository()

		// Create airline repository
		airlineRepository := factory.NewAirlineRepository()

		// Create airport repository
		airportRepository := factory.NewAirportRepository()

		// Create bus repository
		busRepository := factory.NewBusRepository()

		// Create hotel repository
		hotelRepository := factory.NewHotelRepository()

		// Create itinerary repository
		itineraryRepository := factory.NewItineraryRepository()

		// Create itinerary day repository
		itineraryDayRepository := factory.NewItineraryDayRepository()

		// Create itinerary widget repository
		itineraryWidgetRepository := factory.NewItineraryWidgetRepository()

		// Create itinerary widget activity repository
		itineraryWidgetActivityRepository := factory.NewItineraryWidgetActivityRepository()

		// Create itinerary widget hotel repository
		itineraryWidgetHotelRepository := factory.NewItineraryWidgetHotelRepository()

		// Create itinerary widget information repository
		itineraryWidgetInformationRepository := factory.NewItineraryWidgetInformationRepository()

		// Create itinerary widget transport repository
		itineraryWidgetTransportRepository := factory.NewItineraryWidgetTransportRepository()

		// Create itinerary widget recommendation repository
		itineraryWidgetRecommendationRepository := factory.NewItineraryWidgetRecommendationRepository()

		// Delete departure flight route
		departureFlightRoute, err := flightRouteRepository.FindById(ctx, packageSessionEntity.DepartureFlightRouteId)
		if err != nil {
			return err
		}
		for departureFlightRoute.NextId.Valid {
			if departureFlightRoute, err = flightRouteRepository.Delete(ctx, departureFlightRoute.NextId.Int64); err != nil {
				return err
			}
		}

		// Create departure flight route
		for _, departureFlight := range slices.Backward(request.DepartureFlights) {
			flightRoute, err := flightRouteRepository.Create(
				ctx,
				entity.FlightRoute{
					FlightId: departureFlight,
					NextId:   null.NewInt(packageSessionEntity.DepartureFlightRouteId, packageSessionEntity.DepartureFlightRouteId != 0),
				},
			)
			if err != nil {
				return err
			}

			packageSessionEntity.DepartureFlightRouteId = flightRoute.Id
		}

		// Delete return flight route
		returnFlightRoute, err := flightRouteRepository.FindById(ctx, packageSessionEntity.ReturnFlightRouteId)
		if err != nil {
			return err
		}
		for returnFlightRoute.NextId.Valid {
			if returnFlightRoute, err = flightRouteRepository.Delete(ctx, returnFlightRoute.NextId.Int64); err != nil {
				return err
			}
		}

		// Create return flight route
		for _, returnFlight := range slices.Backward(request.ReturnFlights) {
			flightRoute, err := flightRouteRepository.Create(
				ctx,
				entity.FlightRoute{
					FlightId: returnFlight,
					NextId:   null.NewInt(packageSessionEntity.ReturnFlightRouteId, packageSessionEntity.ReturnFlightRouteId != 0),
				},
			)
			if err != nil {
				return err
			}

			packageSessionEntity.ReturnFlightRouteId = flightRoute.Id
		}

		// Delete itinerary
		itinerary, err := itineraryRepository.FindById(ctx, packageSessionEntity.ItineraryId)
		if err != nil {
			return err
		}
		for itinerary.NextId.Valid {
			if itinerary, err = itineraryRepository.Delete(ctx, itinerary.NextId.Int64); err != nil {
				return err
			}
		}

		// Create itinerary with repository
		itineraryEntity := null.NewValue(entity.Itinerary{}, false)
		for _, itineraryRequest := range request.Itineraries {
			var err error

			itineraryEntity = null.ValueFrom(entity.Itinerary{
				City: itineraryRequest.City,
				NextId: null.NewInt(
					itineraryEntity.V.Id,
					itineraryEntity.Valid,
				),
			})

			itineraryDayEntity := null.NewValue(entity.ItineraryDay{}, false)
			for _, itineraryDayRequest := range itineraryRequest.Days {
				itineraryDayEntity = null.ValueFrom(entity.ItineraryDay{
					Title:       itineraryDayRequest.Title,
					Description: itineraryDayRequest.Description,
					NextId: null.NewInt(
						itineraryDayEntity.V.Id,
						itineraryDayEntity.Valid,
					),
				})

				itineraryWidgetEntity := null.NewValue(entity.ItineraryWidget{}, false)
				for _, itineraryWidgetRequest := range itineraryDayRequest.Widgets {
					itineraryWidgetEntity = null.ValueFrom(entity.ItineraryWidget{
						NextId: null.NewInt(
							itineraryWidgetEntity.V.Id,
							itineraryWidgetEntity.Valid,
						),
					})

					switch itineraryWidgetRequest.Type() {
					case "Activity":
						itineraryWidgetActivityRequest := itineraryWidgetRequest.(dto.ItineraryWidgetActivityRequest)
						itineraryWidgetActivityEntity, err := itineraryWidgetActivityRepository.Create(
							ctx,
							entity.ItineraryWidgetActivity{
								Title:       itineraryWidgetActivityRequest.Title,
								Description: itineraryWidgetActivityRequest.Description,
							},
						)
						if err != nil {
							return err
						}

						if _, err = itineraryWidgetActivityRepository.AttachImages(
							ctx,
							itineraryWidgetActivityEntity.Id,
							itineraryWidgetActivityRequest.Images,
						); err != nil {
							return err
						}

						itineraryWidgetEntity.V.ActivityId = null.IntFrom(itineraryWidgetActivityEntity.Id)
					case "Hotel":
						itineraryWidgetHotelRequest := itineraryWidgetRequest.(dto.ItineraryWidgetHotelRequest)
						itineraryWidgetHotelEntity, err := itineraryWidgetHotelRepository.Create(
							ctx,
							entity.ItineraryWidgetHotel{
								HotelId: itineraryWidgetHotelRequest.Hotel,
							},
						)
						if err != nil {
							return err
						}

						itineraryWidgetEntity.V.HotelId = null.IntFrom(itineraryWidgetHotelEntity.Id)
					case "Information":
						itineraryWidgetInformationRequest := itineraryWidgetRequest.(dto.ItineraryWidgetInformationRequest)
						itineraryWidgetInformationEntity, err := itineraryWidgetInformationRepository.Create(
							ctx,
							entity.ItineraryWidgetInformation{
								Description: itineraryWidgetInformationRequest.Description,
							},
						)
						if err != nil {
							return err
						}

						itineraryWidgetEntity.V.InformationId = null.IntFrom(itineraryWidgetInformationEntity.Id)
					case "Transport":
						itineraryWidgetTransportRequest := itineraryWidgetRequest.(dto.ItineraryWidgetTransportRequest)
						itineraryWidgetTransportEntity, err := itineraryWidgetTransportRepository.Create(
							ctx,
							entity.ItineraryWidgetTransport{
								Transportation: itineraryWidgetTransportRequest.Transportation,
								From:           itineraryWidgetTransportRequest.From,
								To:             itineraryWidgetTransportRequest.To,
							},
						)
						if err != nil {
							return err
						}

						itineraryWidgetEntity.V.TransportId = null.IntFrom(itineraryWidgetTransportEntity.Id)
					case "Recommendation":
						itineraryWidgetRecommendationRequest := itineraryWidgetRequest.(dto.ItineraryWidgetRecommendationRequest)
						itineraryWidgetRecommendationEntity, err := itineraryWidgetRecommendationRepository.Create(
							ctx,
							entity.ItineraryWidgetRecommendation{
								Description: itineraryWidgetRecommendationRequest.Description,
							},
						)
						if err != nil {
							return err
						}

						if _, err = itineraryWidgetRecommendationRepository.AttachImages(
							ctx,
							itineraryWidgetRecommendationEntity.Id,
							itineraryWidgetRecommendationRequest.Images,
						); err != nil {
							return err
						}

						itineraryWidgetEntity.V.RecommendationId = null.IntFrom(itineraryWidgetRecommendationEntity.Id)
					}

					itineraryWidgetEntity.V, err = itineraryWidgetRepository.Create(ctx, itineraryWidgetEntity.V)
					if err != nil {
						return err
					}
				}
				itineraryDayEntity.V.WidgetId = null.IntFrom(itineraryWidgetEntity.V.Id)

				itineraryDayEntity.V, err = itineraryDayRepository.Create(ctx, itineraryDayEntity.V)
				if err != nil {
					return err
				}
			}
			itineraryEntity.V.DayId = itineraryDayEntity.V.Id

			itineraryEntity.V, err = itineraryRepository.Create(ctx, itineraryEntity.V)
			if err != nil {
				return err
			}

			if _, err = itineraryRepository.AttachImages(
				ctx,
				itineraryEntity.V.Id,
				itineraryRequest.Images,
			); err != nil {
				return err
			}

			packageSessionEntity.ItineraryId = itineraryEntity.V.Id
		}

		// Update entity with repository
		packageSessionEntity, err := packageSessionRepository.Update(ctx, id, packageSessionEntity)
		if err != nil {
			return err
		}

		// Delete guides with repository
		if _, err := packageSessionRepository.DetachGuides(ctx, packageSessionEntity.Id); err != nil {
			return err
		}

		// Create guides with repository
		if _, err := packageSessionRepository.AttachGuides(ctx, packageSessionEntity.Id, request.Guides); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.packageSessionMapper.MapEntityToResponse(
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

		return err
	})

	return response, err
}

func (s packageSessionServiceImpl) DeletePackageSession(ctx context.Context, id int64) (dto.PackageSessionResponse, error) {
	// Validate id
	if err := s.packageSessionValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Delete entity with repository
	packageSessionEntity, err := s.packageSessionRepository.Delete(ctx, id)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map entity into response
	response, err := s.packageSessionMapper.MapEntityToResponse(
		ctx,
		s.embarkationRepository,
		s.packageSessionRepository,
		s.imageRepository,
		s.flightRouteRepository,
		s.flightRepository,
		s.airlineRepository,
		s.airportRepository,
		s.busRepository,
		s.hotelRepository,
		s.itineraryRepository,
		s.itineraryDayRepository,
		s.itineraryWidgetRepository,
		s.itineraryWidgetActivityRepository,
		s.itineraryWidgetHotelRepository,
		s.itineraryWidgetInformationRepository,
		s.itineraryWidgetTransportRepository,
		s.itineraryWidgetRecommendationRepository,
		packageSessionEntity,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, err
}
