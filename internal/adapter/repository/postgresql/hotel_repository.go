package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type hotelRepositoryPostgresql struct {
	db DB
}

func NewHotelRepository(db DB) repository.HotelRepository {
	return hotelRepositoryPostgresql{db}
}

func (r hotelRepositoryPostgresql) Create(ctx context.Context, hotel entity.Hotel) (entity.Hotel, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "hotels" ("name", "rating", "map", "address", "distance", "distance_landmark", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW(), NULL)`, hotel.Name, hotel.Rating, hotel.Map, hotel.Address, hotel.Distance, hotel.DistanceLandmark, hotel.Review, hotel.Description, hotel.Location, hotel.Slug).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "distance_landmark", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.DistanceLandmark, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Hotel, error) {
	hotel := entity.Hotel{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "rating", "map", "address", "distance", "distance_landmark", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "hotels" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.DistanceLandmark, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Hotel, error) {
	hotels := []entity.Hotel{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "rating", "map", "address", "distance", "distance_landmark", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "hotels" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return hotels, err
	}

	for rows.Next() {
		hotel := entity.Hotel{}
		err = rows.Scan(
			&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.DistanceLandmark, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
			&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
		)
		if err != nil {
			return hotels, err
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func (r hotelRepositoryPostgresql) Update(ctx context.Context, id int64, hotel entity.Hotel) (entity.Hotel, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "hotels" SET "name" = $1, "rating" = $2, "map" = $3, "address" = $4, "distance" = $5, "distance_landmark" = $6, "review" = $7, "description" = $8, "location" = $9, "slug" = $10, "updated_at" = NOW()`,
			hotel.Name, hotel.Rating, hotel.Map, hotel.Address, hotel.Distance, hotel.DistanceLandmark, hotel.Review, hotel.Description, hotel.Location, hotel.Slug,
		).
		S(`WHERE "id" = $11 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "distance_landmark", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.DistanceLandmark, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Hotel, error) {
	hotel := entity.Hotel{}

	builder := sqlbuilder.New().
		S(`UPDATE "hotels" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "distance_landmark", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.DistanceLandmark, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) AttachImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error) {
	if len(imageIds) == 0 {
		return []int64{}, nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "hotel_images" ("hotel_id", "image_id", "created_at", "updated_at", "deleted_at") VALUES`)

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

func (r hotelRepositoryPostgresql) FindImages(ctx context.Context, id int64) ([]entity.Image, error) {
	images := []entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id", "images"."src", "images"."alt", "images"."category", "images"."title", "images"."created_at", "images"."updated_at", "images"."deleted_at"`).
		S(`FROM "hotel_images"`).
		S(`INNER JOIN "hotels" ON "hotels"."id" = "hotel_images"."hotel_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "hotel_images"."image_id"`).
		S(`WHERE "hotel_images"."hotel_id" = $1 AND "hotel_images"."deleted_at" IS NULL AND "hotels"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

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

func (r hotelRepositoryPostgresql) FindImageIds(ctx context.Context, id int64) ([]int64, error) {
	imageIds := []int64{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id"`).
		S(`FROM "hotel_images"`).
		S(`INNER JOIN "hotels" ON "hotels"."id" = "hotel_images"."hotel_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "hotel_images"."image_id"`).
		S(`WHERE "hotel_images"."hotel_id" = $1 AND "hotel_images"."deleted_at" IS NULL AND "hotels"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`, id)

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

func (r hotelRepositoryPostgresql) DetachImages(ctx context.Context, id int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "hotel_images"`).
		S(`WHERE "hotel_id" = $1`, id).
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
