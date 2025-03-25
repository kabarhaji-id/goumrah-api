package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type Hotel struct {
	Id               int64
	Name             string
	Rating           int
	Map              string
	Address          string
	Distance         float64
	DistanceLandmark string
	Review           string
	Description      string
	Location         string
	Slug             string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
