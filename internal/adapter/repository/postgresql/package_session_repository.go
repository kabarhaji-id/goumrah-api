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
		S(`INSERT INTO "package_sessions" ("package_id", "embarkation_id", "departure_date", "departure_flight_route_id", "return_flight_route_id", "created_at", "updated_at", "deleted_at")`).
		S(
			`VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), NULL)`,
			packageSession.PackageId, packageSession.EmbarkationId, packageSession.DepartureDate, packageSession.DepartureFlightRouteId, packageSession.ReturnFlightRouteId,
		).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "departure_flight_route_id", "return_flight_route_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate, &packageSession.DepartureFlightRouteId, &packageSession.ReturnFlightRouteId,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.PackageSession, error) {
	packageSession := entity.PackageSession{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "package_id", "embarkation_id", "departure_date", "departure_flight_route_id", "return_flight_route_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "package_sessions" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate, &packageSession.DepartureFlightRouteId, &packageSession.ReturnFlightRouteId,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.PackageSession, error) {
	packageSessions := []entity.PackageSession{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "package_id", "embarkation_id", "departure_date", "departure_flight_route_id", "return_flight_route_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "package_sessions" WHERE "deleted_at" IS NULL`)
	if packageId, ok := opt.Where["package_id"].(int64); ok {
		builder.SA(`AND "package_id" = ?`, packageId)
	}
	builder.S(`ORDER BY "id" ASC`)
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
			&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate, &packageSession.DepartureFlightRouteId, &packageSession.ReturnFlightRouteId,
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
			`UPDATE "package_sessions" SET "package_id" = $1, "embarkation_id" = $2, "departure_date" = $3, "departure_flight_route_id" = $4, "return_flight_route_id" = $5, "updated_at" = NOW()`,
			packageSession.PackageId, packageSession.EmbarkationId, packageSession.DepartureDate, packageSession.DepartureFlightRouteId, packageSession.ReturnFlightRouteId,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "departure_flight_route_id", "return_flight_route_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate, &packageSession.DepartureFlightRouteId, &packageSession.ReturnFlightRouteId,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.PackageSession, error) {
	packageSession := entity.PackageSession{}

	builder := sqlbuilder.New().
		S(`UPDATE "package_sessions" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "package_id", "embarkation_id", "departure_date", "departure_flight_route_id", "return_flight_route_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&packageSession.Id, &packageSession.PackageId, &packageSession.EmbarkationId, &packageSession.DepartureDate, &packageSession.DepartureFlightRouteId, &packageSession.ReturnFlightRouteId,
		&packageSession.CreatedAt, &packageSession.UpdatedAt, &packageSession.DeletedAt,
	)

	return packageSession, err
}

func (r packageSessionRepositoryPostgresql) AttachGuides(ctx context.Context, id int64, guideIds []int64) ([]int64, error) {
	if len(guideIds) == 0 {
		return []int64{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "package_session_guides" ("package_session_id", "guide_id", "created_at", "updated_at", "deleted_at") VALUES`)

	guideIdsLen := len(guideIds)
	for index, guideId := range guideIds {
		builder.SA(`(?, ?, NOW(), NOW(), NULL)`, id, guideId)
		if index+1 < guideIdsLen {
			builder.S(",")
		}
	}

	_, err := r.db.Exec(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	return guideIds, nil
}

func (r packageSessionRepositoryPostgresql) FindGuides(ctx context.Context, id int64) ([]entity.Guide, error) {
	guides := []entity.Guide{}

	builder := sqlbuilder.New().
		S(`SELECT "guides"."id", "guides"."avatar_id", "guides"."name", "guides"."type", "guides"."description", "guides"."created_at", "guides"."updated_at", "guides"."deleted_at"`).
		S(`FROM "package_session_guides"`).
		S(`INNER JOIN "package_sessions" ON "package_sessions"."id" = "package_session_guides"."package_session_id"`).
		S(`INNER JOIN "guides" ON "guides"."id" = "package_session_guides"."guide_id"`).
		S(`WHERE "package_session_guides"."package_session_id" = $1 AND "package_session_guides"."deleted_at" IS NULL AND "package_sessions"."deleted_at" IS NULL AND "guides"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return guides, err
	}

	for rows.Next() {
		guide := entity.Guide{}
		err = rows.Scan(
			&guide.Id, &guide.AvatarId, &guide.Name, &guide.Type, &guide.Description,
			&guide.CreatedAt, &guide.UpdatedAt, &guide.DeletedAt,
		)
		if err != nil {
			return guides, err
		}

		guides = append(guides, guide)
	}

	return guides, nil
}

func (r packageSessionRepositoryPostgresql) FindGuideIds(ctx context.Context, id int64) ([]int64, error) {
	guideIds := []int64{}

	builder := sqlbuilder.New().
		S(`SELECT "guides"."id"`).
		S(`FROM "package_session_guides"`).
		S(`INNER JOIN "package_sessions" ON "package_sessions"."id" = "package_session_guides"."package_session_id"`).
		S(`INNER JOIN "guides" ON "guides"."id" = "package_session_guides"."guide_id"`).
		S(`WHERE "package_session_guides"."package_session_id" = $1 AND "package_session_guides"."deleted_at" IS NULL AND "package_sessions"."deleted_at" IS NULL AND "guides"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return guideIds, err
	}

	for rows.Next() {
		var guideId int64
		err = rows.Scan(&guideId)
		if err != nil {
			return guideIds, err
		}

		guideIds = append(guideIds, guideId)
	}

	return guideIds, nil
}

func (r packageSessionRepositoryPostgresql) DetachGuides(ctx context.Context, id int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "package_session_guides"`).
		S(`WHERE "package_session_id" = $1`, id).
		S(`RETURNING "guide_id"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	guideIds := []int64{}
	for rows.Next() {
		var guideId int64
		if err = rows.Scan(&guideId); err != nil {
			return nil, err
		}

		guideIds = append(guideIds, guideId)
	}

	return guideIds, nil
}
