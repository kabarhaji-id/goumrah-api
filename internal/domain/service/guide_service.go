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

type guideServiceImpl struct {
	guideRepository repository.GuideRepository
	guideValidator  validator.GuideValidator
	guideMapper     mapper.GuideMapper

	imageRepository repository.ImageRepository
}

func NewGuideService(
	guideRepository repository.GuideRepository,
	guideValidator validator.GuideValidator,
	guideMapper mapper.GuideMapper,
	imageRepository repository.ImageRepository,
) serviceport.GuideService {
	return guideServiceImpl{
		guideRepository,
		guideValidator,
		guideMapper,
		imageRepository,
	}
}

func (s guideServiceImpl) CreateGuide(ctx context.Context, request dto.GuideRequest) (dto.GuideResponse, error) {
	// Validate request
	if err := s.guideValidator.ValidateRequest(ctx, request); err != nil {
		return dto.GuideResponse{}, err
	}

	// Map request into entity
	guideEntity := s.guideMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	guideEntity, err := s.guideRepository.Create(ctx, guideEntity)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	// Map entity into response
	response, err := s.guideMapper.MapEntityToResponse(ctx, s.imageRepository, guideEntity)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	return response, err
}

func (s guideServiceImpl) GetGuideById(ctx context.Context, id int64) (dto.GuideResponse, error) {
	// Validate id
	if err := s.guideValidator.ValidateId(ctx, id); err != nil {
		return dto.GuideResponse{}, err
	}

	// Find entity by id with repository
	guideEntity, err := s.guideRepository.FindById(ctx, id)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	// Map entity into response
	response, err := s.guideMapper.MapEntityToResponse(ctx, s.imageRepository, guideEntity)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	return response, nil
}

func (s guideServiceImpl) GetAllGuide(ctx context.Context, request dto.GetAllGuideRequest) ([]dto.GuideResponse, error) {
	// Validate request
	if err := s.guideValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	guideEntities, err := s.guideRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.guideMapper.MapEntitiesToResponses(ctx, s.imageRepository, guideEntities)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s guideServiceImpl) UpdateGuide(ctx context.Context, id int64, request dto.GuideRequest) (dto.GuideResponse, error) {
	// Validate id
	if err := s.guideValidator.ValidateId(ctx, id); err != nil {
		return dto.GuideResponse{}, err
	}

	// Validate request
	if err := s.guideValidator.ValidateRequest(ctx, request); err != nil {
		return dto.GuideResponse{}, err
	}

	// Map request into entity
	guideEntity := s.guideMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	guideEntity, err := s.guideRepository.Update(ctx, id, guideEntity)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	// Map entity into response
	response, err := s.guideMapper.MapEntityToResponse(ctx, s.imageRepository, guideEntity)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	return response, err
}

func (s guideServiceImpl) DeleteGuide(ctx context.Context, id int64) (dto.GuideResponse, error) {
	// Validate id
	if err := s.guideValidator.ValidateId(ctx, id); err != nil {
		return dto.GuideResponse{}, err
	}

	// Delete entity with repository
	guideEntity, err := s.guideRepository.Delete(ctx, id)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	// Map entity into response
	response, err := s.guideMapper.MapEntityToResponse(ctx, s.imageRepository, guideEntity)
	if err != nil {
		return dto.GuideResponse{}, err
	}

	return response, err
}
