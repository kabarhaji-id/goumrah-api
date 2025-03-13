package validator

import (
	"context"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FlightValidator struct {
}

func NewFlightValidator() FlightValidator {
	return FlightValidator{}
}

func (v FlightValidator) ValidateRequest(ctx context.Context, request dto.FlightRequest) error {
	if request.Airline < 1 {
		return newError("Airline", mustBeGte(1))
	}

	aircraftLength := len(request.Aircraft)
	if aircraftLength < 1 {
		return newError("Aircraft", mustBeNotEmpty)
	}
	if aircraftLength > 100 {
		return newError("Aircraft", maxChars(100))
	}

	if request.Baggage < 1 {
		return newError("Baggage", mustBeGte(1))
	}

	if request.CabinBaggage < 1 {
		return newError("CabinBaggage", mustBeGte(1))
	}

	if request.DepartureAirport < 1 {
		return newError("DepartureAirportId", mustBeGte(1))
	}

	if request.DepartureTerminal.Valid {
		arrivalTerminalLength := len(request.DepartureTerminal.String)
		if arrivalTerminalLength < 1 {
			return newError("DepartureTerminal", mustBeNotEmpty)
		}
		if arrivalTerminalLength > 100 {
			return newError("DepartureTerminal", maxChars(100))
		}
	}

	if len(request.DepartureAt) < 1 {
		return newError("DepartureAt", mustBeNotEmpty)
	}
	if _, err := time.Parse("15:04", request.DepartureAt); err != nil {
		return newError("DepartureAt", invalidDate("HH:mm"))
	}

	if request.ArrivalAirport < 1 {
		return newError("ArrivalAirportId", mustBeGte(1))
	}

	if request.ArrivalTerminal.Valid {
		arrivalTerminalLength := len(request.ArrivalTerminal.String)
		if arrivalTerminalLength < 1 {
			return newError("ArrivalTerminal", mustBeNotEmpty)
		}
		if arrivalTerminalLength > 100 {
			return newError("ArrivalTerminal", maxChars(100))
		}
	}

	if len(request.ArrivalAt) < 1 {
		return newError("ArrivalAt", mustBeNotEmpty)
	}
	if _, err := time.Parse("15:04", request.ArrivalAt); err != nil {
		return newError("ArrivalAt", invalidDate("HH:mm"))
	}

	codeLength := len(request.Code)
	if codeLength < 1 {
		return newError("Code", mustBeNotEmpty)
	}
	if codeLength > 10 {
		return newError("Code", maxChars(10))
	}

	seatLayoutLength := len(request.SeatLayout)
	if seatLayoutLength < 1 {
		return newError("SeatLayout", mustBeNotEmpty)
	}
	if seatLayoutLength > 10 {
		return newError("SeatLayout", maxChars(10))
	}

	classLength := len(request.Class)
	if classLength < 1 {
		return newError("Class", mustBeNotEmpty)
	}
	switch request.Class {
	case entity.FlightEconomy, entity.FlightBusiness, entity.FlightFirst:
		break
	default:
		return newError("Class", mustBe(
			string(entity.FlightEconomy),
			string(entity.FlightBusiness),
			string(entity.FlightFirst),
		))
	}

	return nil
}

func (v FlightValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v FlightValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllFlightRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
