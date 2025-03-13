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

type cityTourServiceImpl struct {
	cityTourRepository repository.CityTourRepository
	cityTourValidator  validator.CityTourValidator
	cityTourMapper     mapper.CityTourMapper
}

func NewCityTourService(
	cityTourRepository repository.CityTourRepository,
	cityTourValidator validator.CityTourValidator,
	cityTourMapper mapper.CityTourMapper,
) serviceport.CityTourService {
	return cityTourServiceImpl{
		cityTourRepository,
		cityTourValidator,
		cityTourMapper,
	}
}

func (s cityTourServiceImpl) CreateCityTour(ctx context.Context, request dto.CityTourRequest) (dto.CityTourResponse, error) {
	// Validate request
	if err := s.cityTourValidator.ValidateRequest(ctx, request); err != nil {
		return dto.CityTourResponse{}, err
	}

	// Map request into entity
	cityTourEntity := s.cityTourMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	cityTourEntity, err := s.cityTourRepository.Create(ctx, cityTourEntity)
	if err != nil {
		return dto.CityTourResponse{}, err
	}

	// Map entity into response
	response := s.cityTourMapper.MapEntityToResponse(ctx, cityTourEntity)

	return response, err
}

func (s cityTourServiceImpl) GetCityTourById(ctx context.Context, id int64) (dto.CityTourResponse, error) {
	// Validate id
	if err := s.cityTourValidator.ValidateId(ctx, id); err != nil {
		return dto.CityTourResponse{}, err
	}

	// Find entity by id with repository
	cityTourEntity, err := s.cityTourRepository.FindById(ctx, id)
	if err != nil {
		return dto.CityTourResponse{}, err
	}

	// Map entity into response
	response := s.cityTourMapper.MapEntityToResponse(ctx, cityTourEntity)

	return response, nil
}

func (s cityTourServiceImpl) GetAllCityTour(ctx context.Context, request dto.GetAllCityTourRequest) ([]dto.CityTourResponse, error) {
	// Validate request
	if err := s.cityTourValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	cityTourEntities, err := s.cityTourRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.cityTourMapper.MapEntitiesToResponses(ctx, cityTourEntities)

	return responses, nil
}

func (s cityTourServiceImpl) UpdateCityTour(ctx context.Context, id int64, request dto.CityTourRequest) (dto.CityTourResponse, error) {
	// Validate id
	if err := s.cityTourValidator.ValidateId(ctx, id); err != nil {
		return dto.CityTourResponse{}, err
	}

	// Validate request
	if err := s.cityTourValidator.ValidateRequest(ctx, request); err != nil {
		return dto.CityTourResponse{}, err
	}

	// Map request into entity
	cityTourEntity := s.cityTourMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	cityTourEntity, err := s.cityTourRepository.Update(ctx, id, cityTourEntity)
	if err != nil {
		return dto.CityTourResponse{}, err
	}

	// Map entity into response
	response := s.cityTourMapper.MapEntityToResponse(ctx, cityTourEntity)

	return response, err
}

func (s cityTourServiceImpl) DeleteCityTour(ctx context.Context, id int64) (dto.CityTourResponse, error) {
	// Validate id
	if err := s.cityTourValidator.ValidateId(ctx, id); err != nil {
		return dto.CityTourResponse{}, err
	}

	// Delete entity with repository
	cityTourEntity, err := s.cityTourRepository.Delete(ctx, id)
	if err != nil {
		return dto.CityTourResponse{}, err
	}

	// Map entity into response
	response := s.cityTourMapper.MapEntityToResponse(ctx, cityTourEntity)

	return response, err
}
