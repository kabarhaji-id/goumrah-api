package service

import (
	"context"
	"slices"

	"github.com/guregu/null/v5"
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

		// Delete departure flight route
		departureFlightRoute, err := s.flightRouteRepository.FindById(ctx, packageSessionEntity.DepartureFlightRouteId)
		if err != nil {
			return err
		}
		for departureFlightRoute.NextId.Valid {
			if departureFlightRoute, err = s.flightRouteRepository.Delete(ctx, departureFlightRoute.NextId.Int64); err != nil {
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
		returnFlightRoute, err := s.flightRouteRepository.FindById(ctx, packageSessionEntity.ReturnFlightRouteId)
		if err != nil {
			return err
		}
		for returnFlightRoute.NextId.Valid {
			if returnFlightRoute, err = s.flightRouteRepository.Delete(ctx, returnFlightRoute.NextId.Int64); err != nil {
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
		packageSessionEntity,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, err
}
