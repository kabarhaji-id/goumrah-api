package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type PackageSessionRequest struct {
	Package          int64
	Embarkation      int64
	DepartureDate    string
	DepartureFlights []int64
	ReturnFlights    []int64
	Guides           []int64
	Bus              int64
}

type GetAllPackageSessionRequest struct {
	Package null.Int64
	Page    int
	PerPage int
}

type PackageSessionResponse struct {
	Id               int64
	Package          int64
	Embarkation      EmbarkationResponse
	DepartureDate    time.Time
	DepartureFlights []FlightResponse
	ReturnFlights    []FlightResponse
	Guides           []GuideResponse
	Bus              BusResponse

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type PackageSessionListResponse struct {
	Id               int64
	Package          int64
	Embarkation      EmbarkationResponse
	DepartureDate    time.Time
	DepartureFlights []int64
	ReturnFlights    []int64
	Guides           []int64
	Bus              int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
