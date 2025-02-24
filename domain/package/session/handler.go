package pkgsession

import (
	"context"
	"errors"
	"maps"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
	"github.com/kabarhaji-id/goumrah-api/domain/embarkation"
)

func handleError(c *fiber.Ctx, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return api.ErrNotFound(c, err)
	}

	pgError := new(pgconn.PgError)
	if errors.As(err, &pgError) {
		if pgError.Code == "23503" {
			field := ""
			switch pgError.ConstraintName {
			case "package_sessions_package_id_fkey":
				field = "package_id"
			case "package_sessions_embarkation_id_fkey":
				field = "embarkation"
			}

			return api.ErrInvalidRequestField(c, field, "Not found")
		}
	}

	return api.ErrInternalServer(c, err)
}

func CreateHandler(c *fiber.Ctx) error {
	// Get request and validate
	request, success, err := validateRequest(c)
	if !success {
		return err
	}

	// Get package id param and validate
	packageId, success, err := api.ValidateId(c, "package_id")
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return handleError(c, err)
	}

	// Insert package into database
	response, err := Dao.Insert(tx, packageId, request)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Select embarkation from database
	response.Embarkation, err = embarkation.Dao.SelectById(tx, response.EmbarkationId)
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

func GetAllByPackageHandler(c *fiber.Ctx) error {
	// Get request and validate query for pagination
	paginationQuery, success, err := api.ValidatePaginationQuery(c)
	if !success {
		return err
	}

	// Get package id param and validate
	packageId, success, err := api.ValidateId(c, "package_id")
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return handleError(c, err)
	}

	// Select all package sessions from database
	responses, err := Dao.SelectByPackageId(tx, packageId, paginationQuery)
	if err != nil {
		tx.Rollback(context.Background())

		return handleError(c, err)
	}

	// Check if package sessions is not empty
	if len(responses) > 0 {
		// Map all package session embarkation id to package session index
		embarkationMaps := map[int64]int{}
		for index, response := range responses {
			embarkationMaps[response.EmbarkationId] = index
		}

		// Select embarkations from database
		embarkations, err := embarkation.Dao.SelectByIds(tx, slices.Collect(maps.Keys(embarkationMaps)))
		if err != nil {
			tx.Rollback(context.Background())

			return handleError(c, err)
		}

		// Iterate over embarkations and assign package session embarkation
		for _, embarkation := range embarkations {
			responses[embarkationMaps[embarkation.Id]].Embarkation = embarkation
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

func GetAllHandler(c *fiber.Ctx) error {
	// Get request and validate query for pagination
	paginationQuery, success, err := api.ValidatePaginationQuery(c)
	if !success {
		return err
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

	// Check if package sessions is not empty
	if len(responses) > 0 {
		// Map all package session embarkation id to package session index
		embarkationMaps := map[int64]int{}
		for index, response := range responses {
			embarkationMaps[response.EmbarkationId] = index
		}

		// Select embarkations from database
		embarkations, err := embarkation.Dao.SelectByIds(tx, slices.Collect(maps.Keys(embarkationMaps)))
		if err != nil {
			tx.Rollback(context.Background())

			return handleError(c, err)
		}

		// Iterate over embarkations and assign package session embarkation
		for _, embarkation := range embarkations {
			responses[embarkationMaps[embarkation.Id]].Embarkation = embarkation
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
	id, success, err := api.ValidateId(c, "id")
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

	// Select embarkation from database
	response.Embarkation, err = embarkation.Dao.SelectById(tx, response.EmbarkationId)
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
	id, success, err := api.ValidateId(c, "id")
	if !success {
		return err
	}

	// Get request and validate
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

	// Select embarkation from database
	response.Embarkation, err = embarkation.Dao.SelectById(tx, response.EmbarkationId)
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
	id, success, err := api.ValidateId(c, "id")
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

	// Select embarkation from database
	response.Embarkation, err = embarkation.Dao.SelectById(tx, response.EmbarkationId)
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
