package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type CityTourValidator struct {
}

func NewCityTourValidator() CityTourValidator {
	return CityTourValidator{}
}

func (v CityTourValidator) ValidateRequest(ctx context.Context, request dto.CityTourRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	cityLength := len(request.City)
	if cityLength < 1 {
		return newError("City", mustBeNotEmpty)
	}
	if cityLength > 100 {
		return newError("City", maxChars(100))
	}

	descriptionLength := len(request.Description)
	if descriptionLength < 1 {
		return newError("Description", mustBeNotEmpty)
	}
	if descriptionLength > 500 {
		return newError("Description", maxChars(500))
	}

	return nil
}

func (v CityTourValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v CityTourValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllCityTourRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
