package airline

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
)

type dao struct {
}

var Dao = dao{}

func (r dao) Insert(tx pgx.Tx, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := queryBuilder.
		InsertInto("airlines").
		Cols("name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at").
		Values(request.Name, request.SkytraxType, request.SkytraxRating, request.Logo, "NOW()", "NOW()").
		Returning("id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Name, &response.SkytraxType, &response.SkytraxRating, &response.LogoId, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) SelectAll(tx pgx.Tx, paginationQuery api.PaginationQuery) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at").
		From("airlines").
		OrderBy("id ASC").
		Limit(paginationQuery.PerPage).
		Offset(paginationQuery.PerPage * (paginationQuery.Page - 1)).
		Where(queryBuilder.IsNull("deleted_at")).
		Build()

	rows, err := database.Pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	responses, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (response Response, err error) {
		err = row.Scan(
			&response.Id, &response.Name, &response.SkytraxType, &response.SkytraxRating, &response.LogoId, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
		)
		return
	})
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (r dao) CountAll(tx pgx.Tx) (int, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("COUNT(*)").
		From("airlines").
		Where(queryBuilder.IsNull("deleted_at")).
		Build()

	var count int
	if err := tx.QueryRow(context.Background(), query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r dao) SelectById(tx pgx.Tx, id int64) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "name", "skytrax_type", "skytrax_rating", "logo_id", "created_at", "updated_at", "deleted_at").
		From("airlines").
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Name, &response.SkytraxType, &response.SkytraxRating, &response.LogoId, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Update(tx pgx.Tx, id int64, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("airlines").
		Set(
			queryBuilder.Assign("name", request.Name),
			queryBuilder.Assign("skytrax_type", request.SkytraxType),
			queryBuilder.Assign("skytrax_rating", request.SkytraxRating),
			queryBuilder.Assign("logo_id", request.Logo),
			queryBuilder.Assign("updated_at", "NOW()"),
		).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, name, skytrax_type, skytrax_rating, logo_id, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Name, &response.SkytraxType, &response.SkytraxRating, &response.LogoId, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Delete(tx pgx.Tx, id int64) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("airlines").
		Set(queryBuilder.Assign("deleted_at", "NOW()")).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, name, skytrax_type, skytrax_rating, logo_id, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Name, &response.SkytraxType, &response.SkytraxRating, &response.LogoId, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}
