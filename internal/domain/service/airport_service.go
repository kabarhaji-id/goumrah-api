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

type airportServiceImpl struct {
	airportRepository repository.AirportRepository
	airportValidator  validator.AirportValidator
	airportMapper     mapper.AirportMapper
}

func NewAirportService(
	airportRepository repository.AirportRepository,
	airportValidator validator.AirportValidator,
	airportMapper mapper.AirportMapper,
) serviceport.AirportService {
	return airportServiceImpl{
		airportRepository,
		airportValidator,
		airportMapper,
	}
}

func (s airportServiceImpl) CreateAirport(ctx context.Context, request dto.AirportRequest) (dto.AirportResponse, error) {
	// Validate request
	if err := s.airportValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AirportResponse{}, err
	}

	// Map request into entity
	airportEntity := s.airportMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	airportEntity, err := s.airportRepository.Create(ctx, airportEntity)
	if err != nil {
		return dto.AirportResponse{}, err
	}

	// Map entity into response
	response := s.airportMapper.MapEntityToResponse(ctx, airportEntity)

	return response, err
}

func (s airportServiceImpl) GetAirportById(ctx context.Context, id int64) (dto.AirportResponse, error) {
	// Validate id
	if err := s.airportValidator.ValidateId(ctx, id); err != nil {
		return dto.AirportResponse{}, err
	}

	// Find entity by id with repository
	airportEntity, err := s.airportRepository.FindById(ctx, id)
	if err != nil {
		return dto.AirportResponse{}, err
	}

	// Map entity into response
	response := s.airportMapper.MapEntityToResponse(ctx, airportEntity)

	return response, nil
}

func (s airportServiceImpl) GetAllAirport(ctx context.Context, request dto.GetAllAirportRequest) ([]dto.AirportResponse, error) {
	// Validate request
	if err := s.airportValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	airportEntities, err := s.airportRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.airportMapper.MapEntitiesToResponses(ctx, airportEntities)

	return responses, nil
}

func (s airportServiceImpl) UpdateAirport(ctx context.Context, id int64, request dto.AirportRequest) (dto.AirportResponse, error) {
	// Validate id
	if err := s.airportValidator.ValidateId(ctx, id); err != nil {
		return dto.AirportResponse{}, err
	}

	// Validate request
	if err := s.airportValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AirportResponse{}, err
	}

	// Map request into entity
	airportEntity := s.airportMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	airportEntity, err := s.airportRepository.Update(ctx, id, airportEntity)
	if err != nil {
		return dto.AirportResponse{}, err
	}

	// Map entity into response
	response := s.airportMapper.MapEntityToResponse(ctx, airportEntity)

	return response, err
}

func (s airportServiceImpl) DeleteAirport(ctx context.Context, id int64) (dto.AirportResponse, error) {
	// Validate id
	if err := s.airportValidator.ValidateId(ctx, id); err != nil {
		return dto.AirportResponse{}, err
	}

	// Delete entity with repository
	airportEntity, err := s.airportRepository.Delete(ctx, id)
	if err != nil {
		return dto.AirportResponse{}, err
	}

	// Map entity into response
	response := s.airportMapper.MapEntityToResponse(ctx, airportEntity)

	return response, err
}
