package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type ItineraryDayRepository interface {
	Create(ctx context.Context, itineraryDay entity.ItineraryDay) (entity.ItineraryDay, error)
	FindById(ctx context.Context, id int64) (entity.ItineraryDay, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.ItineraryDay, error)
	Update(ctx context.Context, id int64, itineraryDay entity.ItineraryDay) (entity.ItineraryDay, error)
	Delete(ctx context.Context, id int64) (entity.ItineraryDay, error)
}
