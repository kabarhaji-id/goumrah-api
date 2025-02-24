package image

import (
	"context"
	"crypto/rand"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/database"
)

func CreateHandler(c *fiber.Ctx) error {
	// Get request and validate
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

	// Insert image into database
	response, err := Dao.Insert(tx, imageFileName, request)
	if err != nil {
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
	// Get request and validate query for pagination
	paginationQuery, success, err := api.ValidatePaginationQuery(c)
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Select all images from database
	responses, err := Dao.SelectAll(tx, paginationQuery)
	if err != nil {
		tx.Rollback(context.Background())

		return api.ErrInternalServer(c, err)
	}

	// Count all images from database
	count, err := Dao.CountAll(tx)
	if err != nil {
		tx.Rollback(context.Background())

		return api.ErrInternalServer(c, err)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
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
	id, success, err := api.ValidateId(c, "id")
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Select image from database
	response, err := Dao.SelectById(tx, id)
	if err != nil {
		tx.Rollback(context.Background())

		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound(c, err)
		}

		return api.ErrInternalServer(c, err)
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return api.ErrInternalServer(c, err)
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
		return api.ErrInternalServer(c, err)
	}

	// Update image in database
	response, err := Dao.Update(tx, id, request)
	if err != nil {
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
	id, success, err := api.ValidateId(c, "id")
	if !success {
		return err
	}

	// Start transaction
	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		return api.ErrInternalServer(c, err)
	}

	// Delete image from database
	response, err := Dao.Delete(tx, id)
	if err != nil {
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
