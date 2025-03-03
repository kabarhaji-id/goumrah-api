package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type PackageSessionRepository interface {
	Create(ctx context.Context, image entity.PackageSession) (entity.PackageSession, error)
	FindById(ctx context.Context, id int64) (entity.PackageSession, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.PackageSession, error)
	Update(ctx context.Context, id int64, image entity.PackageSession) (entity.PackageSession, error)
	Delete(ctx context.Context, id int64) (entity.PackageSession, error)
}
