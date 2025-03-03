package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type EmbarkationRepository interface {
	Create(ctx context.Context, image entity.Embarkation) (entity.Embarkation, error)
	FindById(ctx context.Context, id int64) (entity.Embarkation, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.Embarkation, error)
	Update(ctx context.Context, id int64, image entity.Embarkation) (entity.Embarkation, error)
	Delete(ctx context.Context, id int64) (entity.Embarkation, error)
}
