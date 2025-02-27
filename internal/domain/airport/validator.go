package airport

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
	req.City = strings.TrimSpace(req.City)
	req.Name = strings.TrimSpace(req.Name)
	req.Code = strings.TrimSpace(req.Code)

	if cityLength := len(req.City); cityLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("city", constant.ErrMustBeFilled)
	} else if cityLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("city", constant.ErrMax100Chars)
	}

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if codeLength := len(req.Code); codeLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("code", constant.ErrMustBeFilled)
	} else if codeLength != 3 {
		return CreateRequest{}, errorx.NewValidationError("code", constant.ErrInvalidAirportCode)
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.City = strings.TrimSpace(req.City)
	req.Name = strings.TrimSpace(req.Name)
	req.Code = strings.TrimSpace(req.Code)

	if cityLength := len(req.City); cityLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("city", constant.ErrMustBeFilled)
	} else if cityLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("city", constant.ErrMax100Chars)
	}

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if codeLength := len(req.Code); codeLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("code", constant.ErrMustBeFilled)
	} else if codeLength != 3 {
		return UpdateRequest{}, errorx.NewValidationError("code", constant.ErrInvalidAirportCode)
	}

	return req, nil
}
