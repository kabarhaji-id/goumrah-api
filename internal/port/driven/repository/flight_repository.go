package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type FlightRepository interface {
	Create(ctx context.Context, flight entity.Flight) (entity.Flight, error)
	FindById(ctx context.Context, id int64) (entity.Flight, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Flight, error)
	Update(ctx context.Context, id int64, flight entity.Flight) (entity.Flight, error)
	Delete(ctx context.Context, id int64) (entity.Flight, error)
}
