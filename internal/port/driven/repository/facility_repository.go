package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type FacilityRepository interface {
	Create(ctx context.Context, image entity.Facility) (entity.Facility, error)
	FindById(ctx context.Context, id int64) (entity.Facility, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Facility, error)
	Update(ctx context.Context, id int64, image entity.Facility) (entity.Facility, error)
	Delete(ctx context.Context, id int64) (entity.Facility, error)
}
