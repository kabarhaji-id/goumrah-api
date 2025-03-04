package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type AddonRepository interface {
	Create(ctx context.Context, addon entity.Addon) (entity.Addon, error)
	FindById(ctx context.Context, id int64) (entity.Addon, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Addon, error)
	Update(ctx context.Context, id int64, addon entity.Addon) (entity.Addon, error)
	Delete(ctx context.Context, id int64) (entity.Addon, error)
}
