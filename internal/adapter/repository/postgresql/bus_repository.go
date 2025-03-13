package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type busRepositoryPostgresql struct {
	db DB
}

func NewBusRepository(db DB) repository.BusRepository {
	return busRepositoryPostgresql{db}
}

func (r busRepositoryPostgresql) Create(ctx context.Context, bus entity.Bus) (entity.Bus, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "buses" ("name", "seat", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, bus.Name, bus.Seat).
		S(`RETURNING "id", "name", "seat", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&bus.Id, &bus.Name, &bus.Seat,
		&bus.CreatedAt, &bus.UpdatedAt, &bus.DeletedAt,
	)

	return bus, err
}

func (r busRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Bus, error) {
	bus := entity.Bus{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "seat", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "buses" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&bus.Id, &bus.Name, &bus.Seat,
		&bus.CreatedAt, &bus.UpdatedAt, &bus.DeletedAt,
	)

	return bus, err
}

func (r busRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Bus, error) {
	buses := []entity.Bus{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "seat", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "buses" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return buses, err
	}

	for rows.Next() {
		bus := entity.Bus{}
		err = rows.Scan(
			&bus.Id, &bus.Name, &bus.Seat,
			&bus.CreatedAt, &bus.UpdatedAt, &bus.DeletedAt,
		)
		if err != nil {
			return buses, err
		}

		buses = append(buses, bus)
	}

	return buses, nil
}

func (r busRepositoryPostgresql) Update(ctx context.Context, id int64, bus entity.Bus) (entity.Bus, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "buses" SET "name" = $1, "seat" = $2, "updated_at" = NOW()`,
			bus.Name, bus.Seat,
		).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "seat", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&bus.Id, &bus.Name, &bus.Seat,
		&bus.CreatedAt, &bus.UpdatedAt, &bus.DeletedAt,
	)

	return bus, err
}

func (r busRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Bus, error) {
	bus := entity.Bus{}

	builder := sqlbuilder.New().
		S(`UPDATE "buses" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "seat", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&bus.Id, &bus.Name, &bus.Seat,
		&bus.CreatedAt, &bus.UpdatedAt, &bus.DeletedAt,
	)

	return bus, err
}
