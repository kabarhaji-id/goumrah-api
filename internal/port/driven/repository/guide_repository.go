package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type GuideRepository interface {
	Create(ctx context.Context, guide entity.Guide) (entity.Guide, error)
	FindById(ctx context.Context, id int64) (entity.Guide, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Guide, error)
	Update(ctx context.Context, id int64, guide entity.Guide) (entity.Guide, error)
	Delete(ctx context.Context, id int64) (entity.Guide, error)
}
