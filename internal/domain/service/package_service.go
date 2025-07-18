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

type packageServiceImpl struct {
	packageRepository repository.PackageRepository
	packageValidator  validator.PackageValidator
	packageMapper     mapper.PackageMapper

	imageRepository repository.ImageRepository

	unitOfWork repository.UnitOfWork
}

func NewPackageService(
	packageRepository repository.PackageRepository,
	packageValidator validator.PackageValidator,
	packageMapper mapper.PackageMapper,
	imageRepository repository.ImageRepository,
	unitOfWork repository.UnitOfWork,
) serviceport.PackageService {
	return packageServiceImpl{
		packageRepository,
		packageValidator,
		packageMapper,
		imageRepository,
		unitOfWork,
	}
}

func (s packageServiceImpl) CreatePackage(ctx context.Context, request dto.PackageRequest) (dto.PackageResponse, error) {
	// Validate request
	if err := s.packageValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageResponse{}, err
	}

	// Map request into entity
	packageEntity := s.packageMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.PackageResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create package repository
		packageRepository := factory.NewPackageRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create entity with repository
		packageEntity, err := packageRepository.Create(ctx, packageEntity)
		if err != nil {
			return err
		}

		// Create images with repository
		if _, err := packageRepository.AttachImages(ctx, packageEntity.Id, request.Images); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.packageMapper.MapEntityToResponse(
			ctx,
			imageRepository,
			packageRepository,
			packageEntity,
		)

		return err
	})

	return response, err
}

func (s packageServiceImpl) GetPackageById(ctx context.Context, id int64) (dto.PackageResponse, error) {
	// Validate id
	if err := s.packageValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageResponse{}, err
	}

	// Find entity by id with repository
	packageEntity, err := s.packageRepository.FindById(ctx, id)
	if err != nil {
		return dto.PackageResponse{}, err
	}

	// Map entity into response
	response, err := s.packageMapper.MapEntityToResponse(
		ctx,
		s.imageRepository,
		s.packageRepository,
		packageEntity,
	)
	if err != nil {
		return dto.PackageResponse{}, err
	}

	return response, nil
}

func (s packageServiceImpl) GetAllPackage(ctx context.Context, request dto.GetAllPackageRequest) ([]dto.PackageListResponse, error) {
	// Validate request
	if err := s.packageValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	packageEntities, err := s.packageRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses, err := s.packageMapper.MapEntitiesToListResponses(ctx,
		s.imageRepository,
		s.packageRepository,
		packageEntities,
	)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s packageServiceImpl) UpdatePackage(ctx context.Context, id int64, request dto.PackageRequest) (dto.PackageResponse, error) {
	// Validate id
	if err := s.packageValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageResponse{}, err
	}

	// Validate request
	if err := s.packageValidator.ValidateRequest(ctx, request); err != nil {
		return dto.PackageResponse{}, err
	}

	// Map request into entity
	packageEntity := s.packageMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.PackageResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create package repository
		packageRepository := factory.NewPackageRepository()

		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Update entity with repository
		packageEntity, err := packageRepository.Update(ctx, id, packageEntity)
		if err != nil {
			return err
		}

		// Delete images with repository
		if _, err := packageRepository.DetachImages(ctx, packageEntity.Id); err != nil {
			return err
		}

		// Create images with repository
		if _, err := packageRepository.AttachImages(ctx, packageEntity.Id, request.Images); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.packageMapper.MapEntityToResponse(ctx,
			imageRepository,
			packageRepository,
			packageEntity,
		)

		return err
	})

	return response, err
}

func (s packageServiceImpl) DeletePackage(ctx context.Context, id int64) (dto.PackageResponse, error) {
	// Validate id
	if err := s.packageValidator.ValidateId(ctx, id); err != nil {
		return dto.PackageResponse{}, err
	}

	// Delete entity with repository
	packageEntity, err := s.packageRepository.Delete(ctx, id)
	if err != nil {
		return dto.PackageResponse{}, err
	}

	// Map entity into response
	response, err := s.packageMapper.MapEntityToResponse(
		ctx,
		s.imageRepository,
		s.packageRepository,
		packageEntity,
	)
	if err != nil {
		return dto.PackageResponse{}, err
	}

	return response, err
}
