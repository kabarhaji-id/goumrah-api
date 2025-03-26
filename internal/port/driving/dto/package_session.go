package dto

import (
	"time"

	"github.com/guregu/null/v6"
)

type PackageSessionRequest struct {
	Embarkation      int64
	DepartureDate    string
	Quota            int
	DoublePrice      float64
	DoubleFinalPrice null.Float
	TriplePrice      float64
	TripleFinalPrice null.Float
	QuadPrice        float64
	QuadFinalPrice   null.Float
	InfantPrice      null.Float
	InfantFinalPrice null.Float
	DepartureFlights []int64
	ReturnFlights    []int64
	Guides           []int64
	Bus              int64
	Itineraries      []ItineraryRequest
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
	Quota            int
	DoublePrice      float64
	DoubleFinalPrice null.Float
	TriplePrice      float64
	TripleFinalPrice null.Float
	QuadPrice        float64
	QuadFinalPrice   null.Float
	InfantPrice      null.Float
	InfantFinalPrice null.Float
	DepartureFlights []FlightResponse
	ReturnFlights    []FlightResponse
	Guides           []GuideResponse
	Bus              BusResponse
	Itineraries      []ItineraryResponse

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
