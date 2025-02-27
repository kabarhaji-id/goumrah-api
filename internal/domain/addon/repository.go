package addon

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type RepositoryFindAllOption struct {
	Limit  null.Int
	Offset null.Int

	PackageId null.Int
}

type RepositoryCountOption struct {
	PackageId null.Int
}

type Repository struct {
	db database.DB
}

func NewRepository(db database.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(ctx context.Context, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`INSERT INTO "addons" ("category_id", "name", "price", "created_at", "updated_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW())`, entity.CategoryId, entity.Name, entity.Price).
		S(`RETURNING "id", "category_id", "name", "price", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.CategoryId, &entity.Name, &entity.Price, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) FindAll(ctx context.Context, opt RepositoryFindAllOption) ([]Entity, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "category_id", "name", "price", "created_at", "updated_at"`).
		S(`FROM "addons"`).
		S(`WHERE "deleted_at" IS NULL`)
	if opt.PackageId.Valid {
		builder.SA(`AND "package_id" = ?`, opt.PackageId.Int64)
	}
	builder.S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit.Int64)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset.Int64)
	}

	query, args := builder.Build()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var entities []Entity
	for rows.Next() {
		entity := Entity{}
		if err := rows.Scan(
			&entity.Id, &entity.CategoryId, &entity.Name, &entity.Price, &entity.CreatedAt, &entity.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (r Repository) FindById(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "id", "category_id", "name", "price", "created_at", "updated_at"`).
		S(`FROM "addons"`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.CategoryId, &entity.Name, &entity.Price, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Count(ctx context.Context, opt RepositoryCountOption) (int, error) {
	builder := sqlbuilder.New().
		S(`SELECT COUNT(*)`).
		S(`FROM "addons"`).
		S(`WHERE "deleted_at" IS NULL`)
	if opt.PackageId.Valid {
		builder.S(`AND "package_id" = $1`, opt.PackageId.Int64)
	}

	query, args := builder.Build()

	var count int
	if err := r.db.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r Repository) Update(ctx context.Context, id int64, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "addons"`).
		S(`SET "category_id" = $1, "name" = $2, "price" = $3, "updated_at" = NOW()`, entity.CategoryId, entity.Name, entity.Price).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "category_id", "name", "price", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.CategoryId, &entity.Name, &entity.Price, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "addons"`).
		S(`SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "category_id", "name", "price", "created_at", "updated_at"`).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.CategoryId, &entity.Name, &entity.Price, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}
