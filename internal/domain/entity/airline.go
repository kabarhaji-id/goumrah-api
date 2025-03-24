package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type SkytraxType string

const (
	SkytraxFullService SkytraxType = "Full Service"
	SkytraxLowCost     SkytraxType = "Low Cost"
)

type Airline struct {
	Id            int64
	Name          string
	SkytraxType   SkytraxType
	SkytraxRating int
	LogoId        null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
