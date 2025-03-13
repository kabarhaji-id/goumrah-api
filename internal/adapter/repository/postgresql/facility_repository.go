package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type facilityRepositoryPostgresql struct {
	db DB
}

func NewFacilityRepository(db DB) repository.FacilityRepository {
	return facilityRepositoryPostgresql{db}
}

func (r facilityRepositoryPostgresql) Create(ctx context.Context, facility entity.Facility) (entity.Facility, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "facilities" ("name", "icon", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, NOW(), NOW(), NULL)`, facility.Name, facility.Icon).
		S(`RETURNING "id", "name", "icon", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&facility.Id, &facility.Name, &facility.Icon,
		&facility.CreatedAt, &facility.UpdatedAt, &facility.DeletedAt,
	)

	return facility, err
}

func (r facilityRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Facility, error) {
	facility := entity.Facility{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "icon", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "facilities" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&facility.Id, &facility.Name, &facility.Icon,
		&facility.CreatedAt, &facility.UpdatedAt, &facility.DeletedAt,
	)

	return facility, err
}

func (r facilityRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Facility, error) {
	facilities := []entity.Facility{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "icon", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "facilities" WHERE "deleted_at" IS NULL`).
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
		facility := entity.Facility{}
		err = rows.Scan(
			&facility.Id, &facility.Name, &facility.Icon,
			&facility.CreatedAt, &facility.UpdatedAt, &facility.DeletedAt,
		)
		if err != nil {
			return facilities, err
		}

		facilities = append(facilities, facility)
	}

	return facilities, nil
}

func (r facilityRepositoryPostgresql) Update(ctx context.Context, id int64, facility entity.Facility) (entity.Facility, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "facilities" SET "name" = $1, "icon" = $2, "updated_at" = NOW()`,
			facility.Name, facility.Icon,
		).
		S(`WHERE "id" = $3 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "icon", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&facility.Id, &facility.Name, &facility.Icon,
		&facility.CreatedAt, &facility.UpdatedAt, &facility.DeletedAt,
	)

	return facility, err
}

func (r facilityRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Facility, error) {
	facility := entity.Facility{}

	builder := sqlbuilder.New().
		S(`UPDATE "facilities" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "icon", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&facility.Id, &facility.Name, &facility.Icon,
		&facility.CreatedAt, &facility.UpdatedAt, &facility.DeletedAt,
	)

	return facility, err
}
