package mapper

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
)

type PackageMapper struct {
	imageRepository repository.ImageRepository
	imageMapper     ImageMapper

	packageRepository repository.PackageRepository
}

func NewPackageMapper(
	imageRepository repository.ImageRepository,
	imageMapper ImageMapper,
	packageRepository repository.PackageRepository,
) PackageMapper {
	return PackageMapper{
		imageRepository, imageMapper,
		packageRepository,
	}
}

func (PackageMapper) MapRequestToEntity(ctx context.Context, request dto.PackageRequest) entity.Package {
	return entity.Package{
		ThumbnailId:   request.Thumbnail,
		Name:          request.Name,
		Description:   request.Description,
		IsActive:      request.IsActive,
		Category:      request.Category,
		Type:          request.Type,
		Slug:          sluger.Slug(request.Name),
		IsRecommended: request.IsRecommended,
	}
}

func (m PackageMapper) MapEntityToResponse(ctx context.Context, packageEntity entity.Package) (dto.PackageResponse, error) {
	thumbnailResponse := null.NewValue(dto.ImageResponse{}, false)
	if packageEntity.ThumbnailId.Valid {
		thumbnailEntity, err := m.imageRepository.FindById(ctx, packageEntity.ThumbnailId.Int64)
		if err != nil {
			return dto.PackageResponse{}, err
		}

		thumbnailResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, thumbnailEntity))
	}

	imageEntities, err := m.packageRepository.FindImages(ctx, packageEntity.Id)
	if err != nil {
		return dto.PackageResponse{}, err
	}
	imageResponses := m.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

	return dto.PackageResponse{
		Id:            packageEntity.Id,
		Thumbnail:     thumbnailResponse,
		Name:          packageEntity.Name,
		Description:   packageEntity.Description,
		IsActive:      packageEntity.IsActive,
		Category:      packageEntity.Category,
		Type:          packageEntity.Type,
		Slug:          packageEntity.Slug,
		IsRecommended: packageEntity.IsRecommended,
		Images:        imageResponses,
		CreatedAt:     packageEntity.CreatedAt,
		UpdatedAt:     packageEntity.UpdatedAt,
		DeletedAt:     packageEntity.DeletedAt,
	}, nil
}

func (m PackageMapper) MapEntityToListResponse(ctx context.Context, packageEntity entity.Package) (dto.PackageListResponse, error) {
	thumbnailResponse := null.NewValue(dto.ImageResponse{}, false)
	if packageEntity.ThumbnailId.Valid {
		thumbnailEntity, err := m.imageRepository.FindById(ctx, packageEntity.ThumbnailId.Int64)
		if err != nil {
			return dto.PackageListResponse{}, err
		}

		thumbnailResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, thumbnailEntity))
	}

	imageIds, err := m.packageRepository.FindImageIds(ctx, packageEntity.Id)
	if err != nil {
		return dto.PackageListResponse{}, err
	}

	return dto.PackageListResponse{
		Id:            packageEntity.Id,
		Thumbnail:     thumbnailResponse,
		Name:          packageEntity.Name,
		Description:   packageEntity.Description,
		IsActive:      packageEntity.IsActive,
		Category:      packageEntity.Category,
		Type:          packageEntity.Type,
		Slug:          packageEntity.Slug,
		IsRecommended: packageEntity.IsRecommended,
		Images:        imageIds,
		CreatedAt:     packageEntity.CreatedAt,
		UpdatedAt:     packageEntity.UpdatedAt,
		DeletedAt:     packageEntity.DeletedAt,
	}, nil
}

func (m PackageMapper) MapEntitiesToResponses(ctx context.Context, packageEntities []entity.Package) ([]dto.PackageResponse, error) {
	packageResponses := make([]dto.PackageResponse, len(packageEntities))
	var err error

	for i, packageEntity := range packageEntities {
		packageResponses[i], err = m.MapEntityToResponse(ctx, packageEntity)
		if err != nil {
			return nil, err
		}
	}

	return packageResponses, nil
}

func (m PackageMapper) MapEntitiesToListResponses(ctx context.Context, packageEntities []entity.Package) ([]dto.PackageListResponse, error) {
	packageListResponses := make([]dto.PackageListResponse, len(packageEntities))
	var err error

	for i, packageEntity := range packageEntities {
		packageListResponses[i], err = m.MapEntityToListResponse(ctx, packageEntity)
		if err != nil {
			return nil, err
		}
	}

	return packageListResponses, nil
}
