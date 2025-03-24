package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FlightRequest struct {
	Airline           int64              `json:"airline"`
	Aircraft          string             `json:"aircraft"`
	Baggage           float64            `json:"baggage"`
	CabinBaggage      float64            `json:"cabin_baggage"`
	DepartureAirport  int64              `json:"departure_airport"`
	DepartureTerminal null.String        `json:"departure_terminal"`
	DepartureAt       string             `json:"departure_at"`
	ArrivalAirport    int64              `json:"arrival_airport"`
	ArrivalTerminal   null.String        `json:"arrival_terminal"`
	ArrivalAt         string             `json:"arrival_at"`
	Code              string             `json:"code"`
	SeatLayout        string             `json:"seat_layout"`
	Class             entity.FlightClass `json:"class"`
}

func (r FlightRequest) ToDtoRequest() dto.FlightRequest {
	return dto.FlightRequest{
		Airline:           r.Airline,
		Aircraft:          r.Aircraft,
		Baggage:           r.Baggage,
		CabinBaggage:      r.CabinBaggage,
		DepartureAirport:  r.DepartureAirport,
		DepartureTerminal: r.DepartureTerminal,
		DepartureAt:       r.DepartureAt,
		ArrivalAirport:    r.ArrivalAirport,
		ArrivalTerminal:   r.ArrivalTerminal,
		ArrivalAt:         r.ArrivalAt,
		Code:              r.Code,
		SeatLayout:        r.SeatLayout,
		Class:             r.Class,
	}
}

type GetAllFlightQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type FlightParams struct {
	Id int64 `params:"id"`
}

type FlightResponse struct {
	Id                int64              `json:"id"`
	Airline           AirlineResponse    `json:"airline"`
	Aircraft          string             `json:"aircraft"`
	Baggage           float64            `json:"baggage"`
	CabinBaggage      float64            `json:"cabin_baggage"`
	DepartureAirport  AirportResponse    `json:"departure_airport"`
	DepartureTerminal null.String        `json:"departure_terminal"`
	DepartureAt       string             `json:"departure_at"`
	ArrivalAirport    AirportResponse    `json:"arrival_airport"`
	ArrivalTerminal   null.String        `json:"arrival_terminal"`
	ArrivalAt         string             `json:"arrival_at"`
	Duration          int                `json:"duration"`
	Code              string             `json:"code"`
	SeatLayout        string             `json:"seat_layout"`
	Class             entity.FlightClass `json:"class"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewFlightResponse(dtoResponse dto.FlightResponse) FlightResponse {
	airline := NewAirlineResponse(dtoResponse.Airline)
	departureAirport := NewAirportResponse(dtoResponse.DepartureAirport)
	arrivalAirport := NewAirportResponse(dtoResponse.ArrivalAirport)

	return FlightResponse{
		Id:                dtoResponse.Id,
		Airline:           airline,
		Aircraft:          dtoResponse.Aircraft,
		Baggage:           dtoResponse.Baggage,
		CabinBaggage:      dtoResponse.CabinBaggage,
		DepartureAirport:  departureAirport,
		DepartureTerminal: dtoResponse.DepartureTerminal,
		DepartureAt:       dtoResponse.DepartureAt.Format("15:04"),
		ArrivalAirport:    arrivalAirport,
		ArrivalTerminal:   dtoResponse.ArrivalTerminal,
		ArrivalAt:         dtoResponse.ArrivalAt.Format("15:04"),
		Duration:          int(dtoResponse.Duration.Minutes()),
		Code:              dtoResponse.Code,
		SeatLayout:        dtoResponse.SeatLayout,
		Class:             dtoResponse.Class,
		CreatedAt:         dtoResponse.CreatedAt,
		UpdatedAt:         dtoResponse.UpdatedAt,
		DeletedAt:         dtoResponse.DeletedAt,
	}
}

func NewFlightResponses(dtoResponses []dto.FlightResponse) []FlightResponse {
	addonResponses := make([]FlightResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		addonResponses[i] = NewFlightResponse(dtoResponse)
	}

	return addonResponses
}
