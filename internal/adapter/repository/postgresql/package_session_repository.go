package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type packageSessionRepositoryPostgresql struct {
	db DB
}

func NewPackageSessionRepository(db DB) repository.PackageSessionRepository {
	return packageSessionRepositoryPostgresql{db}
}

func (r packageSessionRepositoryPostgresql) Create(ctx context.Context, packageSession entity.PackageSession) (entity.PackageSession, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "package_sessions" ("package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, packageSession.PackageId, packageSession.EmbarkationId, packageSession.DepartureDate).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.PackageSession, error) {
	packageSession := entity.PackageSession{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "package_sessions" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.PackageSession, error) {
	packageSessions := []entity.PackageSession{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "package_sessions" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return packageSessions, err
	}

	for rows.Next() {
		packageSession := entity.PackageSession{}
		err = rows.Scan(
			&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate,
			&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
		)
		if err != nil {
			return packageSessions, err
		}

		packageSessions = append(packageSessions, packageSession)
	}

	return packageSessions, nil
}

func (r packageSessionRepositoryPostgresql) Update(ctx context.Context, id int64, packageSession entity.PackageSession) (entity.PackageSession, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "package_sessions" SET "package_id" = $1, "embarkation_id" = $2, "departure_date" = $3, "updated_at" = NOW()`,
			packageSession.PackageId, packageSession.EmbarkationId, packageSession.DepartureDate,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.PackageSession, error) {
	packageSession := entity.PackageSession{}

	builder := sqlbuilder.New().
		S(`UPDATE "package_sessions" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}
