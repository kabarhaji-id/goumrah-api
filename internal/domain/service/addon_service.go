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

type addonServiceImpl struct {
	addonRepository repository.AddonRepository
	addonValidator  validator.AddonValidator
	addonMapper     mapper.AddonMapper

	addonCategoryRepository repository.AddonCategoryRepository
}

func NewAddonService(
	addonRepository repository.AddonRepository,
	addonValidator validator.AddonValidator,
	addonMapper mapper.AddonMapper,
	addonCategoryRepository repository.AddonCategoryRepository,
) serviceport.AddonService {
	return addonServiceImpl{
		addonRepository,
		addonValidator,
		addonMapper,
		addonCategoryRepository,
	}
}

func (s addonServiceImpl) CreateAddon(ctx context.Context, request dto.AddonRequest) (dto.AddonResponse, error) {
	// Validate request
	if err := s.addonValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AddonResponse{}, err
	}

	// Map request into entity
	addonEntity := s.addonMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	addonEntity, err := s.addonRepository.Create(ctx, addonEntity)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	// Map entity into response
	response, err := s.addonMapper.MapEntityToResponse(ctx, s.addonCategoryRepository, addonEntity)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	return response, err
}

func (s addonServiceImpl) GetAddonById(ctx context.Context, id int64) (dto.AddonResponse, error) {
	// Validate id
	if err := s.addonValidator.ValidateId(ctx, id); err != nil {
		return dto.AddonResponse{}, err
	}

	// Find entity by id with repository
	addonEntity, err := s.addonRepository.FindById(ctx, id)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	// Map entity into response
	response, err := s.addonMapper.MapEntityToResponse(ctx, s.addonCategoryRepository, addonEntity)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	return response, nil
}

func (s addonServiceImpl) GetAllAddon(ctx context.Context, request dto.GetAllAddonRequest) ([]dto.AddonResponse, error) {
	// Validate request
	if err := s.addonValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	addonEntities, err := s.addonRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.addonMapper.MapEntitiesToResponses(ctx, s.addonCategoryRepository, addonEntities)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s addonServiceImpl) UpdateAddon(ctx context.Context, id int64, request dto.AddonRequest) (dto.AddonResponse, error) {
	// Validate id
	if err := s.addonValidator.ValidateId(ctx, id); err != nil {
		return dto.AddonResponse{}, err
	}

	// Validate request
	if err := s.addonValidator.ValidateRequest(ctx, request); err != nil {
		return dto.AddonResponse{}, err
	}

	// Map request into entity
	addonEntity := s.addonMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	addonEntity, err := s.addonRepository.Update(ctx, id, addonEntity)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	// Map entity into response
	response, err := s.addonMapper.MapEntityToResponse(ctx, s.addonCategoryRepository, addonEntity)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	return response, err
}

func (s addonServiceImpl) DeleteAddon(ctx context.Context, id int64) (dto.AddonResponse, error) {
	// Validate id
	if err := s.addonValidator.ValidateId(ctx, id); err != nil {
		return dto.AddonResponse{}, err
	}

	// Delete entity with repository
	addonEntity, err := s.addonRepository.Delete(ctx, id)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	// Map entity into response
	response, err := s.addonMapper.MapEntityToResponse(ctx, s.addonCategoryRepository, addonEntity)
	if err != nil {
		return dto.AddonResponse{}, err
	}

	return response, err
}
