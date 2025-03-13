package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type CityTourRequest struct {
	Name        string
	City        string
	Description string
}

type GetAllCityTourRequest struct {
	Page    int
	PerPage int
}

type CityTourResponse struct {
	Id          int64
	Name        string
	City        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
