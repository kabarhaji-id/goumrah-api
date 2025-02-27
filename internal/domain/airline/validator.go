package airline

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
	req.SkytraxType = strings.TrimSpace(req.SkytraxType)

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if skytraxTypeLength := len(req.SkytraxType); skytraxTypeLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("skytrax_type", constant.ErrMustBeFilled)
	}
	switch req.SkytraxType {
	case "Full Service", "Low Cost":
		break
	default:
		return CreateRequest{}, errorx.NewValidationError("skytrax_type", constant.ErrInvalidSkytraxType)
	}

	if req.SkytraxRating < 1 || req.SkytraxRating > 5 {
		return CreateRequest{}, errorx.NewValidationError("skytrax_rating", constant.ErrInvalidRating)
	}

	if req.Logo.Valid {
		if logo := req.Logo.Int64; logo < 1 {
			return CreateRequest{}, errorx.NewValidationError("logo", constant.ErrMin1)
		}
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.SkytraxType = strings.TrimSpace(req.SkytraxType)

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if skytraxTypeLength := len(req.SkytraxType); skytraxTypeLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("skytrax_type", constant.ErrMustBeFilled)
	}
	switch req.SkytraxType {
	case "Full Service", "Low Cost":
		break
	default:
		return UpdateRequest{}, errorx.NewValidationError("skytrax_type", constant.ErrInvalidSkytraxType)
	}

	if req.SkytraxRating < 1 || req.SkytraxRating > 5 {
		return UpdateRequest{}, errorx.NewValidationError("skytrax_rating", constant.ErrInvalidRating)
	}

	if req.Logo.Valid {
		if logo := req.Logo.Int64; logo < 1 {
			return UpdateRequest{}, errorx.NewValidationError("logo", constant.ErrMin1)
		}
	}

	return req, nil
}
