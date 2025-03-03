package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirportValidator struct {
}

func NewAirportValidator() AirportValidator {
	return AirportValidator{}
}

func (v AirportValidator) ValidateRequest(ctx context.Context, request dto.AirportRequest) error {
	cityLength := len(request.City)
	if cityLength < 1 {
		return newError("City", mustBeNotEmpty)
	}
	if cityLength > 100 {
		return newError("City", maxChars(100))
	}

	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	codeLength := len(request.Code)
	if codeLength < 1 {
		return newError("Code", mustBeNotEmpty)
	}
	if codeLength != 3 {
		return newError("Code", mustBeChars(3))
	}

	return nil
}

func (v AirportValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v AirportValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllAirportRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
