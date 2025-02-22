package pkg

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
	"github.com/kabarhaji-id/goumrah-api/util"
)

type dao struct {
}

var Dao = dao{}

func (r dao) Insert(tx pgx.Tx, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := queryBuilder.
		InsertInto("packages").
		Cols("thumbnail_id", "name", "description", "is_active", "category", "type", "is_recommended", "slug", "created_at", "updated_at").
		Values(request.Thumbnail, request.Name, request.Description, request.IsActive, request.Category, request.Type, request.IsRecommended, util.Slug(request.Name), "NOW()", "NOW()").
		Returning("id", "thumbnail_id", "name", "description", "is_active", "category", "type", "is_recommended", "slug", "created_at", "updated_at", "deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.ThumbnailId, &response.Name, &response.Description, &response.IsActive, &response.Category, &response.Type, &response.IsRecommended, &response.Slug, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) SelectAll(tx pgx.Tx, paginationQuery api.PaginationQuery) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "thumbnail_id", "name", "description", "is_active", "category", "type", "is_recommended", "slug", "created_at", "updated_at", "deleted_at").
		From("packages").
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
			&response.Id, &response.ThumbnailId, &response.Name, &response.Description, &response.IsActive, &response.Category, &response.Type, &response.IsRecommended, &response.Slug, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
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
		From("packages").
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
		Select("id", "thumbnail_id", "name", "description", "is_active", "category", "type", "is_recommended", "slug", "created_at", "updated_at", "deleted_at").
		From("packages").
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.ThumbnailId, &response.Name, &response.Description, &response.IsActive, &response.Category, &response.Type, &response.IsRecommended, &response.Slug, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Update(tx pgx.Tx, id int64, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("packages").
		Set(
			queryBuilder.Assign("thumbnail_id", request.Thumbnail),
			queryBuilder.Assign("name", request.Name),
			queryBuilder.Assign("description", request.Description),
			queryBuilder.Assign("is_active", request.IsActive),
			queryBuilder.Assign("category", request.Category),
			queryBuilder.Assign("type", request.Type),
			queryBuilder.Assign("is_recommended", request.IsRecommended),
			queryBuilder.Assign("slug", util.Slug(request.Name)),
			queryBuilder.Assign("updated_at", "NOW()"),
		).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, thumbnail_id, name, description, is_active, category, type, is_recommended, slug, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.ThumbnailId, &response.Name, &response.Description, &response.IsActive, &response.Category, &response.Type, &response.IsRecommended, &response.Slug, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Delete(tx pgx.Tx, id int64) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("packages").
		Set(queryBuilder.Assign("deleted_at", "NOW()")).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, thumbnail_id, name, description, is_active, category, type, is_recommended, slug, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.ThumbnailId, &response.Name, &response.Description, &response.IsActive, &response.Category, &response.Type, &response.IsRecommended, &response.Slug, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}
