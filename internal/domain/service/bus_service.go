package service

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	serviceport "github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type busServiceImpl struct {
	busRepository repository.BusRepository
	busValidator  validator.BusValidator
	busMapper     mapper.BusMapper
}

func NewBusService(
	busRepository repository.BusRepository,
	busValidator validator.BusValidator,
	busMapper mapper.BusMapper,
) serviceport.BusService {
	return busServiceImpl{
		busRepository,
		busValidator,
		busMapper,
	}
}

func (s busServiceImpl) CreateBus(ctx context.Context, request dto.BusRequest) (dto.BusResponse, error) {
	// Validate request
	if err := s.busValidator.ValidateRequest(ctx, request); err != nil {
		return dto.BusResponse{}, err
	}

	// Map request into entity
	busEntity := s.busMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	busEntity, err := s.busRepository.Create(ctx, busEntity)
	if err != nil {
		return dto.BusResponse{}, err
	}

	// Map entity into response
	response := s.busMapper.MapEntityToResponse(ctx, busEntity)

	return response, err
}

func (s busServiceImpl) GetBusById(ctx context.Context, id int64) (dto.BusResponse, error) {
	// Validate id
	if err := s.busValidator.ValidateId(ctx, id); err != nil {
		return dto.BusResponse{}, err
	}

	// Find entity by id with repository
	busEntity, err := s.busRepository.FindById(ctx, id)
	if err != nil {
		return dto.BusResponse{}, err
	}

	// Map entity into response
	response := s.busMapper.MapEntityToResponse(ctx, busEntity)

	return response, nil
}

func (s busServiceImpl) GetAllBus(ctx context.Context, request dto.GetAllBusRequest) ([]dto.BusResponse, error) {
	// Validate request
	if err := s.busValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	busEntities, err := s.busRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.busMapper.MapEntitiesToResponses(ctx, busEntities)

	return responses, nil
}

func (s busServiceImpl) UpdateBus(ctx context.Context, id int64, request dto.BusRequest) (dto.BusResponse, error) {
	// Validate id
	if err := s.busValidator.ValidateId(ctx, id); err != nil {
		return dto.BusResponse{}, err
	}

	// Validate request
	if err := s.busValidator.ValidateRequest(ctx, request); err != nil {
		return dto.BusResponse{}, err
	}

	// Map request into entity
	busEntity := s.busMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	busEntity, err := s.busRepository.Update(ctx, id, busEntity)
	if err != nil {
		return dto.BusResponse{}, err
	}

	// Map entity into response
	response := s.busMapper.MapEntityToResponse(ctx, busEntity)

	return response, err
}

func (s busServiceImpl) DeleteBus(ctx context.Context, id int64) (dto.BusResponse, error) {
	// Validate id
	if err := s.busValidator.ValidateId(ctx, id); err != nil {
		return dto.BusResponse{}, err
	}

	// Delete entity with repository
	busEntity, err := s.busRepository.Delete(ctx, id)
	if err != nil {
		return dto.BusResponse{}, err
	}

	// Map entity into response
	response := s.busMapper.MapEntityToResponse(ctx, busEntity)

	return response, err
}
