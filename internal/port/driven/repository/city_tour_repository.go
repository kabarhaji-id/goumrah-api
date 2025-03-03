package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type CityTourRepository interface {
	Create(ctx context.Context, image entity.CityTour) (entity.CityTour, error)
	FindById(ctx context.Context, id int64) (entity.CityTour, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.CityTour, error)
	Update(ctx context.Context, id int64, image entity.CityTour) (entity.CityTour, error)
	Delete(ctx context.Context, id int64) (entity.CityTour, error)
}
