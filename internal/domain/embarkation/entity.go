package embarkation

import (
	"time"

	"github.com/guregu/null/v5"
)

type Entity struct {
	Id        int64
	Name      string
	Latitude  float64
	Longitude float64
	Slug      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
