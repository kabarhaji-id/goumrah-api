package mapper

import (
	"context"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionMapper struct {
	embarkationMapper EmbarkationMapper
	guideMapper       GuideMapper
}

func NewPackageSessionMapper(embarkationMapper EmbarkationMapper, guideMapper GuideMapper) PackageSessionMapper {
	return PackageSessionMapper{
		embarkationMapper, guideMapper,
	}
}

func (PackageSessionMapper) MapRequestToEntity(ctx context.Context, request dto.PackageSessionRequest) entity.PackageSession {
	departureDate, _ := time.Parse("02/01/2006", request.DepartureDate)

	return entity.PackageSession{
		PackageId:     request.Package,
		EmbarkationId: request.Embarkation,
		DepartureDate: departureDate,
	}
}

func (m PackageSessionMapper) MapEntityToResponse(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	imageRepository repository.ImageRepository,
	packageSessionEntity entity.PackageSession,
) (dto.PackageSessionResponse, error) {
	embarkationEntity, err := embarkationRepository.FindById(ctx, packageSessionEntity.EmbarkationId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	embarkationResponse := m.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	guideEntities, err := packageSessionRepository.FindGuides(ctx, packageSessionEntity.Id)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}
	guideResponses, err := m.guideMapper.MapEntitiesToResponses(ctx, imageRepository, guideEntities)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	return dto.PackageSessionResponse{
		Id:            packageSessionEntity.Id,
		Package:       packageSessionEntity.PackageId,
		Embarkation:   embarkationResponse,
		DepartureDate: packageSessionEntity.DepartureDate,
		Guides:        guideResponses,
		CreatedAt:     packageSessionEntity.CreatedAt,
		UpdatedAt:     packageSessionEntity.UpdatedAt,
		DeletedAt:     packageSessionEntity.DeletedAt,
	}, nil
}

func (m PackageSessionMapper) MapEntityToListResponse(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	packageSessionEntity entity.PackageSession,
) (dto.PackageSessionListResponse, error) {
	embarkationEntity, err := embarkationRepository.FindById(ctx, packageSessionEntity.EmbarkationId)
	if err != nil {
		return dto.PackageSessionListResponse{}, err
	}
	embarkationResponse := m.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	guideIds, err := packageSessionRepository.FindGuideIds(ctx, packageSessionEntity.Id)
	if err != nil {
		return dto.PackageSessionListResponse{}, err
	}

	return dto.PackageSessionListResponse{
		Id:            packageSessionEntity.Id,
		Package:       packageSessionEntity.PackageId,
		Embarkation:   embarkationResponse,
		DepartureDate: packageSessionEntity.DepartureDate,
		Guides:        guideIds,
		CreatedAt:     packageSessionEntity.CreatedAt,
		UpdatedAt:     packageSessionEntity.UpdatedAt,
		DeletedAt:     packageSessionEntity.DeletedAt,
	}, nil
}

func (m PackageSessionMapper) MapEntitiesToResponses(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	imageRepository repository.ImageRepository,
	packageSessionEntities []entity.PackageSession,
) ([]dto.PackageSessionResponse, error) {
	packageSessionResponses := make([]dto.PackageSessionResponse, len(packageSessionEntities))
	var err error

	for i, packageSessionEntity := range packageSessionEntities {
		packageSessionResponses[i], err = m.MapEntityToResponse(
			ctx, embarkationRepository, packageSessionRepository, imageRepository, packageSessionEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return packageSessionResponses, nil
}

func (m PackageSessionMapper) MapEntitiesToListResponses(
	ctx context.Context,
	embarkationRepository repository.EmbarkationRepository,
	packageSessionRepository repository.PackageSessionRepository,
	packageSessionEntities []entity.PackageSession,
) ([]dto.PackageSessionListResponse, error) {
	packageSessionListResponses := make([]dto.PackageSessionListResponse, len(packageSessionEntities))
	var err error

	for i, packageSessionEntity := range packageSessionEntities {
		packageSessionListResponses[i], err = m.MapEntityToListResponse(
			ctx,
			embarkationRepository,
			packageSessionRepository,
			packageSessionEntity,
		)
		if err != nil {
			return nil, err
		}
	}

	return packageSessionListResponses, nil
}
