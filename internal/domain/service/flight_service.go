package service

import (
	"context"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	serviceport "github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type flightServiceImpl struct {
	flightRepository repository.FlightRepository
	flightValidator  validator.FlightValidator
	flightMapper     mapper.FlightMapper

	imageRepository   repository.ImageRepository
	airlineRepository repository.AirlineRepository
	airportRepository repository.AirportRepository
}

func NewFlightService(
	flightRepository repository.FlightRepository,
	flightValidator validator.FlightValidator,
	flightMapper mapper.FlightMapper,
	imageRepository repository.ImageRepository,
	airlineRepository repository.AirlineRepository,
	airportRepository repository.AirportRepository,
) serviceport.FlightService {
	return flightServiceImpl{
		flightRepository,
		flightValidator,
		flightMapper,
		imageRepository,
		airlineRepository,
		airportRepository,
	}
}

func (s flightServiceImpl) CreateFlight(ctx context.Context, request dto.FlightRequest) (dto.FlightResponse, error) {
	// Validate request
	if err := s.flightValidator.ValidateRequest(ctx, request); err != nil {
		return dto.FlightResponse{}, err
	}

	// Map request into entity
	flightEntity := s.flightMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	flightEntity, err := s.flightRepository.Create(ctx, flightEntity)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	// Map entity into response
	response, err := s.flightMapper.MapEntityToResponse(
		ctx,
		s.imageRepository, s.airlineRepository, s.airportRepository,
		flightEntity,
	)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	return response, err
}

func (s flightServiceImpl) GetFlightById(ctx context.Context, id int64) (dto.FlightResponse, error) {
	// Validate id
	if err := s.flightValidator.ValidateId(ctx, id); err != nil {
		return dto.FlightResponse{}, err
	}

	// Find entity by id with repository
	flightEntity, err := s.flightRepository.FindById(ctx, id)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	// Map entity into response
	response, err := s.flightMapper.MapEntityToResponse(
		ctx,
		s.imageRepository, s.airlineRepository, s.airportRepository,
		flightEntity,
	)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	return response, nil
}

func (s flightServiceImpl) GetAllFlight(ctx context.Context, request dto.GetAllFlightRequest) ([]dto.FlightResponse, error) {
	// Validate request
	if err := s.flightValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	flightEntities, err := s.flightRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.flightMapper.MapEntitiesToResponses(
		ctx,
		s.imageRepository, s.airlineRepository, s.airportRepository,
		flightEntities,
	)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s flightServiceImpl) UpdateFlight(ctx context.Context, id int64, request dto.FlightRequest) (dto.FlightResponse, error) {
	// Validate id
	if err := s.flightValidator.ValidateId(ctx, id); err != nil {
		return dto.FlightResponse{}, err
	}

	// Validate request
	if err := s.flightValidator.ValidateRequest(ctx, request); err != nil {
		return dto.FlightResponse{}, err
	}

	// Map request into entity
	flightEntity := s.flightMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	flightEntity, err := s.flightRepository.Update(ctx, id, flightEntity)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	// Map entity into response
	response, err := s.flightMapper.MapEntityToResponse(
		ctx,
		s.imageRepository, s.airlineRepository, s.airportRepository,
		flightEntity,
	)

	return response, err
}

func (s flightServiceImpl) DeleteFlight(ctx context.Context, id int64) (dto.FlightResponse, error) {
	// Validate id
	if err := s.flightValidator.ValidateId(ctx, id); err != nil {
		return dto.FlightResponse{}, err
	}

	// Delete entity with repository
	flightEntity, err := s.flightRepository.Delete(ctx, id)
	if err != nil {
		return dto.FlightResponse{}, err
	}

	// Map entity into response
	response, err := s.flightMapper.MapEntityToResponse(
		ctx,
		s.imageRepository, s.airlineRepository, s.airportRepository,
		flightEntity,
	)

	return response, err
}
