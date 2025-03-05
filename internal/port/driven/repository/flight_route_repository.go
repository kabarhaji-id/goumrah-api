package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type FlightRouteRepository interface {
	Create(ctx context.Context, flightRoute entity.FlightRoute) (entity.FlightRoute, error)
	FindById(ctx context.Context, id int64) (entity.FlightRoute, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.FlightRoute, error)
	Update(ctx context.Context, id int64, flightRoute entity.FlightRoute) (entity.FlightRoute, error)
	Delete(ctx context.Context, id int64) (entity.FlightRoute, error)
}
