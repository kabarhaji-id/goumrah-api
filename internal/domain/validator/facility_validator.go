package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FacilityValidator struct {
}

func NewFacilityValidator() FacilityValidator {
	return FacilityValidator{}
}

func (v FacilityValidator) ValidateRequest(ctx context.Context, request dto.FacilityRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	iconLength := len(request.Icon)
	if iconLength < 1 {
		return newError("Icon", mustBeNotEmpty)
	}
	if iconLength > 100 {
		return newError("Icon", maxChars(100))
	}

	return nil
}

func (v FacilityValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v FacilityValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllFacilityRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
