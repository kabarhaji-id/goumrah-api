package bus

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
)

func handleError(c *fiber.Ctx, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return api.ErrNotFound(c, err)
	}

	pgError := new(pgconn.PgError)
	if errors.As(err, &pgError) {
		if pgError.Code == "23505" && pgError.ConstraintName == "buses_name_unique" {
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

	// Insert bus into database
	response, err := Dao.Insert(tx, request)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
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
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return handleError(c, err)
	}

	// Select all buses from database
	responses, err := Dao.SelectAll(tx, paginationQuery)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Count all buses from database
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

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return handleError(c, err)
	}

	return c.JSON(api.ResponseData(response))
}
