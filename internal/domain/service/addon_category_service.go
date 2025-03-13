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

type addonCategoryServiceImpl struct {
	addonCategoryRepository repository.AddonCategoryRepository
	addonCategoryValidator  validator.AddonCategoryValidator
	addonCategoryMapper     mapper.AddonCategoryMapper
}

func NewAddonCategoryService(
	addonCategoryRepository repository.AddonCategoryRepository,
	addonCategoryValidator validator.AddonCategoryValidator,
	addonCategoryMapper mapper.AddonCategoryMapper,
) serviceport.AddonCategoryService {
	return addonCategoryServiceImpl{
		addonCategoryRepository,
		addonCategoryValidator,
		addonCategoryMapper,
	}
}

func (s addonCategoryServiceImpl) CreateAddonCategory(ctx context.Context, request dto.AddonCategoryRequest) (dto.AddonCategoryResponse, error) {
	// Validate request
	if err := s.addonCategoryValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Map request into entity
	addonCategoryEntity := s.addonCategoryMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	addonCategoryEntity, err := s.addonCategoryRepository.Create(ctx, addonCategoryEntity)
	if err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Map entity into response
	response := s.addonCategoryMapper.MapEntityToResponse(ctx, addonCategoryEntity)

	return response, err
}

func (s addonCategoryServiceImpl) GetAddonCategoryById(ctx context.Context, id int64) (dto.AddonCategoryResponse, error) {
	// Validate id
	if err := s.addonCategoryValidator.ValidateId(ctx, id); err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Find entity by id with repository
	addonCategoryEntity, err := s.addonCategoryRepository.FindById(ctx, id)
	if err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Map entity into response
	response := s.addonCategoryMapper.MapEntityToResponse(ctx, addonCategoryEntity)

	return response, nil
}

func (s addonCategoryServiceImpl) GetAllAddonCategory(ctx context.Context, request dto.GetAllAddonCategoryRequest) ([]dto.AddonCategoryResponse, error) {
	// Validate request
	if err := s.addonCategoryValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	addonCategoryEntities, err := s.addonCategoryRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.addonCategoryMapper.MapEntitiesToResponses(ctx, addonCategoryEntities)

	return responses, nil
}

func (s addonCategoryServiceImpl) UpdateAddonCategory(ctx context.Context, id int64, request dto.AddonCategoryRequest) (dto.AddonCategoryResponse, error) {
	// Validate id
	if err := s.addonCategoryValidator.ValidateId(ctx, id); err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Validate request
	if err := s.addonCategoryValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Map request into entity
	addonCategoryEntity := s.addonCategoryMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	addonCategoryEntity, err := s.addonCategoryRepository.Update(ctx, id, addonCategoryEntity)
	if err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Map entity into response
	response := s.addonCategoryMapper.MapEntityToResponse(ctx, addonCategoryEntity)

	return response, err
}

func (s addonCategoryServiceImpl) DeleteAddonCategory(ctx context.Context, id int64) (dto.AddonCategoryResponse, error) {
	// Validate id
	if err := s.addonCategoryValidator.ValidateId(ctx, id); err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Delete entity with repository
	addonCategoryEntity, err := s.addonCategoryRepository.Delete(ctx, id)
	if err != nil {
		return dto.AddonCategoryResponse{}, err
	}

	// Map entity into response
	response := s.addonCategoryMapper.MapEntityToResponse(ctx, addonCategoryEntity)

	return response, err
}
