package service

import (
	"context"
	"os"
	"path/filepath"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	serviceport "github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type imageServiceImpl struct {
	imageRepository repository.ImageRepository
	imageValidator  validator.ImageValidator
	imageMapper     mapper.ImageMapper

	unitOfWork repository.UnitOfWork
}

func NewImageService(
	imageRepository repository.ImageRepository,
	imageValidator validator.ImageValidator,
	imageMapper mapper.ImageMapper,
	unitOfWork repository.UnitOfWork,
) serviceport.ImageService {
	return imageServiceImpl{
		imageRepository,
		imageValidator,
		imageMapper,
		unitOfWork,
	}
}

func (s imageServiceImpl) CreateImage(ctx context.Context, request dto.ImageRequest) (dto.ImageResponse, error) {
	// Validate request
	if err := s.imageValidator.ValidateRequest(ctx, request); err != nil {
		return dto.ImageResponse{}, err
	}

	// Map request into entity
	imageEntity := s.imageMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.ImageResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Create entity with repository
		imageEntity, err := imageRepository.Create(ctx, imageEntity)
		if err != nil {
			return err
		}

		// Write image file into public folder
		if err := os.WriteFile(
			filepath.Join("public", imageEntity.Src),
			request.FileData,
			0644,
		); err != nil {
			return err
		}

		// Map entity into response
		response = s.imageMapper.MapEntityToResponse(ctx, imageEntity)

		return nil
	})

	return response, err
}

func (s imageServiceImpl) GetImageById(ctx context.Context, id int64) (dto.ImageResponse, error) {
	// Validate id
	if err := s.imageValidator.ValidateId(ctx, id); err != nil {
		return dto.ImageResponse{}, err
	}

	// Find entity by id with repository
	imageEntity, err := s.imageRepository.FindById(ctx, id)
	if err != nil {
		return dto.ImageResponse{}, err
	}

	// Map entity into response
	response := s.imageMapper.MapEntityToResponse(ctx, imageEntity)

	return response, nil
}

func (s imageServiceImpl) GetAllImage(ctx context.Context, request dto.GetAllImageRequest) ([]dto.ImageResponse, error) {
	// Validate request
	if err := s.imageValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	imageEntities, err := s.imageRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

	return responses, nil
}

func (s imageServiceImpl) UpdateImage(ctx context.Context, id int64, request dto.ImageRequest) (dto.ImageResponse, error) {
	// Validate id
	if err := s.imageValidator.ValidateId(ctx, id); err != nil {
		return dto.ImageResponse{}, err
	}

	// Validate request
	if err := s.imageValidator.ValidateRequest(ctx, request); err != nil {
		return dto.ImageResponse{}, err
	}

	// Map request into entity
	imageEntity := s.imageMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.ImageResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Update entity with repository
		imageEntity, err := imageRepository.Update(ctx, id, imageEntity)
		if err != nil {
			return err
		}

		// Write image file into public folder
		if err := os.WriteFile(
			filepath.Join("public", imageEntity.Src),
			request.FileData,
			0644,
		); err != nil {
			return err
		}

		// Map entity into response
		response = s.imageMapper.MapEntityToResponse(ctx, imageEntity)

		return nil
	})

	return response, err
}

func (s imageServiceImpl) DeleteImage(ctx context.Context, id int64) (dto.ImageResponse, error) {
	// Validate id
	if err := s.imageValidator.ValidateId(ctx, id); err != nil {
		return dto.ImageResponse{}, err
	}

	// Create response
	response := dto.ImageResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create image repository
		imageRepository := factory.NewImageRepository()

		// Delete entity with repository
		imageEntity, err := imageRepository.Delete(ctx, id)
		if err != nil {
			return err
		}

		// Delete image file from public folder
		if err := os.Remove(filepath.Join("public", imageEntity.Src)); err != nil {
			return err
		}

		// Map entity into response
		response = s.imageMapper.MapEntityToResponse(ctx, imageEntity)

		return nil
	})

	return response, err
}
