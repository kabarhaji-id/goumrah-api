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

type airlineServiceImpl struct {
	airlineRepository repository.AirlineRepository
	airlineValidator  validator.AirlineValidator
	airlineMapper     mapper.AirlineMapper

	imageRepository repository.ImageRepository
}

func NewAirlineService(
	airlineRepository repository.AirlineRepository,
	airlineValidator validator.AirlineValidator,
	airlineMapper mapper.AirlineMapper,
	imageRepository repository.ImageRepository,
) serviceport.AirlineService {
	return airlineServiceImpl{
		airlineRepository,
		airlineValidator,
		airlineMapper,
		imageRepository,
	}
}

func (s airlineServiceImpl) CreateAirline(ctx context.Context, request dto.AirlineRequest) (dto.AirlineResponse, error) {
	// Validate request
	if err := s.airlineValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AirlineResponse{}, err
	}

	// Map request into entity
	airlineEntity := s.airlineMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	airlineEntity, err := s.airlineRepository.Create(ctx, airlineEntity)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	// Map entity into response
	response, err := s.airlineMapper.MapEntityToResponse(ctx, s.imageRepository, airlineEntity)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	return response, err
}

func (s airlineServiceImpl) GetAirlineById(ctx context.Context, id int64) (dto.AirlineResponse, error) {
	// Validate id
	if err := s.airlineValidator.ValidateId(ctx, id); err != nil {
		return dto.AirlineResponse{}, err
	}

	// Find entity by id with repository
	airlineEntity, err := s.airlineRepository.FindById(ctx, id)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	// Map entity into response
	response, err := s.airlineMapper.MapEntityToResponse(ctx, s.imageRepository, airlineEntity)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	return response, nil
}

func (s airlineServiceImpl) GetAllAirline(ctx context.Context, request dto.GetAllAirlineRequest) ([]dto.AirlineResponse, error) {
	// Validate request
	if err := s.airlineValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	airlineEntities, err := s.airlineRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.airlineMapper.MapEntitiesToResponses(ctx, s.imageRepository, airlineEntities)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s airlineServiceImpl) UpdateAirline(ctx context.Context, id int64, request dto.AirlineRequest) (dto.AirlineResponse, error) {
	// Validate id
	if err := s.airlineValidator.ValidateId(ctx, id); err != nil {
		return dto.AirlineResponse{}, err
	}

	// Validate request
	if err := s.airlineValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AirlineResponse{}, err
	}

	// Map request into entity
	airlineEntity := s.airlineMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	airlineEntity, err := s.airlineRepository.Update(ctx, id, airlineEntity)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	// Map entity into response
	response, err := s.airlineMapper.MapEntityToResponse(ctx, s.imageRepository, airlineEntity)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	return response, err
}

func (s airlineServiceImpl) DeleteAirline(ctx context.Context, id int64) (dto.AirlineResponse, error) {
	// Validate id
	if err := s.airlineValidator.ValidateId(ctx, id); err != nil {
		return dto.AirlineResponse{}, err
	}

	// Delete entity with repository
	airlineEntity, err := s.airlineRepository.Delete(ctx, id)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	// Map entity into response
	response, err := s.airlineMapper.MapEntityToResponse(ctx, s.imageRepository, airlineEntity)
	if err != nil {
		return dto.AirlineResponse{}, err
	}

	return response, err
}
