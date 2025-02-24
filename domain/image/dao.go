package image

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

func (r dao) Insert(tx pgx.Tx, src string, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := queryBuilder.
		InsertInto("images").
		Cols("src", "alt", "category", "title", "created_at", "updated_at").
		Values(src, request.Alt, request.Category, request.Title, "NOW()", "NOW()").
		Returning("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) SelectAll(tx pgx.Tx, paginationQuery api.PaginationQuery) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		From("images").
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
			&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
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
		From("images").
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
		Select("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		From("images").
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) SelectByIds(tx pgx.Tx, ids []int64) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		From("images").
		Where(
			queryBuilder.In("id", ids),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	rows, err := database.Pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	responses, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (response Response, err error) {
		err = row.Scan(
			&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
		)
		return
	})
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (r dao) Update(tx pgx.Tx, id int64, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("images").
		Set(
			queryBuilder.Assign("alt", request.Alt),
			queryBuilder.Assign("category", request.Category),
			queryBuilder.Assign("title", request.Title),
			queryBuilder.Assign("updated_at", "NOW()"),
		).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, src, alt, category, title, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Delete(tx pgx.Tx, id int64) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("images").
		Set(queryBuilder.Assign("deleted_at", "NOW()")).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, src, alt, category, title, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}
