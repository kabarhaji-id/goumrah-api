package image

import (
	"context"
	"crypto/rand"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
)

func CreateHandler(c *fiber.Ctx) error {
	// Validate and get request
	request, success, err := validateRequest(c)
	if !success {
		return err
	}

	// Generate image file name
	imageFileName := strings.ToLower(rand.Text()) + filepath.Ext(request.Image.Filename)

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Build query
	queryBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := queryBuilder.
		InsertInto("images").
		Cols("src", "alt", "category", "title", "created_at", "updated_at").
		Values(imageFileName, request.Alt, request.Category, request.Title, "NOW()", "NOW()").
		Returning("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		Build()

	// Insert image into database and scan returning into response
	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		tx.Rollback(context.Background())

		return api.ErrInternalServer(c, err)
	}

	// Save image into public folder
	if err := c.SaveFile(request.Image, filepath.Join("public", imageFileName)); err != nil {
		tx.Rollback(context.Background())

		return api.ErrInternalServer(c, err)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return api.ErrInternalServer(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(api.ResponseData(response))
}

func GetAllHandler(c *fiber.Ctx) error {
	// Validate and get request query for pagination
	paginationQuery, success, err := api.ValidatePaginationQuery(c)
	if !success {
		return err
	}

	// Build query for select all images
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		From("images").
		OrderBy("id ASC").
		Limit(paginationQuery.PerPage).
		Offset(paginationQuery.PerPage * (paginationQuery.Page - 1)).
		Where(queryBuilder.IsNull("deleted_at")).
		Build()

	// Query images from database
	rows, err := database.Pool.Query(context.Background(), query, args...)
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Collect rows into response
	responses, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (response Response, err error) {
		err = row.Scan(
			&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
		)
		return
	})
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Build query for count all images
	queryBuilder = sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args = queryBuilder.
		Select("COUNT(*)").
		From("images").
		Where(queryBuilder.IsNull("deleted_at")).
		Build()

	// Query count all images from database
	var count int
	if err := database.Pool.QueryRow(context.Background(), query, args...).Scan(&count); err != nil {
		return api.ErrInternalServer(c, err)
	}

	return c.JSON(api.ResponseData(responses, api.PaginationMeta{
		Page:      paginationQuery.Page,
		PerPage:   paginationQuery.PerPage,
		FirstPage: 1,
		LastPage:  (count + paginationQuery.PerPage - 1) / paginationQuery.PerPage,
		Total:     count,
	}))
}

func GetOneHandler(c *fiber.Ctx) error {
	// Validate and get id param
	id, success, err := api.ValidateId(c)
	if !success {
		return err
	}

	// Build query
	queryBuilder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := queryBuilder.
		Select("id", "src", "alt", "category", "title", "created_at", "updated_at", "deleted_at").
		From("images").
		Where(
			queryBuilder.Equal("id", id),
			queryBuilder.IsNull("deleted_at"),
		).
		Build()

	// Query image from database and scan into response
	response := Response{}
	if err := database.Pool.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound(c, err)
		}

		return api.ErrInternalServer(c, err)
	}

	return c.JSON(api.ResponseData(response))
}

func UpdateHandler(c *fiber.Ctx) error {
	// Validate and get id param
	id, success, err := api.ValidateId(c)
	if !success {
		return err
	}

	// Validate and get request
	request, success, err := validateRequest(c)
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Build query
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

	// Update image in database and scan returning into response
	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		tx.Rollback(context.Background())

		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound(c, err)
		}

		return api.ErrInternalServer(c, err)
	}

	// Save image into public folder
	if err := c.SaveFile(request.Image, filepath.Join("public", response.Src)); err != nil {
		tx.Rollback(context.Background())

		return api.ErrInternalServer(c, err)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return api.ErrInternalServer(c, err)
	}

	return c.JSON(api.ResponseData(response))
}

func DeleteHandler(c *fiber.Ctx) error {
	// Validate and get id param
	id, success, err := api.ValidateId(c)
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Build query
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

	// Delete image in database and scan returning into response
	response := Response{}
	if err := tx.QueryRow(context.Background(), query, args...).Scan(
		&response.Id, &response.Src, &response.Alt, &response.Category, &response.Title, &response.CreatedAt, &response.UpdatedAt, &response.DeletedAt,
	); err != nil {
		tx.Rollback(context.Background())

		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound(c, err)
		}

		return api.ErrInternalServer(c, err)
	}

	// Remove image from public folder
	if err := os.Remove(filepath.Join("public", response.Src)); err != nil {
		tx.Rollback(context.Background())

		return api.ErrInternalServer(c, err)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return api.ErrInternalServer(c, err)
	}

	return c.JSON(api.ResponseData(response))
}
