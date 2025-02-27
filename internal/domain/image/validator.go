package image

import (
	"strings"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/constant"
	"github.com/kabarhaji-id/goumrah-api/internal/common/errorx"
)

type Validator struct {
}

func NewValidator() Validator {
	return Validator{}
}

func (v Validator) ValidateCreateRequest(req CreateRequest) (CreateRequest, error) {
	req.Alt = strings.TrimSpace(req.Alt)
	if req.Category.Valid {
		req.Category = null.StringFrom(strings.TrimSpace(req.Category.String))
	}
	req.Title = strings.TrimSpace(req.Title)

	if altLength := len(req.Alt); altLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("alt", constant.ErrMustBeFilled)
	} else if altLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("alt", constant.ErrMax100Chars)
	}

	if req.Category.Valid {
		if categoryLength := len(req.Category.String); categoryLength < 1 {
			return CreateRequest{}, errorx.NewValidationError("category", constant.ErrMustBeFilled)
		} else if categoryLength > 100 {
			return CreateRequest{}, errorx.NewValidationError("category", constant.ErrMax100Chars)
		}
	}

	if titleLength := len(req.Title); titleLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("title", constant.ErrMustBeFilled)
	} else if titleLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("title", constant.ErrMax100Chars)
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Alt = strings.TrimSpace(req.Alt)
	if req.Category.Valid {
		req.Category = null.StringFrom(strings.TrimSpace(req.Category.String))
	}
	req.Title = strings.TrimSpace(req.Title)

	if altLength := len(req.Alt); altLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("alt", constant.ErrMustBeFilled)
	} else if altLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("alt", constant.ErrMax100Chars)
	}

	if req.Category.Valid {
		if categoryLength := len(req.Category.String); categoryLength < 1 {
			return UpdateRequest{}, errorx.NewValidationError("category", constant.ErrMustBeFilled)
		} else if categoryLength > 100 {
			return UpdateRequest{}, errorx.NewValidationError("category", constant.ErrMax100Chars)
		}
	}

	if titleLength := len(req.Title); titleLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("title", constant.ErrMustBeFilled)
	} else if titleLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("title", constant.ErrMax100Chars)
	}

	return req, nil
}
