package addon

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

	if req.Category < 1 {
		return CreateRequest{}, errorx.NewValidationError("category", constant.ErrMin1)
	}

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if req.Price < 1 {
		return CreateRequest{}, errorx.NewValidationError("price", constant.ErrMin1)
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)

	if req.Category < 1 {
		return UpdateRequest{}, errorx.NewValidationError("category", constant.ErrMin1)
	}

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if req.Price < 1 {
		return UpdateRequest{}, errorx.NewValidationError("price", constant.ErrMin1)
	}

	return req, nil
}
