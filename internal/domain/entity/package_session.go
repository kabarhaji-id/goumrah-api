package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type PackageSession struct {
	Id                     int64
	PackageId              int64
	EmbarkationId          int64
	DepartureDate          time.Time
	DepartureFlightRouteId int64
	ReturnFlightRouteId    int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
