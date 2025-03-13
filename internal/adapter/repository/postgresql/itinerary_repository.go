package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type itineraryRepositoryPostgresql struct {
	db DB
}

func NewItineraryRepository(db DB) repository.ItineraryRepository {
	return itineraryRepositoryPostgresql{db}
}

func (r itineraryRepositoryPostgresql) Create(ctx context.Context, itinerary entity.Itinerary) (entity.Itinerary, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "itineraries" ("city", "day_id", "next_id", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, itinerary.City, itinerary.DayId, itinerary.NextId).
		S(`RETURNING "id", "city", "day_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itinerary.Id, &itinerary.City, &itinerary.DayId, &itinerary.NextId,
		&itinerary.CreatedAt, &itinerary.UpdatedAt, &itinerary.DeletedAt,
	)

	return itinerary, err
}

func (r itineraryRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Itinerary, error) {
	itinerary := entity.Itinerary{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "city", "day_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itineraries" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itinerary.Id, &itinerary.City, &itinerary.DayId, &itinerary.NextId,
		&itinerary.CreatedAt, &itinerary.UpdatedAt, &itinerary.DeletedAt,
	)

	return itinerary, err
}

func (r itineraryRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Itinerary, error) {
	itineraries := []entity.Itinerary{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "city", "day_id", "next_id", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "itineraries" WHERE "deleted_at" IS NULL`).
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
		itinerary := entity.Itinerary{}
		err = rows.Scan(
			&itinerary.Id, &itinerary.City, &itinerary.DayId, &itinerary.NextId,
			&itinerary.CreatedAt, &itinerary.UpdatedAt, &itinerary.DeletedAt,
		)
		if err != nil {
			return itineraries, err
		}

		itineraries = append(itineraries, itinerary)
	}

	return itineraries, nil
}

func (r itineraryRepositoryPostgresql) Update(ctx context.Context, id int64, itinerary entity.Itinerary) (entity.Itinerary, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "itineraries" SET "city" = $1, "day_id" = $2, "next_id" = $3, "updated_at" = NOW()`,
			itinerary.City, itinerary.DayId, itinerary.NextId,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "city", "day_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itinerary.Id, &itinerary.City, &itinerary.DayId, &itinerary.NextId,
		&itinerary.CreatedAt, &itinerary.UpdatedAt, &itinerary.DeletedAt,
	)

	return itinerary, err
}

func (r itineraryRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Itinerary, error) {
	itinerary := entity.Itinerary{}

	builder := sqlbuilder.New().
		S(`UPDATE "itineraries" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "city", "day_id", "next_id", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&itinerary.Id, &itinerary.City, &itinerary.DayId, &itinerary.NextId,
		&itinerary.CreatedAt, &itinerary.UpdatedAt, &itinerary.DeletedAt,
	)

	return itinerary, err
}

func (r itineraryRepositoryPostgresql) AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error) {
	if len(imageIds) == 0 {
		return []int64{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "itinerary_images" ("itinerary_id", "image_id", "created_at", "updated_at", "deleted_at") VALUES`)

	imageIdsLen := len(imageIds)
	for index, imageId := range imageIds {
		builder.SA(`(?, ?, NOW(), NOW(), NULL)`, id, imageId)
		if index+1 < imageIdsLen {
			builder.S(",")
		}
	}

	_, err := r.db.Exec(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}

	return imageIds, nil
}

func (r itineraryRepositoryPostgresql) FindImages(ctx context.Context, id int64) ([]entity.Image, error) {
	images := []entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id", "images"."src", "images"."alt", "images"."category", "images"."title", "images"."created_at", "images"."updated_at", "images"."deleted_at"`).
		S(`FROM "itinerary_images"`).
		S(`INNER JOIN "itineraries" ON "itineraries"."id" = "itinerary_images"."itinerary_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "itinerary_images"."image_id"`).
		S(`WHERE "itinerary_images"."itinerary_id" = $1 AND "itinerary_images"."deleted_at" IS NULL AND "itineraries"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return images, err
	}

	for rows.Next() {
		image := entity.Image{}
		err = rows.Scan(
			&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
			&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
		)
		if err != nil {
			return images, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (r itineraryRepositoryPostgresql) FindImageIds(ctx context.Context, id int64) ([]int64, error) {
	imageIds := []int64{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id"`).
		S(`FROM "itinerary_images"`).
		S(`INNER JOIN "itineraries" ON "itineraries"."id" = "itinerary_images"."itinerary_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "itinerary_images"."image_id"`).
		S(`WHERE "itinerary_images"."itinerary_id" = $1 AND "itinerary_images"."deleted_at" IS NULL AND "itineraries"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return imageIds, err
	}

	for rows.Next() {
		var imageId int64
		err = rows.Scan(&imageId)
		if err != nil {
			return imageIds, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}

func (r itineraryRepositoryPostgresql) DetachImages(ctx context.Context, id int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "itinerary_images"`).
		S(`WHERE "itinerary_id" = $1`, id).
		S(`RETURNING "image_id"`)

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	imageIds := []int64{}
	for rows.Next() {
		var imageId int64
		if err = rows.Scan(&imageId); err != nil {
			return nil, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}
