package mapper

import (
	"context"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
)

type PackageMapper struct {
	imageMapper ImageMapper
}

func NewPackageMapper(
	imageMapper ImageMapper,
) PackageMapper {
	return PackageMapper{
		imageMapper,
	}
}

func (PackageMapper) MapRequestToEntity(ctx context.Context, request dto.PackageRequest) entity.Package {
	return entity.Package{
		ThumbnailId: request.Thumbnail,
		Name:        request.Name,
		Category:    request.Category,
		Type:        request.Type,
		FastTrain:   request.FastTrain,
		Slug:        sluger.Slug(request.Name),
	}
}

func (m PackageMapper) MapEntityToResponse(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	packageRepository repository.PackageRepository,
	packageEntity entity.Package,
) (dto.PackageResponse, error) {
	thumbnailResponse := null.NewValue(dto.ImageResponse{}, false)
	if packageEntity.ThumbnailId.Valid {
		thumbnailEntity, err := imageRepository.FindById(ctx, packageEntity.ThumbnailId.Int64)
		if err != nil {
			return dto.PackageResponse{}, err
		}

		thumbnailResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, thumbnailEntity))
	}

	imageEntities, err := packageRepository.FindImages(ctx, packageEntity.Id)
	if err != nil {
		return dto.PackageResponse{}, err
	}
	imageResponses := m.imageMapper.MapEntitiesToResponses(ctx, imageEntities)

	return dto.PackageResponse{
		Id:        packageEntity.Id,
		Thumbnail: thumbnailResponse,
		Name:      packageEntity.Name,
		Category:  packageEntity.Category,
		Type:      packageEntity.Type,
		FastTrain: packageEntity.FastTrain,
		Slug:      packageEntity.Slug,
		Images:    imageResponses,
		CreatedAt: packageEntity.CreatedAt,
		UpdatedAt: packageEntity.UpdatedAt,
		DeletedAt: packageEntity.DeletedAt,
	}, nil
}

func (m PackageMapper) MapEntityToListResponse(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	packageRepository repository.PackageRepository,
	packageEntity entity.Package,
) (dto.PackageListResponse, error) {
	thumbnailResponse := null.NewValue(dto.ImageResponse{}, false)
	if packageEntity.ThumbnailId.Valid {
		thumbnailEntity, err := imageRepository.FindById(ctx, packageEntity.ThumbnailId.Int64)
		if err != nil {
			return dto.PackageListResponse{}, err
		}

		thumbnailResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, thumbnailEntity))
	}

	imageIds, err := packageRepository.FindImageIds(ctx, packageEntity.Id)
	if err != nil {
		return dto.PackageListResponse{}, err
	}

	return dto.PackageListResponse{
		Id:        packageEntity.Id,
		Thumbnail: thumbnailResponse,
		Name:      packageEntity.Name,
		Category:  packageEntity.Category,
		Type:      packageEntity.Type,
		FastTrain: packageEntity.FastTrain,
		Slug:      packageEntity.Slug,
		Images:    imageIds,
		CreatedAt: packageEntity.CreatedAt,
		UpdatedAt: packageEntity.UpdatedAt,
		DeletedAt: packageEntity.DeletedAt,
	}, nil
}

func (m PackageMapper) MapEntitiesToResponses(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	packageRepository repository.PackageRepository,
	packageEntities []entity.Package,
) ([]dto.PackageResponse, error) {
	packageResponses := make([]dto.PackageResponse, len(packageEntities))
	var err error

	for i, packageEntity := range packageEntities {
		packageResponses[i], err = m.MapEntityToResponse(ctx, imageRepository, packageRepository, packageEntity)
		if err != nil {
			return nil, err
		}
	}

	return packageResponses, nil
}

func (m PackageMapper) MapEntitiesToListResponses(
	ctx context.Context,
	imageRepository repository.ImageRepository,
	packageRepository repository.PackageRepository,
	packageEntities []entity.Package,
) ([]dto.PackageListResponse, error) {
	packageListResponses := make([]dto.PackageListResponse, len(packageEntities))
	var err error

	for i, packageEntity := range packageEntities {
		packageListResponses[i], err = m.MapEntityToListResponse(ctx, imageRepository, packageRepository, packageEntity)
		if err != nil {
			return nil, err
		}
	}

	return packageListResponses, nil
}
