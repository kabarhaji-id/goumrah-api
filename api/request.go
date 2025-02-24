package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
)

type PaginationQuery struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

type paginationQueryOriginal struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

func ValidatePaginationQuery(c *fiber.Ctx) (PaginationQuery, bool, error) {
	queryOriginal := paginationQueryOriginal{}
	if err := c.QueryParser(&queryOriginal); err != nil {
		return PaginationQuery{}, false, ErrInvalidRequestQuery(c, err)
	}

	query := PaginationQuery{
		Page:    int(queryOriginal.Page.Int64),
		PerPage: int(queryOriginal.PerPage.Int64),
	}
	if !queryOriginal.Page.Valid {
		query.Page = 1
	}
	if !queryOriginal.PerPage.Valid {
		query.PerPage = 10
	}

	// Validate page
	if query.Page < 1 {
		return PaginationQuery{}, false, ErrInvalidRequestField(c, "page", "Must be greater than 1")
	}

	// Validate per page
	if query.PerPage < 1 {
		return PaginationQuery{}, false, ErrInvalidRequestField(c, "per_page", "Must be greater than 1")
	}

	return query, true, nil
}

func ValidateId(c *fiber.Ctx, name string) (int64, bool, error) {
	id := c.Params(name, "")
	if id == "" {
		return 0, false, ErrInvalidRequestField(c, name, "Must be filled")
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, false, ErrInvalidRequestField(c, name, "Must be a number", err)
	}

	return idInt, true, nil
}
