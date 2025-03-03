package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type imageRepositoryPostgresql struct {
	db DB
}

func NewImageRepository(db DB) repository.ImageRepository {
	return imageRepositoryPostgresql{db}
}

func (r imageRepositoryPostgresql) Create(ctx context.Context, image entity.Image) (entity.Image, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "images" ("src", "alt", "category", "title", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, image.Src, image.Alt, image.Category, image.Title).
		S(`RETURNING "id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
		&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
	)

	return image, err
}

func (r imageRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Image, error) {
	image := entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "images" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
		&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
	)

	return image, err
}

func (r imageRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Image, error) {
	images := []entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "images" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

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

func (r imageRepositoryPostgresql) Update(ctx context.Context, id int64, image entity.Image) (entity.Image, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "images" SET "alt" = $1, "category" = $2, "title" = $3, "updated_at" = NOW()`,
			image.Alt, image.Category, image.Title,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
		&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
	)

	return image, err
}

func (r imageRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Image, error) {
	image := entity.Image{}

	builder := sqlbuilder.New().
		S(`UPDATE "images" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&image.Id, &image.Src, &image.Alt, &image.Category, &image.Title,
		&image.CreatedAt, &image.UpdatedAt, &image.DeletedAt,
	)

	return image, err
}
