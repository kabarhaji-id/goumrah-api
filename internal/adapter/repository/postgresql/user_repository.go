package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/pkg/sqlbuilder"
)

type userRepositoryPostgresql struct {
	db DB
}

func NewUserRepository(db DB) repository.UserRepository {
	return userRepositoryPostgresql{db}
}

func (r userRepositoryPostgresql) Create(ctx context.Context, user entity.User) (entity.User, error) {
	builder := sqlbuilder.New().
		S(`INSERT INTO "users" ("full_name", "phone_number", "email", "address", "created_at", "updated_at", "deleted_at")`).
		S(`VALUES ($1, $2, $3, $4, NOW(), NOW(), NULL)`, user.FullName, user.PhoneNumber, user.Email, user.Address).
		S(`RETURNING "id", "full_name", "phone_number", "email", "address", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&user.Id, &user.FullName, &user.PhoneNumber, &user.Email, &user.Address,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)

	return user, err
}

func (r userRepositoryPostgresql) FindById(ctx context.Context, id int64) (entity.User, error) {
	user := entity.User{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "full_name", "phone_number", "email", "address", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "users" WHERE "id" = $1 AND "deleted_at" IS NULL`, id)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&user.Id, &user.FullName, &user.PhoneNumber, &user.Email, &user.Address,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)

	return user, err
}

func (r userRepositoryPostgresql) FindAll(ctx context.Context, opt repository.FindAllOptions) ([]entity.User, error) {
	users := []entity.User{}

	builder := sqlbuilder.New().
		S(`SELECT "id", "full_name", "phone_number", "email", "address", "created_at", "updated_at", "deleted_at"`).
		S(`FROM "users" WHERE "deleted_at" IS NULL`).
		S(`ORDER BY "id" ASC`)
	if opt.Limit.Valid {
		builder.SA(`LIMIT ?`, opt.Limit)
	}
	if opt.Offset.Valid {
		builder.SA(`OFFSET ?`, opt.Offset)
	}

	rows, err := r.db.Query(ctx, builder.Query(), builder.Args()...)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(
			&user.Id, &user.FullName, &user.PhoneNumber, &user.Email, &user.Address,
			&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r userRepositoryPostgresql) Update(ctx context.Context, id int64, user entity.User) (entity.User, error) {
	builder := sqlbuilder.New().
		S(
			`UPDATE "users" SET "full_name" = $1, "phone_number" = $2, "email" = $3, "address" = $4, "updated_at" = NOW()`,
			user.FullName, user.PhoneNumber, user.Email, user.Address,
		).
		S(`WHERE "id" = $5 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "full_name", "phone_number", "email", "address", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&user.Id, &user.FullName, &user.PhoneNumber, &user.Email, &user.Address,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)

	return user, err
}

func (r userRepositoryPostgresql) Delete(ctx context.Context, id int64) (entity.User, error) {
	user := entity.User{}

	builder := sqlbuilder.New().
		S(`UPDATE "users" SET "deleted_at" = NOW()`).
		S(`WHERE "id" = $1 AND "deleted_at" IS NULL`, id).
		S(`RETURNING "id", "full_name", "phone_number", "email", "address", "created_at", "updated_at", "deleted_at"`)

	err := r.db.QueryRow(ctx, builder.Query(), builder.Args()...).Scan(
		&user.Id, &user.FullName, &user.PhoneNumber, &user.Email, &user.Address,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)

	return user, err
}
