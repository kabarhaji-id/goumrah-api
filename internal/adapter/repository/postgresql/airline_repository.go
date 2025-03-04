package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type airlineRepositoryPostgresql struct {
	db DB
}

func NewAirlineRepository(db DB) repository.AirlineRepository {
	return airlineRepositoryPostgresql{db}
}

func (r airlineRepositoryPostgresql) Create(ctx context.Context, airline entity.Airline) (entity.Airline, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "airlines" ("name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, airline.Name, airline.SkytraxType, airline.SkytraxRating, airline.LogoId).
		S(`RETURNING "id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airline.Id, &airline.Name, &airline.SkytraxType, &airline.SkytraxRating, &airline.LogoId,
		&airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt,
	)

	return airline, err
}

func (r airlineRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Airline, error) {
	airline := entity.Airline{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "airlines" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airline.Id, &airline.Name, &airline.SkytraxType, &airline.SkytraxRating, &airline.LogoId,
		&airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt,
	)

	return airline, err
}

func (r airlineRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Airline, error) {
	airlines := []entity.Airline{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "airlines" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return airlines, err
	}

	for rows.Next() {
		airline := entity.Airline{}
		err = rows.Scan(
			&airline.Id, &airline.Name, &airline.SkytraxType, &airline.SkytraxRating, &airline.LogoId,
			&airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt,
		)
		if err != nil {
			return airlines, err
		}

		airlines = append(airlines, airline)
	}

	return airlines, nil
}

func (r airlineRepositoryPostgresql) Update(ctx context.Context, id int64, airline entity.Airline) (entity.Airline, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "airlines" SET "name" = $1, "skytrax_type" = $2, "skytrax_rating" = $3, "logo_id" = $4, "updated_at" = NOW()`,
			airline.Name, airline.SkytraxType, airline.SkytraxRating, airline.LogoId,
		).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airline.Id, &airline.Name, &airline.SkytraxType, &airline.SkytraxRating, &airline.LogoId,
		&airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt,
	)

	return airline, err
}

func (r airlineRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Airline, error) {
	airline := entity.Airline{}

	builder := sqlbuilder.New().
		S(`UPDATE "airlines" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airline.Id, &airline.Name, &airline.SkytraxType, &airline.SkytraxRating, &airline.LogoId,
		&airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt,
	)

	return airline, err
}
