package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type cityTourRepositoryPostgresql struct {
	db DB
}

func NewCityTourRepository(db DB) repository.CityTourRepository {
	return cityTourRepositoryPostgresql{db}
}

func (r cityTourRepositoryPostgresql) Create(ctx context.Context, cityTour entity.CityTour) (entity.CityTour, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "city_tours" ("name", "city", "description", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, cityTour.Name, cityTour.City, cityTour.Description).
		S(`RETURNING "id", "name", "city", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&cityTour.Id, &cityTour.Name, &cityTour.City, &cityTour.Description,
		&cityTour.CreatedAt, &cityTour.UpdatedAt, &cityTour.DeletedAt,
	)

	return cityTour, err
}

func (r cityTourRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.CityTour, error) {
	cityTour := entity.CityTour{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "city", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "city_tours" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&cityTour.Id, &cityTour.Name, &cityTour.City, &cityTour.Description,
		&cityTour.CreatedAt, &cityTour.UpdatedAt, &cityTour.DeletedAt,
	)

	return cityTour, err
}

func (r cityTourRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.CityTour, error) {
	cityTours := []entity.CityTour{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "city", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "city_tours" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return cityTours, err
	}

	for rows.Next() {
		cityTour := entity.CityTour{}
		err = rows.Scan(
			&cityTour.Id, &cityTour.Name, &cityTour.City, &cityTour.Description,
			&cityTour.CreatedAt, &cityTour.UpdatedAt, &cityTour.DeletedAt,
		)
		if err != nil {
			return cityTours, err
		}

		cityTours = append(cityTours, cityTour)
	}

	return cityTours, nil
}

func (r cityTourRepositoryPostgresql) Update(ctx context.Context, id int64, cityTour entity.CityTour) (entity.CityTour, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "city_tours" SET "name" = $1, "city" = $2, "description" = $3, "updated_at" = NOW()`,
			cityTour.Name, cityTour.City, cityTour.Description,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "city", "description", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&cityTour.Id, &cityTour.Name, &cityTour.City, &cityTour.Description,
		&cityTour.CreatedAt, &cityTour.UpdatedAt, &cityTour.DeletedAt,
	)

	return cityTour, err
}

func (r cityTourRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.CityTour, error) {
	cityTour := entity.CityTour{}

	builder := sqlbuilder.New().
		S(`UPDATE "city_tours" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "city", "description", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&cityTour.Id, &cityTour.Name, &cityTour.City, &cityTour.Description,
		&cityTour.CreatedAt, &cityTour.UpdatedAt, &cityTour.DeletedAt,
	)

	return cityTour, err
}
