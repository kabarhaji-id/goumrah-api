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

type facilityServiceImpl struct {
	facilityRepository repository.FacilityRepository
	facilityValidator  validator.FacilityValidator
	facilityMapper     mapper.FacilityMapper
}

func NewFacilityService(
	facilityRepository repository.FacilityRepository,
	facilityValidator validator.FacilityValidator,
	facilityMapper mapper.FacilityMapper,
) serviceport.FacilityService {
	return facilityServiceImpl{
		facilityRepository,
		facilityValidator,
		facilityMapper,
	}
}

func (s facilityServiceImpl) CreateFacility(ctx context.Context, request dto.FacilityRequest) (dto.FacilityResponse, error) {
	// Validate request
	if err := s.facilityValidator.ValidateRequest(ctx, request); err != nil {
		return dto.FacilityResponse{}, err
	}

	// Map request into entity
	facilityEntity := s.facilityMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	facilityEntity, err := s.facilityRepository.Create(ctx, facilityEntity)
	if err != nil {
		return dto.FacilityResponse{}, err
	}

	// Map entity into response
	response := s.facilityMapper.MapEntityToResponse(ctx, facilityEntity)

	return response, err
}

func (s facilityServiceImpl) GetFacilityById(ctx context.Context, id int64) (dto.FacilityResponse, error) {
	// Validate id
	if err := s.facilityValidator.ValidateId(ctx, id); err != nil {
		return dto.FacilityResponse{}, err
	}

	// Find entity by id with repository
	facilityEntity, err := s.facilityRepository.FindById(ctx, id)
	if err != nil {
		return dto.FacilityResponse{}, err
	}

	// Map entity into response
	response := s.facilityMapper.MapEntityToResponse(ctx, facilityEntity)

	return response, nil
}

func (s facilityServiceImpl) GetAllFacility(ctx context.Context, request dto.GetAllFacilityRequest) ([]dto.FacilityResponse, error) {
	// Validate request
	if err := s.facilityValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	facilityEntities, err := s.facilityRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.facilityMapper.MapEntitiesToResponses(ctx, facilityEntities)

	return responses, nil
}

func (s facilityServiceImpl) UpdateFacility(ctx context.Context, id int64, request dto.FacilityRequest) (dto.FacilityResponse, error) {
	// Validate id
	if err := s.facilityValidator.ValidateId(ctx, id); err != nil {
		return dto.FacilityResponse{}, err
	}

	// Validate request
	if err := s.facilityValidator.ValidateRequest(ctx, request); err != nil {
		return dto.FacilityResponse{}, err
	}

	// Map request into entity
	facilityEntity := s.facilityMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	facilityEntity, err := s.facilityRepository.Update(ctx, id, facilityEntity)
	if err != nil {
		return dto.FacilityResponse{}, err
	}

	// Map entity into response
	response := s.facilityMapper.MapEntityToResponse(ctx, facilityEntity)

	return response, err
}

func (s facilityServiceImpl) DeleteFacility(ctx context.Context, id int64) (dto.FacilityResponse, error) {
	// Validate id
	if err := s.facilityValidator.ValidateId(ctx, id); err != nil {
		return dto.FacilityResponse{}, err
	}

	// Delete entity with repository
	facilityEntity, err := s.facilityRepository.Delete(ctx, id)
	if err != nil {
		return dto.FacilityResponse{}, err
	}

	// Map entity into response
	response := s.facilityMapper.MapEntityToResponse(ctx, facilityEntity)

	return response, err
}
