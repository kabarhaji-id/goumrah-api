package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type packageRepositoryPostgresql struct {
	db DB
}

func NewPackageRepository(db DB) repository.PackageRepository {
	return packageRepositoryPostgresql{db}
}

func (r packageRepositoryPostgresql) Create(ctx context.Context, pkg entity.Package) (entity.Package, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "packages" ("thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW(), NULL)`, pkg.ThumbnailId, pkg.Name, pkg.Description, pkg.IsActive, pkg.Category, pkg.Type, pkg.Slug, pkg.IsRecommended).
		S(`RETURNING "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&pkg.Id, &pkg.ThumbnailId, &pkg.Name, &pkg.Description, &pkg.IsActive, &pkg.Category, &pkg.Type, &pkg.Slug, &pkg.IsRecommended,
		&pkg.CreatedAt, &pkg.UpdatedAt, &pkg.DeletedAt,
	)

	return pkg, err
}

func (r packageRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Package, error) {
	pkg := entity.Package{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "packages" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&pkg.Id, &pkg.ThumbnailId, &pkg.Name, &pkg.Description, &pkg.IsActive, &pkg.Category, &pkg.Type, &pkg.Slug, &pkg.IsRecommended,
		&pkg.CreatedAt, &pkg.UpdatedAt, &pkg.DeletedAt,
	)

	return pkg, err
}

func (r packageRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Package, error) {
	packages := []entity.Package{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "packages" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return packages, err
	}

	for rows.Next() {
		pkg := entity.Package{}
		err = rows.Scan(
			&pkg.Id, &pkg.ThumbnailId, &pkg.Name, &pkg.Description, &pkg.IsActive, &pkg.Category, &pkg.Type, &pkg.Slug, &pkg.IsRecommended,
			&pkg.CreatedAt, &pkg.UpdatedAt, &pkg.DeletedAt,
		)
		if err != nil {
			return packages, err
		}

		packages = append(packages, pkg)
	}

	return packages, nil
}

func (r packageRepositoryPostgresql) Update(ctx context.Context, id int64, pkg entity.Package) (entity.Package, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "packages" SET "thumbnail_id" = $1, "name" = $2, "description" = $3, "is_active" = $4, "category" = $5, "type" = $6, "slug" = $7, "is_recommended" = $8, "updated_at" = NOW()`,
			pkg.ThumbnailId, pkg.Name, pkg.Description, pkg.IsActive, pkg.Category, pkg.Type, pkg.Slug, pkg.IsRecommended,
		).
		S(`WHERE "id" = $9 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&pkg.Id, &pkg.ThumbnailId, &pkg.Name, &pkg.Description, &pkg.IsActive, &pkg.Category, &pkg.Type, &pkg.Slug, &pkg.IsRecommended,
		&pkg.CreatedAt, &pkg.UpdatedAt, &pkg.DeletedAt,
	)

	return pkg, err
}

func (r packageRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Package, error) {
	pkg := entity.Package{}

	builder := sqlbuilder.New().
		S(`UPDATE "packages" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "thumbnail_id", "name", "description", "is_active", "category", "type", "slug", "is_recommended", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&pkg.Id, &pkg.ThumbnailId, &pkg.Name, &pkg.Description, &pkg.IsActive, &pkg.Category, &pkg.Type, &pkg.Slug, &pkg.IsRecommended,
		&pkg.CreatedAt, &pkg.UpdatedAt, &pkg.DeletedAt,
	)

	return pkg, err
}

func (r packageRepositoryPostgresql) CreateImages(ctx context.Context, id int64, imageIds []int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "package_images" ("package_id", "image_id", "created_at", "updated_at", "deleted_at") VALUES`)

	imageIdsLen := len(imageIds)
	for index, imageId := range imageIds {
		builder.S(`(?, ?, NOW(), NOW(), NULL)`, id, imageId)
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

func (r packageRepositoryPostgresql) FindImages(ctx context.Context, id int64) ([]entity.Image, error) {
	images := []entity.Image{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id", "images"."src", "images"."alt", "images"."category", "images"."title", "images"."created_at", "images"."updated_at", "images"."deleted_at"`).
		S(`FROM "package_images"`).
		S(`INNER JOIN "packages" ON "packages"."id" = "package_images"."package_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "package_images"."image_id"`).
		S(`WHERE "package_images"."package_id" = $1 AND "package_images"."deleted_at" IS NULL AND "packages"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`)

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

func (r packageRepositoryPostgresql) FindImageIds(ctx context.Context, id int64) ([]int64, error) {
	imageIds := []int64{}

	builder := sqlbuilder.New().
		S(`SELECT "images"."id"`).
		S(`FROM "package_images"`).
		S(`INNER JOIN "packages" ON "packages"."id" = "package_images"."package_id"`).
		S(`INNER JOIN "images" ON "images"."id" = "package_images"."image_id"`).
		S(`WHERE "package_images"."package_id" = $1 AND "package_images"."deleted_at" IS NULL AND "packages"."deleted_at" IS NULL AND "images"."deleted_at" IS NULL`)

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

func (r packageRepositoryPostgresql) DeleteImages(ctx context.Context, id int64) ([]int64, error) {
	builder := sqlbuilder.New().
		S(`DELETE FROM "package_images"`).
		S(`WHERE "package_id" = $1`, id).
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
