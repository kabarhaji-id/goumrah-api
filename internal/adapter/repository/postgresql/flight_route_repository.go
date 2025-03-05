package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type flightRouteRepositoryPostgresql struct {
	db DB
}

func NewFlightRouteRepository(db DB) repository.FlightRouteRepository {
	return flightRouteRepositoryPostgresql{db}
}

func (r flightRouteRepositoryPostgresql) Create(ctx context.Context, flightRoute entity.FlightRoute) (entity.FlightRoute, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "flight_routes" ("flight_id", "next_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, flightRoute.FlightId, flightRoute.NextId).
		S(`RETURNING "id", "flight_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flightRoute.Id, &flightRoute.FlightId, &flightRoute.NextId,
		&flightRoute.CreatedAt, &flightRoute.UpdatedAt, &flightRoute.DeletedAt,
	)

	return flightRoute, err
}

func (r flightRouteRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.FlightRoute, error) {
	flightRoute := entity.FlightRoute{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "flight_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "flight_routes" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flightRoute.Id, &flightRoute.FlightId, &flightRoute.NextId,
		&flightRoute.CreatedAt, &flightRoute.UpdatedAt, &flightRoute.DeletedAt,
	)

	return flightRoute, err
}

func (r flightRouteRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.FlightRoute, error) {
	facilities := []entity.FlightRoute{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "flight_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "flight_routes" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return facilities, err
	}

	for rows.Next() {
		flightRoute := entity.FlightRoute{}
		err = rows.Scan(
			&flightRoute.Id, &flightRoute.FlightId, &flightRoute.NextId,
			&flightRoute.CreatedAt, &flightRoute.UpdatedAt, &flightRoute.DeletedAt,
		)
		if err != nil {
			return facilities, err
		}

		facilities = append(facilities, flightRoute)
	}

	return facilities, nil
}

func (r flightRouteRepositoryPostgresql) Update(ctx context.Context, id int64, flightRoute entity.FlightRoute) (entity.FlightRoute, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "flight_routes" SET "flight_id" = $1, "next_id" = $2, "updated_at" = NOW()`,
			flightRoute.FlightId, flightRoute.NextId,
		).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "flight_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flightRoute.Id, &flightRoute.FlightId, &flightRoute.NextId,
		&flightRoute.CreatedAt, &flightRoute.UpdatedAt, &flightRoute.DeletedAt,
	)

	return flightRoute, err
}

func (r flightRouteRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.FlightRoute, error) {
	flightRoute := entity.FlightRoute{}

	builder := sqlbuilder.New().
		S(`UPDATE "flight_routes" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "flight_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&flightRoute.Id, &flightRoute.FlightId, &flightRoute.NextId,
		&flightRoute.CreatedAt, &flightRoute.UpdatedAt, &flightRoute.DeletedAt,
	)

	return flightRoute, err
}
