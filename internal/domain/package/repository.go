package pkg

import (
	"context"
	"fmt"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
	"github.com/kabarhaji-id/goumrah-api/pkg/sluger"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type RepositoryFindAllOption struct {
	Limit  null.Int
	Offset null.Int
}

type Repository struct {
	db database.DB
}

func NewRepository(db database.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(ctx context.Context, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`INSERT INTO "packages" ("thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())`, entity.ThumbnailId, entity.Name, entity.Description, entity.IsActive, entity.Category, entity.Type, sluger.Slug(entity.Name), entity.IsRecommended).
		S(`RETURNING "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.ThumbnailId, &entity.Name, &entity.Description, &entity.IsActive, &entity.Category, &entity.Type, &entity.Slug, &entity.IsRecommended, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) FindAll(ctx context.Context, opt RepositoryFindAllOption) ([]Entity, error) {
	builder := sqlbuilder.New().
		S(`SELECT "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at"`).
		S(`FROM "packages"`).
		S(`WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit.Int64)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset.Int64)
	}

	query, args := builder.Build()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var entities []Entity
	for rows.Next() {
		entity := Entity{}
		if err := rows.Scan(
			&entity.Id, &entity.ThumbnailId, &entity.Name, &entity.Description, &entity.IsActive, &entity.Category, &entity.Type, &entity.Slug, &entity.IsRecommended, &entity.CreatedAt, &entity.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (r Repository) FindById(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at"`).
		S(`FROM "packages"`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.ThumbnailId, &entity.Name, &entity.Description, &entity.IsActive, &entity.Category, &entity.Type, &entity.Slug, &entity.IsRecommended, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Count(ctx context.Context) (int, error) {
	query, args := sqlbuilder.New().
		S(`SELECT COUNT(*)`).
		S(`FROM "packages"`).
		S(`WHERE "deleted_at" IS NULL`).
		Build()

	var count int
	if err := r.db.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r Repository) Update(ctx context.Context, id int64, entity Entity) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "packages"`).
		S(`SET "thumbnail_id" = $1, "name" = $2, "description" = $3, "is_active" = $4, "category" = $5, "type" = $6, "slug" = $7, "is_recommended" = $8, "updated_at" = NOW()`, entity.ThumbnailId, entity.Name, entity.Description, entity.IsActive, entity.Category, entity.Type, sluger.Slug(entity.Slug), entity.IsRecommended).
		S(`WHERE "id" = $9 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at"`).
		Build()

	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.ThumbnailId, &entity.Name, &entity.Description, &entity.IsActive, &entity.Category, &entity.Type, &entity.Slug, &entity.IsRecommended, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id int64) (Entity, error) {
	query, args := sqlbuilder.New().
		S(`UPDATE "packages"`).
		S(`SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at"`).
		Build()

	entity := Entity{}
	if err := r.db.QueryRow(ctx, query, args...).Scan(
		&entity.Id, &entity.ThumbnailId, &entity.Name, &entity.Description, &entity.IsActive, &entity.Category, &entity.Type, &entity.Slug, &entity.IsRecommended, &entity.CreatedAt, &entity.UpdatedAt,
	); err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r Repository) CreateImages(ctx context.Context, packageId int64, imageIds []int64) error {
	if len(imageIds) == 0 {
		return nil
	}

	builder := sqlbuilder.New().
		S(`INSERT INTO "package_images" ("package_id", "image_id", "created_at", "updated_at") VALUES`)

	imageIdsLength := len(imageIds)
	for id, imageId := range imageIds {
		builder.SA(`(?, ?, NOW(), NOW())`, packageId, imageId)

		if id < imageIdsLength-1 {
			builder.S(`, `)
		}
	}

	query, args := builder.Build()

	fmt.Println(query)

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r Repository) FindImageIds(ctx context.Context, packageId int64) ([]int64, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "image_id"`).
		S(`FROM "package_images"`).
		S(`WHERE "package_id" = $1`, packageId).
		Build()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	imageIds := make([]int64, 0)
	for rows.Next() {
		var imageId int64
		if err := rows.Scan(&imageId); err != nil {
			return nil, err
		}

		imageIds = append(imageIds, imageId)
	}

	return imageIds, nil
}

func (r Repository) FindImages(ctx context.Context, packageId int64) ([]image.Entity, error) {
	query, args := sqlbuilder.New().
		S(`SELECT "images"."id", "images"."src", "images"."alt", "images"."category", "images"."title", "images"."created_at", "images"."updated_at", "images"."deleted_at"`).
		S(`FROM "package_images"`).
		S(`JOIN "images" ON "package_images"."image_id" = "images"."id"`).
		S(`WHERE "package_images"."package_id" = $1`, packageId).
		Build()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var imageEntities []image.Entity
	for rows.Next() {
		imageEntity := image.Entity{}
		if err := rows.Scan(
			&imageEntity.Id, &imageEntity.Src, &imageEntity.Alt, &imageEntity.Category, &imageEntity.Title, &imageEntity.CreatedAt, &imageEntity.UpdatedAt, &imageEntity.DeletedAt,
		); err != nil {
			return nil, err
		}

		imageEntities = append(imageEntities, imageEntity)
	}

	return imageEntities, nil
}

func (r Repository) DeleteImages(ctx context.Context, packageId int64) error {
	query, args := sqlbuilder.New().
		S(`DELETE FROM "package_images"`).
		S(`WHERE "package_id" = $1`, packageId).
		Build()

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
