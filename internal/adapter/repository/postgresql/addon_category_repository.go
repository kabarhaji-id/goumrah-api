package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type addonCategoryRepositoryPostgresql struct {
	db DB
}

func NewAddonCategoryRepository(db DB) repository.AddonCategoryRepository {
	return addonCategoryRepositoryPostgresql{db}
}

func (r addonCategoryRepositoryPostgresql) Create(ctx context.Context, addonCategory entity.AddonCategory) (entity.AddonCategory, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "addon_categories" ("name", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, NOW(), NOW(), NULL)`, addonCategory.Name).
		S(`RETURNING "id", "name", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addonCategory.Id, &addonCategory.Name,
		&addonCategory.CreatedAt, &addonCategory.UpdatedAt, &addonCategory.DeletedAt,
	)

	return addonCategory, err
}

func (r addonCategoryRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.AddonCategory, error) {
	addonCategory := entity.AddonCategory{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "addon_categories" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addonCategory.Id, &addonCategory.Name,
		&addonCategory.CreatedAt, &addonCategory.UpdatedAt, &addonCategory.DeletedAt,
	)

	return addonCategory, err
}

func (r addonCategoryRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.AddonCategory, error) {
	addonCategories := []entity.AddonCategory{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "name", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "addon_categories" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return addonCategories, err
	}

	for rows.Next() {
		addonCategory := entity.AddonCategory{}
		err = rows.Scan(
			&addonCategory.Id, &addonCategory.Name,
			&addonCategory.CreatedAt, &addonCategory.UpdatedAt, &addonCategory.DeletedAt,
		)
		if err != nil {
			return addonCategories, err
		}

		addonCategories = append(addonCategories, addonCategory)
	}

	return addonCategories, nil
}

func (r addonCategoryRepositoryPostgresql) Update(ctx context.Context, id int64, addonCategory entity.AddonCategory) (entity.AddonCategory, error) {
	builder := sqlbuilder.New().
		S(`UPDATE "addon_categories" SET "name" = $1, "updated_at" = NOW()`, addonCategory.Name).
		S(`WHERE "id" = $2 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addonCategory.Id, &addonCategory.Name,
		&addonCategory.CreatedAt, &addonCategory.UpdatedAt, &addonCategory.DeletedAt,
	)

	return addonCategory, err
}

func (r addonCategoryRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.AddonCategory, error) {
	addonCategory := entity.AddonCategory{}

	builder := sqlbuilder.New().
		S(`UPDATE "addon_categories" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "name", "created_at", "updated_at", "deleted_at`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addonCategory.Id, &addonCategory.Name,
		&addonCategory.CreatedAt, &addonCategory.UpdatedAt, &addonCategory.DeletedAt,
	)

	return addonCategory, err
}
