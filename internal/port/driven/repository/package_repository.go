package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type PackageRepository interface {
	Create(ctx context.Context, pkg entity.Package) (entity.Package, error)
	FindById(ctx context.Context, id int64) (entity.Package, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Package, error)
	Update(ctx context.Context, id int64, pkg entity.Package) (entity.Package, error)
	Delete(ctx context.Context, id int64) (entity.Package, error)

	AttachImages(ctx context.Context, id int64, pkgIds []int64) ([]int64, error)
	FindImages(ctx context.Context, id int64) ([]entity.Image, error)
	FindImageIds(ctx context.Context, id int64) ([]int64, error)
	DetachImages(ctx context.Context, id int64) ([]int64, error)
}
