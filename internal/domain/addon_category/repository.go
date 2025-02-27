package addon_category

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type RepositoryFindAllOption struct {
	Limit  null.Int
	Offset null.Int
}

type Repository struct {
	db database.DB
}

func NewRepository(db database.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(ctx context.Context, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`INSERT INTO "addon_categories" ("name", "created_at", "updated_at")`).
		S(`VALUES ($1, NOW(), NOW())`, entity.Name).
		S(`RETURNING "id", "name", "created_at", "updated_at", "deleted_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) FindAll(ctx context.Context, opt RepositoryFindAllOption) ([]Entity, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "addon_categories"`).
		S(`WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.S(`LIMIT ?`, opt.Limit.Int64)
	}
	if opt.Offset.Valid {
		builder.S(`OFFSET ?`, opt.Offset.Int64)
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
			&entity.Id, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (r Repository) FindByID(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "id", "name", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "addon_categories"`).
		S(`WHERE "id" = $1`, id).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Update(ctx context.Context, id int64, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "addon_categories"`).
		S(`SET "name" = $1, "updated_at" = NOW()`, entity.Name).
		S(`WHERE "id" = $2`, id).
		S(`RETURNING "id", "name", "created_at", "updated_at", "deleted_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "addon_categories"`).
		S(`SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1`, id).
		S(`RETURNING "id", "name", "created_at", "updated_at", "deleted_at"`).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}
