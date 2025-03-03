package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type HotelRepository interface {
	Create(ctx context.Context, image entity.Hotel) (entity.Hotel, error)
	FindById(ctx context.Context, id int64) (entity.Hotel, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Hotel, error)
	Update(ctx context.Context, id int64, image entity.Hotel) (entity.Hotel, error)
	Delete(ctx context.Context, id int64) (entity.Hotel, error)
}
