package hotel

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
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
		S(`INSERT INTO "hotels" ("name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())`, entity.Name, entity.Rating, entity.Map, entity.Address, entity.Distance, entity.Review, entity.Description, entity.Location, sluger.Slug(entity.Name)).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.Rating, &entity.Map, &entity.Address, &entity.Distance, &entity.Review, &entity.Description, &entity.Location, &entity.Slug, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) FindAll(ctx context.Context, opt RepositoryFindAllOption) ([]Entity, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at"`).
		S(`FROM "hotels"`).
		S(`WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
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
			&entity.Id, &entity.Name, &entity.Rating, &entity.Map, &entity.Address, &entity.Distance, &entity.Review, &entity.Description, &entity.Location, &entity.Slug, &entity.CreatedAt, &entity.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (r Repository) FindById(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at"`).
		S(`FROM "hotels"`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.Rating, &entity.Map, &entity.Address, &entity.Distance, &entity.Review, &entity.Description, &entity.Location, &entity.Slug, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Count(ctx context.Context) (int, error) {
	query, args := sqlbuilder.New().
		S(`SELECT COUNT(*)`).
		S(`FROM "hotels"`).
		S(`WHERE "deleted_at" IS NULL`).
		Build()

	var count int
	if err := r.db.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r Repository) Update(ctx context.Context, id int64, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "hotels"`).
		S(
			`SET "name" = $1, "rating" = $2, "map" = $3, "address" = $4, "distance" = $5, "review" = $6, "description" = $7, "location" = $8, "slug" = $9, "updated_at" = NOW()`,
			entity.Name, entity.Rating, entity.Map, entity.Address, entity.Distance, entity.Review, entity.Description, entity.Location, sluger.Slug(entity.Name),
		).
		S(`WHERE "id" = $10 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.Rating, &entity.Map, &entity.Address, &entity.Distance, &entity.Review, &entity.Description, &entity.Location, &entity.Slug, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "hotels"`).
		S(`SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at"`).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.Name, &entity.Rating, &entity.Map, &entity.Address, &entity.Distance, &entity.Review, &entity.Description, &entity.Location, &entity.Slug, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}
