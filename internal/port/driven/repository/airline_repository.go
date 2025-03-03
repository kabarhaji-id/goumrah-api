package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type AirlineRepository interface {
	Create(ctx context.Context, image entity.Airline) (entity.Airline, error)
	FindById(ctx context.Context, id int64) (entity.Airline, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Airline, error)
	Update(ctx context.Context, id int64, image entity.Airline) (entity.Airline, error)
	Delete(ctx context.Context, id int64) (entity.Airline, error)
}
