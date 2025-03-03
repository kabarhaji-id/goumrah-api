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

type packageSessionServiceImpl struct {
	packageSessionRepository repository.PackageSessionRepository
	packageSessionValidator  validator.PackageSessionValidator
	packageSessionMapper     mapper.PackageSessionMapper
}

func NewPackageSessionService(
	packageSessionRepository repository.PackageSessionRepository,
	packageSessionValidator validator.PackageSessionValidator,
	packageSessionMapper mapper.PackageSessionMapper,
) serviceport.PackageSessionService {
	return packageSessionServiceImpl{
		packageSessionRepository,
		packageSessionValidator,
		packageSessionMapper,
	}
}

func (s packageSessionServiceImpl) CreatePackageSession(ctx context.Context, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error) {
	// Validate request
	if err := s.packageSessionValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map request into entity
	packageSessionEntity := s.packageSessionMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	packageSessionEntity, err := s.packageSessionRepository.Create(ctx, packageSessionEntity)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map entity into response
	response, err := s.packageSessionMapper.MapEntityToResponse(ctx, packageSessionEntity)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, err
}

func (s packageSessionServiceImpl) GetPackageSessionById(ctx context.Context, id int64) (dto.PackageSessionResponse, error) {
	// Validate id
	if err := s.packageSessionValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Find entity by id with repository
	packageSessionEntity, err := s.packageSessionRepository.FindById(ctx, id)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map entity into response
	response, err := s.packageSessionMapper.MapEntityToResponse(ctx, packageSessionEntity)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, nil
}

func (s packageSessionServiceImpl) GetAllPackageSession(ctx context.Context, request dto.GetAllPackageSessionRequest) ([]dto.PackageSessionResponse, error) {
	// Validate request
	if err := s.packageSessionValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	packageSessionEntities, err := s.packageSessionRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.packageSessionMapper.MapEntitiesToResponses(ctx, packageSessionEntities)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s packageSessionServiceImpl) UpdatePackageSession(ctx context.Context, id int64, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error) {
	// Validate id
	if err := s.packageSessionValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Validate request
	if err := s.packageSessionValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map request into entity
	packageSessionEntity := s.packageSessionMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	packageSessionEntity, err := s.packageSessionRepository.Update(ctx, id, packageSessionEntity)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map entity into response
	response, err := s.packageSessionMapper.MapEntityToResponse(ctx, packageSessionEntity)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, err
}

func (s packageSessionServiceImpl) DeletePackageSession(ctx context.Context, id int64) (dto.PackageSessionResponse, error) {
	// Validate id
	if err := s.packageSessionValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Delete entity with repository
	packageSessionEntity, err := s.packageSessionRepository.Delete(ctx, id)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map entity into response
	response, err := s.packageSessionMapper.MapEntityToResponse(ctx, packageSessionEntity)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, err
}
