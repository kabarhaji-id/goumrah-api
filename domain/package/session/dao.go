package pkgsession

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

func (r dao) Insert(tx pgx.Tx, packageId int64, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := queryBuilder.
		InsertInto("package_sessions").
		Cols("package_id", "embarkation_id", "departure_date", "created_at", "updated_at").
		Values(packageId, request.Embarkation, request.DepartureDate, "NOW()", "NOW()").
		Returning("id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.PackageId, &response.EmbarkationId, &response.DepartureDate, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) SelectAll(tx pgx.Tx, paginationQuery api.PaginationQuery) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at").
		From("package_sessions").
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
			&response.Id, &response.PackageId, &response.EmbarkationId, &response.DepartureDate, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
		)
		return
	})
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (r dao) SelectByPackageId(tx pgx.Tx, packageId int64, paginationQuery api.PaginationQuery) ([]Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at").
		From("package_sessions").
		OrderBy("id ASC").
		Limit(paginationQuery.PerPage).
		Offset(paginationQuery.PerPage*(paginationQuery.Page-1)).
		Where(
			queryBuilder.Equal("package_id", packageId),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	rows, err := database.Pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	responses, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (response Response, err error) {
		err = row.Scan(
			&response.Id, &response.PackageId, &response.EmbarkationId, &response.DepartureDate, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
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
		From("package_sessions").
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
		Select("id", "package_id", "embarkation_id", "departure_date", "created_at", "updated_at", "deleted_at").
		From("package_sessions").
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.PackageId, &response.EmbarkationId, &response.DepartureDate, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Update(tx pgx.Tx, id int64, request Request) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("package_sessions").
		Set(
			queryBuilder.Assign("embarkation_id", request.Embarkation),
			queryBuilder.Assign("departure_date", request.DepartureDate),
			queryBuilder.Assign("updated_at", "NOW()"),
		).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, package_id, embarkation_id, departure_date, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.PackageId, &response.EmbarkationId, &response.DepartureDate, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (r dao) Delete(tx pgx.Tx, id int64) (Response, error) {
	queryBuilder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := queryBuilder.
		Update("package_sessions").
		Set(queryBuilder.Assign("deleted_at", "NOW()")).
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		SQL("RETURNING id, package_id, embarkation_id, departure_date, created_at, updated_at, deleted_at").
		Build()

	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.PackageId, &response.EmbarkationId, &response.DepartureDate, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		return Response{}, err
	}

	return response, nil
}
