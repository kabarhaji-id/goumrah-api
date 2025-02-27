package guide

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
	req.Type = strings.TrimSpace(req.Type)
	req.Description = strings.TrimSpace(req.Description)

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if typeLength := len(req.Type); typeLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("type", constant.ErrMustBeFilled)
	}
	switch req.Type {
	case "Perjalanan", "Ibadah":
		break
	default:
		return CreateRequest{}, errorx.NewValidationError("type", constant.ErrInvalidSkytraxType)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	if req.Avatar.Valid {
		if logo := req.Avatar.Int64; logo < 1 {
			return CreateRequest{}, errorx.NewValidationError("avatar", constant.ErrMin1)
		}
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Type = strings.TrimSpace(req.Type)
	req.Description = strings.TrimSpace(req.Description)

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if typeLength := len(req.Type); typeLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("type", constant.ErrMustBeFilled)
	}
	switch req.Type {
	case "Perjalanan", "Ibadah":
		break
	default:
		return UpdateRequest{}, errorx.NewValidationError("type", constant.ErrInvalidSkytraxType)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	if req.Avatar.Valid {
		if logo := req.Avatar.Int64; logo < 1 {
			return UpdateRequest{}, errorx.NewValidationError("avatar", constant.ErrMin1)
		}
	}

	return req, nil
}
