package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type EmbarkationValidator struct {
}

func NewEmbarkationValidator() EmbarkationValidator {
	return EmbarkationValidator{}
}

func (v EmbarkationValidator) ValidateRequest(ctx context.Context, request dto.EmbarkationRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	if request.Latitude < -90 || request.Latitude > 90 {
		return newError("Latitude", mustBetween(-90, 90))
	}

	if request.Longitude < -180 || request.Longitude > 180 {
		return newError("Longitude", mustBetween(-180, 180))
	}

	return nil
}

func (v EmbarkationValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v EmbarkationValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllEmbarkationRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
