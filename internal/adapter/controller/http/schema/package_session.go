package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionRequest struct {
	Embarkation      int64              `json:"embarkation"`
	DepartureDate    string             `json:"departure_date"`
	Quota            int                `json:"quota"`
	DoublePrice      float64            `json:"double_price"`
	DoubleFinalPrice null.Float         `json:"double_final_price"`
	TriplePrice      float64            `json:"triple_price"`
	TripleFinalPrice null.Float         `json:"triple_final_price"`
	QuadPrice        float64            `json:"quad_price"`
	QuadFinalPrice   null.Float         `json:"quad_final_price"`
	InfantPrice      null.Float         `json:"infant_price"`
	InfantFinalPrice null.Float         `json:"infant_final_price"`
	DepartureFlights []int64            `json:"departure_flights"`
	ReturnFlights    []int64            `json:"return_flights"`
	Guides           []int64            `json:"guides"`
	Bus              int64              `json:"bus"`
	Itineraries      []ItineraryRequest `json:"itineraries"`
}

func (r PackageSessionRequest) ToDtoRequest() dto.PackageSessionRequest {
	dtoItineraries := make([]dto.ItineraryRequest, len(r.Itineraries))
	for i, itinerary := range r.Itineraries {
		dtoItineraries[i] = itinerary.ToDtoRequest()
	}

	return dto.PackageSessionRequest{
		Embarkation:      r.Embarkation,
		DepartureDate:    r.DepartureDate,
		Quota:            r.Quota,
		DoublePrice:      r.DoublePrice,
		DoubleFinalPrice: r.DoubleFinalPrice,
		TriplePrice:      r.TriplePrice,
		TripleFinalPrice: r.TripleFinalPrice,
		QuadPrice:        r.QuadPrice,
		QuadFinalPrice:   r.QuadFinalPrice,
		InfantPrice:      r.InfantPrice,
		InfantFinalPrice: r.InfantFinalPrice,
		DepartureFlights: r.DepartureFlights,
		ReturnFlights:    r.ReturnFlights,
		Guides:           r.Guides,
		Bus:              r.Bus,
		Itineraries:      dtoItineraries,
	}
}

type GetAllPackageSessionQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type PackageSessionParams struct {
	Id int64 `params:"id"`
}

type PackageSessionResponse struct {
	Id               int64               `json:"id"`
	Package          int64               `json:"package"`
	Embarkation      EmbarkationResponse `json:"embarkation"`
	DepartureDate    time.Time           `json:"departure_date"`
	Quota            int                 `json:"quota"`
	DoublePrice      float64             `json:"double_price"`
	DoubleFinalPrice null.Float          `json:"double_final_price"`
	TriplePrice      float64             `json:"triple_price"`
	TripleFinalPrice null.Float          `json:"triple_final_price"`
	QuadPrice        float64             `json:"quad_price"`
	QuadFinalPrice   null.Float          `json:"quad_final_price"`
	InfantPrice      null.Float          `json:"infant_price"`
	InfantFinalPrice null.Float          `json:"infant_final_price"`
	DepartureFlights []FlightResponse    `json:"departure_flights"`
	ReturnFlights    []FlightResponse    `json:"return_flights"`
	Guides           []GuideResponse     `json:"guides"`
	Bus              BusResponse         `json:"bus"`
	Itineraries      []ItineraryResponse `json:"itineraries"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewPackageSessionResponse(dtoResponse dto.PackageSessionResponse) PackageSessionResponse {
	embarkation := NewEmbarkationResponse(dtoResponse.Embarkation)
	departureFlights := NewFlightResponses(dtoResponse.DepartureFlights)
	returnFlights := NewFlightResponses(dtoResponse.ReturnFlights)
	guides := NewGuideResponses(dtoResponse.Guides)
	bus := NewBusResponse(dtoResponse.Bus)

	return PackageSessionResponse{
		Id:               dtoResponse.Id,
		Package:          dtoResponse.Id,
		Embarkation:      embarkation,
		DepartureDate:    dtoResponse.DepartureDate,
		Quota:            dtoResponse.Quota,
		DoublePrice:      dtoResponse.DoublePrice,
		DoubleFinalPrice: dtoResponse.DoubleFinalPrice,
		TriplePrice:      dtoResponse.TriplePrice,
		TripleFinalPrice: dtoResponse.TripleFinalPrice,
		QuadPrice:        dtoResponse.QuadPrice,
		QuadFinalPrice:   dtoResponse.QuadFinalPrice,
		InfantPrice:      dtoResponse.InfantPrice,
		InfantFinalPrice: dtoResponse.InfantFinalPrice,
		DepartureFlights: departureFlights,
		ReturnFlights:    returnFlights,
		Guides:           guides,
		Bus:              bus,
		Itineraries:      NewItineraryResponses(dtoResponse.Itineraries),
		CreatedAt:        dtoResponse.CreatedAt,
		UpdatedAt:        dtoResponse.UpdatedAt,
		DeletedAt:        dtoResponse.DeletedAt,
	}
}

func NewPackageSessionResponses(dtoResponses []dto.PackageSessionResponse) []PackageSessionResponse {
	packageSessionResponses := make([]PackageSessionResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		packageSessionResponses[i] = NewPackageSessionResponse(dtoResponse)
	}

	return packageSessionResponses
}

type PackageSessionListResponse struct {
	Id               int64               `json:"id"`
	Package          int64               `json:"package"`
	Embarkation      EmbarkationResponse `json:"embarkation"`
	DepartureDate    time.Time           `json:"departure_date"`
	DepartureFlights []int64             `json:"departure_flights"`
	ReturnFlights    []int64             `json:"return_flights"`
	Guides           []int64             `json:"guides"`
	Bus              int64               `json:"bus"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewPackageSessionListResponse(dtoResponse dto.PackageSessionListResponse) PackageSessionListResponse {
	embarkation := NewEmbarkationResponse(dtoResponse.Embarkation)

	return PackageSessionListResponse{
		Id:               dtoResponse.Id,
		Package:          dtoResponse.Id,
		Embarkation:      embarkation,
		DepartureDate:    dtoResponse.DepartureDate,
		DepartureFlights: dtoResponse.DepartureFlights,
		ReturnFlights:    dtoResponse.ReturnFlights,
		Guides:           dtoResponse.Guides,
		Bus:              dtoResponse.Bus,
		CreatedAt:        dtoResponse.CreatedAt,
		UpdatedAt:        dtoResponse.UpdatedAt,
		DeletedAt:        dtoResponse.DeletedAt,
	}
}

func NewPackageSessionListResponses(dtoResponses []dto.PackageSessionListResponse) []PackageSessionListResponse {
	packageSessionListResponses := make([]PackageSessionListResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		packageSessionListResponses[i] = NewPackageSessionListResponse(dtoResponse)
	}

	return packageSessionListResponses
}
