package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type ImageRepository interface {
	Create(ctx context.Context, image entity.Image) (entity.Image, error)
	FindById(ctx context.Context, id int64) (entity.Image, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Image, error)
	Update(ctx context.Context, id int64, image entity.Image) (entity.Image, error)
	Delete(ctx context.Context, id int64) (entity.Image, error)
}
