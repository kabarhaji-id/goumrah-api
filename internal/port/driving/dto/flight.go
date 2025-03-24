package dto

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type FlightRequest struct {
	Airline           int64
	Aircraft          string
	Baggage           float64
	CabinBaggage      float64
	DepartureAirport  int64
	DepartureTerminal null.String
	DepartureAt       string
	ArrivalAirport    int64
	ArrivalTerminal   null.String
	ArrivalAt         string
	Code              string
	SeatLayout        string
	Class             entity.FlightClass
}

type GetAllFlightRequest struct {
	Page    int
	PerPage int
}

type FlightResponse struct {
	Id                int64
	Airline           AirlineResponse
	Aircraft          string
	Baggage           float64
	CabinBaggage      float64
	DepartureAirport  AirportResponse
	DepartureTerminal null.String
	DepartureAt       time.Time
	ArrivalAirport    AirportResponse
	ArrivalTerminal   null.String
	ArrivalAt         time.Time
	Duration          time.Duration
	Code              string
	SeatLayout        string
	Class             entity.FlightClass

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
