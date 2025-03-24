package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type HotelRepository interface {
	Create(ctx context.Context, hotel entity.Hotel) (entity.Hotel, error)
	FindById(ctx context.Context, id int64) (entity.Hotel, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Hotel, error)
	Update(ctx context.Context, id int64, hotel entity.Hotel) (entity.Hotel, error)
	Delete(ctx context.Context, id int64) (entity.Hotel, error)

	AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error)
	FindImages(ctx context.Context, id int64) ([]entity.Image, error)
	FindImageIds(ctx context.Context, id int64) ([]int64, error)
	DetachImages(ctx context.Context, id int64) ([]int64, error)
}
