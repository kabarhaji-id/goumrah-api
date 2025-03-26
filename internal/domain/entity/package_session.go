package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type PackageSession struct {
	Id                     int64
	PackageId              int64
	EmbarkationId          int64
	DepartureDate          time.Time
	Quota                  int
	DoublePrice            float64
	DoubleFinalPrice       null.Float
	TriplePrice            float64
	TripleFinalPrice       null.Float
	QuadPrice              float64
	QuadFinalPrice         null.Float
	InfantPrice            null.Float
	InfantFinalPrice       null.Float
	DepartureFlightRouteId int64
	ReturnFlightRouteId    int64
	BusId                  int64
	ItineraryId            int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
