package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type PackageRepository interface {
	Create(ctx context.Context, image entity.Package) (entity.Package, error)
	FindById(ctx context.Context, id int64) (entity.Package, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Package, error)
	Update(ctx context.Context, id int64, image entity.Package) (entity.Package, error)
	Delete(ctx context.Context, id int64) (entity.Package, error)

	CreateImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error)
	FindImages(ctx context.Context, id int64) ([]entity.Image, error)
	FindImageIds(ctx context.Context, id int64) ([]int64, error)
	DeleteImages(ctx context.Context, id int64) ([]int64, error)
}
