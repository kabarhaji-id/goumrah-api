package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type addonRepositoryPostgresql struct {
	db DB
}

func NewAddonRepository(db DB) repository.AddonRepository {
	return addonRepositoryPostgresql{db}
}

func (r addonRepositoryPostgresql) Create(ctx context.Context, addon entity.Addon) (entity.Addon, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "addons" ("category_id", "name", "price", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, NOW(), NOW(), NULL)`, addon.CategoryId, addon.Name, addon.Price).
		S(`RETURNING "id", "category_id", "name", "price", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addon.Id, &addon.CategoryId, &addon.Name, &addon.Price,
		&addon.CreatedAt, &addon.UpdatedAt, &addon.DeletedAt,
	)

	return addon, err
}

func (r addonRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Addon, error) {
	addon := entity.Addon{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "category_id", "name", "price", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "addons" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addon.Id, &addon.CategoryId, &addon.Name, &addon.Price,
		&addon.CreatedAt, &addon.UpdatedAt, &addon.DeletedAt,
	)

	return addon, err
}

func (r addonRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Addon, error) {
	addons := []entity.Addon{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "category_id", "name", "price", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "addons" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return addons, err
	}

	for rows.Next() {
		addon := entity.Addon{}
		err = rows.Scan(
			&addon.Id, &addon.CategoryId, &addon.Name, &addon.Price,
			&addon.CreatedAt, &addon.UpdatedAt, &addon.DeletedAt,
		)
		if err != nil {
			return addons, err
		}

		addons = append(addons, addon)
	}

	return addons, nil
}

func (r addonRepositoryPostgresql) Update(ctx context.Context, id int64, addon entity.Addon) (entity.Addon, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "addons" SET "category_id" = $1, "name" = $2, "price" = $3, "updated_at" = NOW()`,
			addon.CategoryId, addon.Name, addon.Price,
		).
		S(`WHERE "id" = $4 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "category_id", "name", "price", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addon.Id, &addon.CategoryId, &addon.Name, &addon.Price,
		&addon.CreatedAt, &addon.UpdatedAt, &addon.DeletedAt,
	)

	return addon, err
}

func (r addonRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Addon, error) {
	addon := entity.Addon{}

	builder := sqlbuilder.New().
		S(`UPDATE "addons" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "category_id", "name", "price", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&addon.Id, &addon.CategoryId, &addon.Name, &addon.Price,
		&addon.CreatedAt, &addon.UpdatedAt, &addon.DeletedAt,
	)

	return addon, err
}
