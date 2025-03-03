package mapper

import (
	"context"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionMapper struct {
	embarkationRepository repository.EmbarkationRepository
	embarkationMapper     EmbarkationMapper

	packageSessionRepository repository.PackageSessionRepository
}

func NewPackageSessionMapper(
	embarkationRepository repository.EmbarkationRepository,
	embarkationMapper EmbarkationMapper,
	packageSessionRepository repository.PackageSessionRepository,
) PackageSessionMapper {
	return PackageSessionMapper{
		embarkationRepository, embarkationMapper,
		packageSessionRepository,
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

func (m PackageSessionMapper) MapEntityToResponse(ctx context.Context, packageSessionEntity entity.PackageSession) (dto.PackageSessionResponse, error) {
	embarkationEntity, err := m.embarkationRepository.FindById(ctx, packageSessionEntity.EmbarkationId)
	if err != nil {
		return dto.PackageSessionResponse{}, err
	}

	embarkationResponse := m.embarkationMapper.MapEntityToResponse(ctx, embarkationEntity)

	return dto.PackageSessionResponse{
		Id:            packageSessionEntity.Id,
		Package:       packageSessionEntity.PackageId,
		Embarkation:   embarkationResponse,
		DepartureDate: packageSessionEntity.DepartureDate,
		CreatedAt:     packageSessionEntity.CreatedAt,
		UpdatedAt:     packageSessionEntity.UpdatedAt,
		DeletedAt:     packageSessionEntity.DeletedAt,
	}, nil
}

func (m PackageSessionMapper) MapEntitiesToResponses(ctx context.Context, packageSessionEntities []entity.PackageSession) ([]dto.PackageSessionResponse, error) {
	packageSessionResponses := make([]dto.PackageSessionResponse, len(packageSessionEntities))
	var err error

	for i, packageSessionEntity := range packageSessionEntities {
		packageSessionResponses[i], err = m.MapEntityToResponse(ctx, packageSessionEntity)
		if err != nil {
			return nil, err
		}
	}

	return packageSessionResponses, nil
}
