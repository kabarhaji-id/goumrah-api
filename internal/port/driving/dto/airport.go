package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type AirportRequest struct {
	City string
	Name string
	Code string
}

type GetAllAirportRequest struct {
	Page    int
	PerPage int
}

type AirportResponse struct {
	Id   int64
	City string
	Name string
	Code string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
