package package_session

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type RepositoryFindAllOption struct {
	Limit  null.Int
	Offset null.Int

	PackageID null.Int
}

type Repository struct {
	db database.DB
}

func NewRepository(db database.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(ctx context.Context, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`INSERT INTO "package_sessions" ("package_id", "embarkation_id", "departure_date", "created_at", "updated_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW())`, entity.PackageId, entity.EmbarkationId, entity.DepartureDate).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.PackageId, &entity.EmbarkationId, &entity.DepartureDate, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) FindAll(ctx context.Context, opt RepositoryFindAllOption) ([]Entity, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at"`).
		S(`FROM "package_sessions"`).
		S(`WHERE "deleted_at" IS NULL`)
	if opt.PackageID.Valid {
		builder.S(`AND "package_id" = ?`, opt.PackageID.Int64)
	}
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit.Int64)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset.Int64)
	}
	builder.S(`ORDER BY "id" ASC`)

	query, args := builder.Build()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var entities []Entity
	for rows.Next() {
		entity := Entity{}
		if err := rows.Scan(
			&entity.Id, &entity.PackageId, &entity.EmbarkationId, &entity.DepartureDate, &entity.CreatedAt, &entity.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (r Repository) FindByID(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at"`).
		S(`FROM "package_sessions"`).
		S(`WHERE "id" = $1`, id).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.PackageId, &entity.EmbarkationId, &entity.DepartureDate, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Update(ctx context.Context, id int64, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "package_sessions"`).
		S(`SET "embarkation_id" = $1, "departure_date" = $2, "updated_at" = NOW()`, entity.EmbarkationId, entity.DepartureDate).
		S(`WHERE "id" = $3`, id).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.PackageId, &entity.EmbarkationId, &entity.DepartureDate, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "package_sessions"`).
		S(`SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1`, id).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at"`).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.PackageId, &entity.EmbarkationId, &entity.DepartureDate, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}
