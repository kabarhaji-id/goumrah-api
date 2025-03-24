package dto

import (
	"time"

	"github.com/guregu/null/v6"
)

type FacilityRequest struct {
	Name string
	Icon string
}

type GetAllFacilityRequest struct {
	Page    int
	PerPage int
}

type FacilityResponse struct {
	Id   int64
	Name string
	Icon string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
