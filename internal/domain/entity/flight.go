package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type FlightClass string

const (
	FlightEconomy  FlightClass = "Economy"
	FlightBusiness FlightClass = "Business"
	FlightFirst    FlightClass = "First"
)

type Flight struct {
	Id                 int64
	AirlineId          int64
	Aircraft           string
	Baggage            float64
	CabinBaggage       float64
	DepartureAirportId int64
	DepartureTerminal  null.String
	DepartureAt        time.Time
	ArrivalAirportId   int64
	ArrivalTerminal    null.String
	ArrivalAt          time.Time
	Code               string
	SeatLayout         string
	Class              FlightClass

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type FlightRoute struct {
	Id       int64
	FlightId int64
	NextId   null.Int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
