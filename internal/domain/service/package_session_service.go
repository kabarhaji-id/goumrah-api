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

	embarkationRepository repository.EmbarkationRepository

	imageRepository repository.ImageRepository

	unitOfWork repository.UnitOfWork
}

func NewPackageSessionService(
	packageSessionRepository repository.PackageSessionRepository,
	packageSessionValidator validator.PackageSessionValidator,
	packageSessionMapper mapper.PackageSessionMapper,
	embarkationRepository repository.EmbarkationRepository,
	imageRepository repository.ImageRepository,
	unitOfWork repository.UnitOfWork,
) serviceport.PackageSessionService {
	return packageSessionServiceImpl{
		packageSessionRepository,
		packageSessionValidator,
		packageSessionMapper,
		embarkationRepository,
		imageRepository,
		unitOfWork,
	}
}

func (s packageSessionServiceImpl) CreatePackageSession(ctx context.Context, request dto.PackageSessionRequest) (dto.PackageSessionResponse, error) {
	// Validate request
	if err := s.packageSessionValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageSessionResponse{}, err
	}

	// Map request into entity
	packageSessionEntity := s.packageSessionMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.PackageSessionResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create package session repository
		packageSessionRepository := factory.NewPackageSessionRepository()

		// Create embarkation repository
		embarkationRepository := factory.NewEmbarkationRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create entity with repository
		packageSessionEntity, err := packageSessionRepository.Create(ctx, packageSessionEntity)
		if err != nil {
			return err
		}

		// Create guides with repository
		if _, err := packageSessionRepository.CreateGuides(ctx, packageSessionEntity.Id, request.Guides); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.packageSessionMapper.MapEntityToResponse(
			ctx,
			embarkationRepository,
			packageSessionRepository,
			imageRepository,
			packageSessionEntity,
		)

		return err
	})

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
	response, err := s.packageSessionMapper.MapEntityToResponse(
		ctx,
		s.embarkationRepository,
		s.packageSessionRepository,
		s.imageRepository,
		packageSessionEntity,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, nil
}

func (s packageSessionServiceImpl) GetAllPackageSession(ctx context.Context, request dto.GetAllPackageSessionRequest) ([]dto.PackageSessionListResponse, error) {
	// Validate request
	if err := s.packageSessionValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Create where find all options
	where := map[string]any{}
	if request.Package.Valid {
		where["package_id"] = request.Package.Int64
	}

	// Find all entities with repository
	packageSessionEntities, err := s.packageSessionRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
		Where:  where,
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.packageSessionMapper.MapEntitiesToListResponses(
		ctx,
		s.embarkationRepository,
		s.packageSessionRepository,
		packageSessionEntities,
	)
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

	// Create response
	response := dto.PackageSessionResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create package session repository
		packageSessionRepository := factory.NewPackageSessionRepository()

		// Create embarkation repository
		embarkationRepository := factory.NewEmbarkationRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Update entity with repository
		packageSessionEntity, err := packageSessionRepository.Update(ctx, id, packageSessionEntity)
		if err != nil {
			return err
		}

		// Delete guides with repository
		if _, err := packageSessionRepository.DeleteGuides(ctx, packageSessionEntity.Id); err != nil {
			return err
		}

		// Create guides with repository
		if _, err := packageSessionRepository.CreateGuides(ctx, packageSessionEntity.Id, request.Guides); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.packageSessionMapper.MapEntityToResponse(
			ctx,
			embarkationRepository,
			packageSessionRepository,
			imageRepository,
			packageSessionEntity,
		)

		return err
	})

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
	response, err := s.packageSessionMapper.MapEntityToResponse(
		ctx,
		s.embarkationRepository,
		s.packageSessionRepository,
		s.imageRepository,
		packageSessionEntity,
	)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return response, err
}
