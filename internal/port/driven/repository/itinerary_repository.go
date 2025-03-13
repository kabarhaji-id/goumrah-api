package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type ItineraryRepository interface {
	Create(ctx context.Context, itinerary entity.Itinerary) (entity.Itinerary, error)
	FindById(ctx context.Context, id int64) (entity.Itinerary, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Itinerary, error)
	Update(ctx context.Context, id int64, itinerary entity.Itinerary) (entity.Itinerary, error)
	Delete(ctx context.Context, id int64) (entity.Itinerary, error)

	AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error)
	FindImages(ctx context.Context, id int64) ([]entity.Image, error)
	FindImageIds(ctx context.Context, id int64) ([]int64, error)
	DetachImages(ctx context.Context, id int64) ([]int64, error)
}
