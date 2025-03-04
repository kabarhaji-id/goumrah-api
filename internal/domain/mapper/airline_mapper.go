package mapper

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirlineMapper struct {
	imageMapper ImageMapper
}

func NewAirlineMapper(imageMapper ImageMapper) AirlineMapper {
	return AirlineMapper{
		imageMapper,
	}
}

func (AirlineMapper) MapRequestToEntity(ctx context.Context, request dto.AirlineRequest) entity.Airline {
	return entity.Airline{
		Name:          request.Name,
		SkytraxType:   request.SkytraxType,
		SkytraxRating: request.SkytraxRating,
		LogoId:        request.Logo,
	}
}

func (m AirlineMapper) MapEntityToResponse(ctx context.Context, imageRepository repository.ImageRepository, airlineEntity entity.Airline) (dto.AirlineResponse, error) {
	logoResponse := null.NewValue(dto.ImageResponse{}, false)
	if airlineEntity.LogoId.Valid {
		logoEntity, err := imageRepository.FindById(ctx, airlineEntity.LogoId.Int64)
		if err != nil {
			return dto.AirlineResponse{}, err
		}

		logoResponse = null.ValueFrom(m.imageMapper.MapEntityToResponse(ctx, logoEntity))
	}

	return dto.AirlineResponse{
		Id:            airlineEntity.Id,
		Name:          airlineEntity.Name,
		SkytraxType:   airlineEntity.SkytraxType,
		SkytraxRating: airlineEntity.SkytraxRating,
		Logo:          logoResponse,
		CreatedAt:     airlineEntity.CreatedAt,
		UpdatedAt:     airlineEntity.UpdatedAt,
		DeletedAt:     airlineEntity.DeletedAt,
	}, nil
}

func (m AirlineMapper) MapEntitiesToResponses(ctx context.Context, imageRepository repository.ImageRepository, airlineEntities []entity.Airline) ([]dto.AirlineResponse, error) {
	airlineResponses := make([]dto.AirlineResponse, len(airlineEntities))
	var err error

	for i, airlineEntity := range airlineEntities {
		airlineResponses[i], err = m.MapEntityToResponse(ctx, imageRepository, airlineEntity)
		if err != nil {
			return nil, err
		}
	}

	return airlineResponses, nil
}
