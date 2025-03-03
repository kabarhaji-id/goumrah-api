package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirlineValidator struct {
}

func NewAirlineValidator() AirlineValidator {
	return AirlineValidator{}
}

func (v AirlineValidator) ValidateRequest(ctx context.Context, request dto.AirlineRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	skytraxTypeLength := len(request.SkytraxType)
	if skytraxTypeLength < 1 {
		return newError("SkytraxType", mustBeNotEmpty)
	}
	switch request.SkytraxType {
	case entity.SkytraxFullService, entity.SkytraxLowCost:
		break
	default:
		return newError("SkytraxType", mustBe(
			string(entity.SkytraxFullService),
			string(entity.SkytraxLowCost),
		))
	}

	if request.SkytraxRating < 1 || request.SkytraxRating > 5 {
		return newError("SkytraxRating", mustBetween(1, 5))
	}

	if request.Logo.Valid {
		logo := request.Logo.Int64
		if logo < 1 {
			return newError("Logo", mustBeGte(1))
		}
	}

	return nil
}

func (v AirlineValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v AirlineValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllAirlineRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
