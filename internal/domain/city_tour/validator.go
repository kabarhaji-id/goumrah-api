package city_tour

import (
	"strings"

	"github.com/kabarhaji-id/goumrah-api/internal/common/constant"
	"github.com/kabarhaji-id/goumrah-api/internal/common/errorx"
)

type Validator struct {
}

func NewValidator() Validator {
	return Validator{}
}

func (v Validator) ValidateCreateRequest(req CreateRequest) (CreateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.City = strings.TrimSpace(req.City)
	req.Description = strings.TrimSpace(req.Description)

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if cityLength := len(req.City); cityLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("city", constant.ErrMustBeFilled)
	} else if cityLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("city", constant.ErrMax100Chars)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.City = strings.TrimSpace(req.City)
	req.Description = strings.TrimSpace(req.Description)

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if cityLength := len(req.City); cityLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("city", constant.ErrMustBeFilled)
	} else if cityLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("city", constant.ErrMax100Chars)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	return req, nil
}
