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

type embarkationServiceImpl struct {
	embarkationRepository repository.EmbarkationRepository
	embarkationValidator  validator.EmbarkationValidator
	embarkationMapper     mapper.EmbarkationMapper
}

func NewEmbarkationService(
	embarkationRepository repository.EmbarkationRepository,
	embarkationValidator validator.EmbarkationValidator,
	embarkationMapper mapper.EmbarkationMapper,
) serviceport.EmbarkationService {
	return embarkationServiceImpl{
		embarkationRepository,
		embarkationValidator,
		embarkationMapper,
	}
}

func (s embarkationServiceImpl) CreateEmbarkation(ctx context.Context, request dto.EmbarkationRequest) (dto.EmbarkationResponse, error) {
	// Validate request
	if err := s.embarkationValidator.ValidateRequest(ctx, request); err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Map request into entity
	embarkationEntity := s.embarkationMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	embarkationEntity, err := s.embarkationRepository.Create(ctx, embarkationEntity)
	if err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Map entity into response
	response := s.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	return response, err
}

func (s embarkationServiceImpl) GetEmbarkationById(ctx context.Context, id int64) (dto.EmbarkationResponse, error) {
	// Validate id
	if err := s.embarkationValidator.ValidateId(ctx, id); err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Find entity by id with repository
	embarkationEntity, err := s.embarkationRepository.FindById(ctx, id)
	if err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Map entity into response
	response := s.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	return response, nil
}

func (s embarkationServiceImpl) GetAllEmbarkation(ctx context.Context, request dto.GetAllEmbarkationRequest) ([]dto.EmbarkationResponse, error) {
	// Validate request
	if err := s.embarkationValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	embarkationEntities, err := s.embarkationRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.embarkationMapper.MapEntitiesToResponses(ctx, embarkationEntities)

	return responses, nil
}

func (s embarkationServiceImpl) UpdateEmbarkation(ctx context.Context, id int64, request dto.EmbarkationRequest) (dto.EmbarkationResponse, error) {
	// Validate id
	if err := s.embarkationValidator.ValidateId(ctx, id); err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Validate request
	if err := s.embarkationValidator.ValidateRequest(ctx, request); err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Map request into entity
	embarkationEntity := s.embarkationMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	embarkationEntity, err := s.embarkationRepository.Update(ctx, id, embarkationEntity)
	if err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Map entity into response
	response := s.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	return response, err
}

func (s embarkationServiceImpl) DeleteEmbarkation(ctx context.Context, id int64) (dto.EmbarkationResponse, error) {
	// Validate id
	if err := s.embarkationValidator.ValidateId(ctx, id); err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Delete entity with repository
	embarkationEntity, err := s.embarkationRepository.Delete(ctx, id)
	if err != nil {
		return dto.EmbarkationResponse{}, err
	}

	// Map entity into response
	response := s.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	return response, err
}
