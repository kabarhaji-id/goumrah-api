package airport

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
		InsertInto("airports").
		Cols("city", "name", "code", "created_at", "updated_at").
		Values(request.City, request.Name, request.Code, "NOW()", "NOW()").
		Returning("id", "city", "name", "code", "created_at", "updated_at", "deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.City, &response.Name, &response.Code, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) SelectAll(tx pgx.Tx, paginationQuery api.PaginationQuery) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "city", "name", "code", "created_at", "updated_at", "deleted_at").
		From("airports").
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
			&response.Id, &response.City, &response.Name, &response.Code, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
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
		From("airports").
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
		Select("id", "city", "name", "code", "created_at", "updated_at", "deleted_at").
		From("airports").
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.City, &response.Name, &response.Code, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Update(tx pgx.Tx, id int64, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("airports").
		Set(
			queryBuilder.Assign("city", request.City),
			queryBuilder.Assign("name", request.Name),
			queryBuilder.Assign("code", request.Code),
			queryBuilder.Assign("updated_at", "NOW()"),
		).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, city, name, code, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.City, &response.Name, &response.Code, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Delete(tx pgx.Tx, id int64) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("airports").
		Set(queryBuilder.Assign("deleted_at", "NOW()")).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, city, name, code, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.City, &response.Name, &response.Code, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}
