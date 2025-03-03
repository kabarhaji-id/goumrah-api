package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type airportRepositoryPostgresql struct {
	db DB
}

func NewAirportRepository(db DB) repository.AirportRepository {
	return airportRepositoryPostgresql{db}
}

func (r airportRepositoryPostgresql) Create(ctx context.Context, airport entity.Airport) (entity.Airport, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "airports" ("city", "name", "code", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, airport.City, airport.Name, airport.Code).
		S(`RETURNING "id", "city", "name", "code", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airport.Id, &airport.City, &airport.Name, &airport.Code,
		&airport.CreatedAt, &airport.UpdatedAt, &airport.DeletedAt,
	)

	return airport, err
}

func (r airportRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Airport, error) {
	airport := entity.Airport{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "city", "name", "code", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "airports" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airport.Id, &airport.City, &airport.Name, &airport.Code,
		&airport.CreatedAt, &airport.UpdatedAt, &airport.DeletedAt,
	)

	return airport, err
}

func (r airportRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Airport, error) {
	airports := []entity.Airport{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "city", "name", "code", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "airports" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return airports, err
	}

	for rows.Next() {
		airport := entity.Airport{}
		err = rows.Scan(
			&airport.Id, &airport.City, &airport.Name, &airport.Code,
			&airport.CreatedAt, &airport.UpdatedAt, &airport.DeletedAt,
		)
		if err != nil {
			return airports, err
		}

		airports = append(airports, airport)
	}

	return airports, nil
}

func (r airportRepositoryPostgresql) Update(ctx context.Context, id int64, airport entity.Airport) (entity.Airport, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "airports" SET "city" = $1, "name" = $2, "code" = $3, "updated_at" = NOW()`,
			airport.City, airport.Name, airport.Code,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "city", "name", "code", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airport.Id, &airport.City, &airport.Name, &airport.Code,
		&airport.CreatedAt, &airport.UpdatedAt, &airport.DeletedAt,
	)

	return airport, err
}

func (r airportRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Airport, error) {
	airport := entity.Airport{}

	builder := sqlbuilder.New().
		S(`UPDATE "airports" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "city", "name", "code", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&airport.Id, &airport.City, &airport.Name, &airport.Code,
		&airport.CreatedAt, &airport.UpdatedAt, &airport.DeletedAt,
	)

	return airport, err
}
