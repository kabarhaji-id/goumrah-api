package pkg

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
	req.Description = strings.TrimSpace(req.Description)
	req.Category = strings.TrimSpace(req.Category)
	req.Type = strings.TrimSpace(req.Type)

	if req.Thumbnail.Valid {
		if req.Thumbnail.Int64 < 1 {
			return CreateRequest{}, errorx.NewValidationError("thumbnail", constant.ErrMin1)
		}
	}

	if nameLength := len(req.Name); nameLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return CreateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return CreateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	if categoryLength := len(req.Category); categoryLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("category", constant.ErrMustBeFilled)
	}
	switch req.Category {
	case "Silver", "Gold", "Platinum", "Luxury":
		break
	default:
		return CreateRequest{}, errorx.NewValidationError("category", constant.ErrInvalidPackageCategory)
	}

	if typeLength := len(req.Type); typeLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("type", constant.ErrMustBeFilled)
	}
	switch req.Type {
	case "Reguler", "Plus":
		break
	default:
		return CreateRequest{}, errorx.NewValidationError("type", constant.ErrInvalidPackageType)
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)
	req.Category = strings.TrimSpace(req.Category)
	req.Type = strings.TrimSpace(req.Type)

	if req.Thumbnail.Valid {
		if req.Thumbnail.Int64 < 1 {
			return UpdateRequest{}, errorx.NewValidationError("thumbnail", constant.ErrMin1)
		}
	}

	if nameLength := len(req.Name); nameLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return UpdateRequest{}, errorx.NewValidationError("name", constant.ErrMax100Chars)
	}

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return UpdateRequest{}, errorx.NewValidationError("description", constant.ErrMax500Chars)
	}

	if categoryLength := len(req.Category); categoryLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("category", constant.ErrMustBeFilled)
	}
	switch req.Category {
	case "Silver", "Gold", "Platinum", "Luxury":
		break
	default:
		return UpdateRequest{}, errorx.NewValidationError("category", constant.ErrInvalidPackageCategory)
	}

	if typeLength := len(req.Type); typeLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("type", constant.ErrMustBeFilled)
	}
	switch req.Type {
	case "Reguler", "Plus":
		break
	default:
		return UpdateRequest{}, errorx.NewValidationError("type", constant.ErrInvalidPackageType)
	}

	return req, nil
}
