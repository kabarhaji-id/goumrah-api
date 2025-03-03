package mapper

import (
	"context"
	"crypto/rand"

	"github.com/gabriel-vasile/mimetype"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ImageMapper struct {
}

func NewImageMapper() ImageMapper {
	return ImageMapper{}
}

func (ImageMapper) MapRequestToEntity(ctx context.Context, request dto.ImageRequest) entity.Image {
	// Detect file mime from file data
	fileMime := mimetype.Detect(request.FileData)

	// Generate src from random (crypto) text and file extension from file mime
	src := rand.Text() + fileMime.Extension()

	return entity.Image{
		Src:      src,
		Alt:      request.Alt,
		Category: request.Category,
		Title:    request.Title,
	}
}

func (ImageMapper) MapEntityToResponse(ctx context.Context, imageEntity entity.Image) dto.ImageResponse {
	return dto.ImageResponse{
		Id:        imageEntity.Id,
		Src:       imageEntity.Src,
		Alt:       imageEntity.Alt,
		Category:  imageEntity.Category,
		Title:     imageEntity.Title,
		CreatedAt: imageEntity.CreatedAt,
		UpdatedAt: imageEntity.UpdatedAt,
		DeletedAt: imageEntity.DeletedAt,
	}
}

func (m ImageMapper) MapEntitiesToResponses(ctx context.Context, imageEntities []entity.Image) []dto.ImageResponse {
	imageResponses := make([]dto.ImageResponse, len(imageEntities))

	for i, imageEntity := range imageEntities {
		imageResponses[i] = m.MapEntityToResponse(ctx, imageEntity)
	}

	return imageResponses
}
