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
		S(`INSERT INTO "hotels" ("name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW(), NULL)`, hotel.Name, hotel.Rating, hotel.Map, hotel.Address, hotel.Distance, hotel.Review, hotel.Description, hotel.Location, hotel.Slug).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Hotel, error) {
	hotel := entity.Hotel{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "hotels" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Hotel, error) {
	hotels := []entity.Hotel{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`).
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
			&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
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
			`UPDATE "hotels" SET "name" = $1, "rating" = $2, "map" = $3, "address" = $4, "distance" = $5, "review" = $6, "description" = $7, "location" = $8, "slug" = $9, "updated_at" = NOW()`,
			hotel.Name, hotel.Rating, hotel.Map, hotel.Address, hotel.Distance, hotel.Review, hotel.Description, hotel.Location, hotel.Slug,
		).
		S(`WHERE "id" = $10 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}

func (r hotelRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Hotel, error) {
	hotel := entity.Hotel{}

	builder := sqlbuilder.New().
		S(`UPDATE "hotels" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "rating", "map", "address", "distance", "review", "description", "location", "slug", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&hotel.Id, &hotel.Name, &hotel.Rating, &hotel.Map, &hotel.Address, &hotel.Distance, &hotel.Review, &hotel.Description, &hotel.Location, &hotel.Slug,
		&hotel.CreatedAt, &hotel.UpdatedAt, &hotel.DeletedAt,
	)

	return hotel, err
}
