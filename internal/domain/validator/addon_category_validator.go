package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonCategoryValidator struct {
}

func NewAddonCategoryValidator() AddonCategoryValidator {
	return AddonCategoryValidator{}
}

func (v AddonCategoryValidator) ValidateRequest(ctx context.Context, request dto.AddonCategoryRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	return nil
}

func (v AddonCategoryValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v AddonCategoryValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllAddonCategoryRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
