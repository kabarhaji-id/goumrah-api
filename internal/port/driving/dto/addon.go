package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type AddonRequest struct {
	Category int64
	Name     string
	Price    float64
}

type GetAllAddonRequest struct {
	Page    int
	PerPage int
}

type AddonResponse struct {
	Id       int64
	Category AddonCategoryResponse
	Name     string
	Price    float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
