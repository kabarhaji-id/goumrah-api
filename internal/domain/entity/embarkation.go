package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type Embarkation struct {
	Id        int64
	Name      string
	Latitude  float64
	Longitude float64
	Slug      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
