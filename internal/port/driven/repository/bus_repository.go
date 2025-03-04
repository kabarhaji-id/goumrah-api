package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type BusRepository interface {
	Create(ctx context.Context, bus entity.Bus) (entity.Bus, error)
	FindById(ctx context.Context, id int64) (entity.Bus, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Bus, error)
	Update(ctx context.Context, id int64, bus entity.Bus) (entity.Bus, error)
	Delete(ctx context.Context, id int64) (entity.Bus, error)
}
