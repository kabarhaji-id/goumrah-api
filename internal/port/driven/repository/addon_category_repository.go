package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type AddonCategoryRepository interface {
	Create(ctx context.Context, addonCategory entity.AddonCategory) (entity.AddonCategory, error)
	FindById(ctx context.Context, id int64) (entity.AddonCategory, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.AddonCategory, error)
	Update(ctx context.Context, id int64, addonCategory entity.AddonCategory) (entity.AddonCategory, error)
	Delete(ctx context.Context, id int64) (entity.AddonCategory, error)
}
