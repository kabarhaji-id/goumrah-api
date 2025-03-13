package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type guideRepositoryPostgresql struct {
	db DB
}

func NewGuideRepository(db DB) repository.GuideRepository {
	return guideRepositoryPostgresql{db}
}

func (r guideRepositoryPostgresql) Create(ctx context.Context, guide entity.Guide) (entity.Guide, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "guides" ("avatar_id", "name", "type", "description", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, guide.AvatarId, guide.Name, guide.Type, guide.Description).
		S(`RETURNING "id", "avatar_id", "name", "type", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&guide.Id, &guide.AvatarId, &guide.Name, &guide.Type, &guide.Description,
		&guide.CreatedAt, &guide.UpdatedAt, &guide.DeletedAt,
	)

	return guide, err
}

func (r guideRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.Guide, error) {
	guide := entity.Guide{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "avatar_id", "name", "type", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "guides" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&guide.Id, &guide.AvatarId, &guide.Name, &guide.Type, &guide.Description,
		&guide.CreatedAt, &guide.UpdatedAt, &guide.DeletedAt,
	)

	return guide, err
}

func (r guideRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.Guide, error) {
	guides := []entity.Guide{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "avatar_id", "name", "type", "description", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "guides" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

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

func (r guideRepositoryPostgresql) Update(ctx context.Context, id int64, guide entity.Guide) (entity.Guide, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "guides" SET "avatar_id" = $1, "name" = $2, "type" = $3, "description" = $4, "updated_at" = NOW()`,
			guide.AvatarId, guide.Name, guide.Type, guide.Description,
		).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "avatar_id", "name", "type", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&guide.Id, &guide.AvatarId, &guide.Name, &guide.Type, &guide.Description,
		&guide.CreatedAt, &guide.UpdatedAt, &guide.DeletedAt,
	)

	return guide, err
}

func (r guideRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.Guide, error) {
	guide := entity.Guide{}

	builder := sqlbuilder.New().
		S(`UPDATE "guides" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "avatar_id", "name", "type", "description", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&guide.Id, &guide.AvatarId, &guide.Name, &guide.Type, &guide.Description,
		&guide.CreatedAt, &guide.UpdatedAt, &guide.DeletedAt,
	)

	return guide, err
}
