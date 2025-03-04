package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type PackageSessionRepository interface {
	Create(ctx context.Context, packageSession entity.PackageSession) (entity.PackageSession, error)
	FindById(ctx context.Context, id int64) (entity.PackageSession, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.PackageSession, error)
	Update(ctx context.Context, id int64, packageSession entity.PackageSession) (entity.PackageSession, error)
	Delete(ctx context.Context, id int64) (entity.PackageSession, error)

	CreateGuides(ctx context.Context, id int64, guideIds []int64) ([]int64, error)
	FindGuides(ctx context.Context, id int64) ([]entity.Guide, error)
	FindGuideIds(ctx context.Context, id int64) ([]int64, error)
	DeleteGuides(ctx context.Context, id int64) ([]int64, error)
}
