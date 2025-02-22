package pkg

import (
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
	"github.com/kabarhaji-id/goumrah-api/domain/image"
)

func handleError(c *fiber.Ctx, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return api.ErrNotFound(c, err)
	}

	pgError := new(pgconn.PgError)
	if errors.As(err, &pgError) {
		if pgError.Code == "23503" && pgError.ConstraintName == "packages_thumbnail_id_fkey" {
			return api.ErrInvalidRequestField(c, "thumbnail", "Not found")
		}
		if pgError.Code == "23505" && pgError.ConstraintName == "packages_name_unique" {
			return api.ErrConflictField(c, "name")
		}
	}

	return api.ErrInternalServer(c, err)
}

func CreateHandler(c *fiber.Ctx) error {
	// Validate and get request
	request, success, err := validateRequest(c)
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return handleError(c, err)
	}

	// Insert package into database
	response, err := Dao.Insert(tx, request)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Select thumbnail if not null from database
	if response.ThumbnailId.Valid {
		thumbnail, err := image.Dao.SelectById(tx, response.ThumbnailId.Int64)
		if err != nil {
			tx.Rollback(context.Background())

			return handleError(c, err)
		}

		response.Thumbnail = null.ValueFrom(thumbnail)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(api.ResponseData(response))
}

func GetAllHandler(c *fiber.Ctx) error {
	// Validate and get request query for pagination
	paginationQuery, success, err := api.ValidatePaginationQuery(c)
	if !success {
		return handleError(c, err)
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return handleError(c, err)
	}

	// Select all packages from database
	responses, err := Dao.SelectAll(tx, paginationQuery)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Select thumbnail if not null from database
	for i, response := range responses {
		if response.ThumbnailId.Valid {
			log.Println(response.ThumbnailId.Int64)
			thumbnail, err := image.Dao.SelectById(tx, response.ThumbnailId.Int64)
			if err != nil {
				tx.Rollback(context.Background())

				return handleError(c, err)
			}

			responses[i].Thumbnail = null.ValueFrom(thumbnail)
		}
	}

	// Count all packages from database
	count, err := Dao.CountAll(tx)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return handleError(c, err)
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

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return handleError(c, err)
	}

	// Select image from database
	response, err := Dao.SelectById(tx, id)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Select thumbnail if not null from database
	if response.ThumbnailId.Valid {
		thumbnail, err := image.Dao.SelectById(tx, response.ThumbnailId.Int64)
		if err != nil {
			tx.Rollback(context.Background())

			return handleError(c, err)
		}

		response.Thumbnail = null.ValueFrom(thumbnail)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return handleError(c, err)
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
		return handleError(c, err)
	}

	// Update image in database
	response, err := Dao.Update(tx, id, request)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Select thumbnail if not null from database
	if response.ThumbnailId.Valid {
		thumbnail, err := image.Dao.SelectById(tx, response.ThumbnailId.Int64)
		if err != nil {
			tx.Rollback(context.Background())

			return handleError(c, err)
		}

		response.Thumbnail = null.ValueFrom(thumbnail)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return handleError(c, err)
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
		return handleError(c, err)
	}

	// Delete image from database
	response, err := Dao.Delete(tx, id)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Select thumbnail if not null from database
	if response.ThumbnailId.Valid {
		thumbnail, err := image.Dao.SelectById(tx, response.ThumbnailId.Int64)
		if err != nil {
			tx.Rollback(context.Background())

			return handleError(c, err)
		}

		response.Thumbnail = null.ValueFrom(thumbnail)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return handleError(c, err)
	}

	return c.JSON(api.ResponseData(response))
}
