package repository

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	FindById(ctx context.Context, id int64) (entity.User, error)
	FindAll(ctx context.Context, opt FindAllOptions) ([]entity.User, error)
	Update(ctx context.Context, id int64, user entity.User) (entity.User, error)
	Delete(ctx context.Context, id int64) (entity.User, error)
}
