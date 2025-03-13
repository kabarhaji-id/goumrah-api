package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type AddonCategoryRequest struct {
	Name string
}

type GetAllAddonCategoryRequest struct {
	Page    int
	PerPage int
}

type AddonCategoryResponse struct {
	Id   int64
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
