package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type embarkationRepositoryPostgresql struct {
	db DB
}

func NewEmbarkationRepository(db DB) repository.EmbarkationRepository {
	return embarkationRepositoryPostgresql{db}
}

func (r embarkationRepositoryPostgresql) Create(ctx context.Context, embarkation entity.Embarkation) (entity.Embarkation, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "embarkations" ("name", "latitude", "longitude", "slug", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, embarkation.Name, embarkation.Latitude, embarkation.Longitude, embarkation.Slug).
		S(`RETURNING "id", "name", "latitude", "longitude", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&embarkation.Id, &embarkation.Name, &embarkation.Latitude, &embarkation.Longitude, &embarkation.Slug,
		&embarkation.CreatedAt, &embarkation.UpdatedAt, &embarkation.DeletedAt,
	)

	return embarkation, err
}

func (r embarkationRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Embarkation, error) {
	embarkation := entity.Embarkation{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "latitude", "longitude", "slug", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "embarkations" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&embarkation.Id, &embarkation.Name, &embarkation.Latitude, &embarkation.Longitude, &embarkation.Slug,
		&embarkation.CreatedAt, &embarkation.UpdatedAt, &embarkation.DeletedAt,
	)

	return embarkation, err
}

func (r embarkationRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Embarkation, error) {
	embarkations := []entity.Embarkation{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "latitude", "longitude", "slug", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "embarkations" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return embarkations, err
	}

	for rows.Next() {
		embarkation := entity.Embarkation{}
		err = rows.Scan(
			&embarkation.Id, &embarkation.Name, &embarkation.Latitude, &embarkation.Longitude, &embarkation.Slug,
			&embarkation.CreatedAt, &embarkation.UpdatedAt, &embarkation.DeletedAt,
		)
		if err != nil {
			return embarkations, err
		}

		embarkations = append(embarkations, embarkation)
	}

	return embarkations, nil
}

func (r embarkationRepositoryPostgresql) Update(ctx context.Context, id int64, embarkation entity.Embarkation) (entity.Embarkation, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "embarkations" SET "name" = $1, "latitude" = $2, "longitude" = $3, "slug" = $4, "updated_at" = NOW()`,
			embarkation.Name, embarkation.Latitude, embarkation.Longitude, embarkation.Slug,
		).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "latitude", "longitude", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&embarkation.Id, &embarkation.Name, &embarkation.Latitude, &embarkation.Longitude, &embarkation.Slug,
		&embarkation.CreatedAt, &embarkation.UpdatedAt, &embarkation.DeletedAt,
	)

	return embarkation, err
}

func (r embarkationRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Embarkation, error) {
	embarkation := entity.Embarkation{}

	builder := sqlbuilder.New().
		S(`UPDATE "embarkations" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "latitude", "longitude", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&embarkation.Id, &embarkation.Name, &embarkation.Latitude, &embarkation.Longitude, &embarkation.Slug,
		&embarkation.CreatedAt, &embarkation.UpdatedAt, &embarkation.DeletedAt,
	)

	return embarkation, err
}
