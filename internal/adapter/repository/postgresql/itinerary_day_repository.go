package postgresql

import (
	"context"
	"log"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type itineraryDayRepositoryPostgresql struct {
	db DB
}

func NewItineraryDayRepository(db DB) repository.ItineraryDayRepository {
	return itineraryDayRepositoryPostgresql{db}
}

func (r itineraryDayRepositoryPostgresql) Create(ctx context.Context, itineraryDay entity.ItineraryDay) (entity.ItineraryDay, error) {
	log.Println(itineraryDay)
	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_days" ("title", "description", "widget_id", "next_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, itineraryDay.Title, itineraryDay.Description, itineraryDay.WidgetId, itineraryDay.NextId).
		S(`RETURNING "id", "title", "description", "widget_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryDay.Id, &itineraryDay.Title, &itineraryDay.Description, &itineraryDay.WidgetId, &itineraryDay.NextId,
		&itineraryDay.CreatedAt, &itineraryDay.UpdatedAt, &itineraryDay.DeletedAt,
	)

	return itineraryDay, err
}

func (r itineraryDayRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.ItineraryDay, error) {
	itineraryDay := entity.ItineraryDay{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "title", "description", "widget_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_days" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryDay.Id, &itineraryDay.Title, &itineraryDay.Description, &itineraryDay.WidgetId, &itineraryDay.NextId,
		&itineraryDay.CreatedAt, &itineraryDay.UpdatedAt, &itineraryDay.DeletedAt,
	)

	return itineraryDay, err
}

func (r itineraryDayRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.ItineraryDay, error) {
	itineraries := []entity.ItineraryDay{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "title", "description", "widget_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itinerary_days" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return itineraries, err
	}

	for rows.Next() {
		itineraryDay := entity.ItineraryDay{}
		err = rows.Scan(
			&itineraryDay.Id, &itineraryDay.Title, &itineraryDay.Description, &itineraryDay.WidgetId, &itineraryDay.NextId,
			&itineraryDay.CreatedAt, &itineraryDay.UpdatedAt, &itineraryDay.DeletedAt,
		)
		if err != nil {
			return itineraries, err
		}

		itineraries = append(itineraries, itineraryDay)
	}

	return itineraries, nil
}

func (r itineraryDayRepositoryPostgresql) Update(ctx context.Context, id int64, itineraryDay entity.ItineraryDay) (entity.ItineraryDay, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itinerary_days" SET "title" = $1, "description" = $2, "widget_id" = $3, "next_id" = $4, "updated_at" = NOW()`,
			itineraryDay.Title, itineraryDay.Description, itineraryDay.WidgetId, itineraryDay.NextId,
		).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "widget_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryDay.Id, &itineraryDay.Title, &itineraryDay.Description, &itineraryDay.WidgetId, &itineraryDay.NextId,
		&itineraryDay.CreatedAt, &itineraryDay.UpdatedAt, &itineraryDay.DeletedAt,
	)

	return itineraryDay, err
}

func (r itineraryDayRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.ItineraryDay, error) {
	itineraryDay := entity.ItineraryDay{}

	builder := sqlbuilder.New().
		S(`UPDATE "itinerary_days" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "title", "description", "widget_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itineraryDay.Id, &itineraryDay.Title, &itineraryDay.Description, &itineraryDay.WidgetId, &itineraryDay.NextId,
		&itineraryDay.CreatedAt, &itineraryDay.UpdatedAt, &itineraryDay.DeletedAt,
	)

	return itineraryDay, err
}
