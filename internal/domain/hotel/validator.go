package hotel

import (
	"net/url"
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
	req.Map = strings.TrimSpace(req.Map)
	req.Address = strings.TrimSpace(req.Address)
	req.Review = strings.TrimSpace(req.Review)
	req.Description = strings.TrimSpace(req.Description)
	req.Location = strings.TrimSpace(req.Location)

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if req.Rating < 1 || req.Rating > 5 {
		return CreateRequest{}, errorx.NewValidationError("rating", constant.ErrInvalidRating)
	}

	if mapLength := len(req.Map); mapLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("map", constant.ErrMustBeFilled)
	}
	if _, err := url.Parse(req.Map); err != nil {
		return CreateRequest{}, errorx.NewValidationError("map", constant.ErrInvalidURL)
	}

	if addressLength := len(req.Address); addressLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("address", constant.ErrMustBeFilled)
	} else if addressLength > 500 {
		return CreateRequest{}, errorx.NewValidationError("address", constant.ErrMax500Chars)
	}

	if req.Distance < 1 {
		return CreateRequest{}, errorx.NewValidationError("distance", constant.ErrMin1)
	}

	if reviewLength := len(req.Review); reviewLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("review", constant.ErrMustBeFilled)
	}
	if _, err := url.Parse(req.Review); err != nil {
		return CreateRequest{}, errorx.NewValidationError("review", constant.ErrInvalidURL)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	if locationLength := len(req.Location); locationLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("location", constant.ErrMustBeFilled)
	} else if locationLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("location", constant.ErrMax100Chars)
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Map = strings.TrimSpace(req.Map)
	req.Address = strings.TrimSpace(req.Address)
	req.Review = strings.TrimSpace(req.Review)
	req.Description = strings.TrimSpace(req.Description)
	req.Location = strings.TrimSpace(req.Location)

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if req.Rating < 1 || req.Rating > 5 {
		return UpdateRequest{}, errorx.NewValidationError("rating", constant.ErrInvalidRating)
	}

	if mapLength := len(req.Map); mapLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("map", constant.ErrMustBeFilled)
	}
	if _, err := url.Parse(req.Map); err != nil {
		return UpdateRequest{}, errorx.NewValidationError("map", constant.ErrInvalidURL)
	}

	if addressLength := len(req.Address); addressLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("address", constant.ErrMustBeFilled)
	} else if addressLength > 500 {
		return UpdateRequest{}, errorx.NewValidationError("address", constant.ErrMax500Chars)
	}

	if req.Distance < 1 {
		return UpdateRequest{}, errorx.NewValidationError("distance", constant.ErrMin1)
	}

	if reviewLength := len(req.Review); reviewLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("review", constant.ErrMustBeFilled)
	}
	if _, err := url.Parse(req.Review); err != nil {
		return UpdateRequest{}, errorx.NewValidationError("review", constant.ErrInvalidURL)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	if locationLength := len(req.Location); locationLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("location", constant.ErrMustBeFilled)
	} else if locationLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("location", constant.ErrMax100Chars)
	}

	return req, nil
}
