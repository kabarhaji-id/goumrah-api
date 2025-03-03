package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type CityTour struct {
	Id          int64
	Name        string
	City        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
