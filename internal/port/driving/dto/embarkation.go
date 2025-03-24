package dto

import (
	"time"

	"github.com/guregu/null/v6"
)

type EmbarkationRequest struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type GetAllEmbarkationRequest struct {
	Page    int
	PerPage int
}

type EmbarkationResponse struct {
	Id        int64
	Name      string
	Latitude  float64
	Longitude float64
	Slug      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
