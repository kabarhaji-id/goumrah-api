package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonValidator struct {
}

func NewAddonValidator() AddonValidator {
	return AddonValidator{}
}

func (v AddonValidator) ValidateRequest(ctx context.Context, request dto.AddonRequest) error {
	if request.Category < 1 {
		return newError("Category", mustBeGte(1))
	}

	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	if request.Price < 1 {
		return newError("Price", mustBeGte(1))
	}

	return nil
}

func (v AddonValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v AddonValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllAddonRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
