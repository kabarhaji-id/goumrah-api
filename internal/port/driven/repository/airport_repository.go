package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type AirportRepository interface {
	Create(ctx context.Context, image entity.Airport) (entity.Airport, error)
	FindById(ctx context.Context, id int64) (entity.Airport, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Airport, error)
	Update(ctx context.Context, id int64, image entity.Airport) (entity.Airport, error)
	Delete(ctx context.Context, id int64) (entity.Airport, error)
}
