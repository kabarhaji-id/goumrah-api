package package_session

import (
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/common/constant"
	"github.com/kabarhaji-id/goumrah-api/internal/common/errorx"
)

type Validator struct {
}

func NewValidator() Validator {
	return Validator{}
}

func (v Validator) ValidateCreateRequest(req CreateRequest) (CreateRequest, error) {
	if req.Embarkation < 1 {
		return CreateRequest{}, errorx.NewValidationError("embarkation", constant.ErrMin1)
	}

	if departureDateLength := len(req.DepartureDate); departureDateLength < 1 {
		return CreateRequest{}, errorx.NewValidationError("departure_date", constant.ErrMustBeFilled)
	}
	if _, err := time.Parse("02/01/2006", req.DepartureDate); err != nil {
		return CreateRequest{}, errorx.NewValidationError("departure_date", "Must be in format DD/MM/YYYY")
	}

	return req, nil
}

func (v Validator) ValidateUpdateRequest(req UpdateRequest) (UpdateRequest, error) {
	if req.Embarkation < 1 {
		return UpdateRequest{}, errorx.NewValidationError("embarkation", constant.ErrMin1)
	}

	if departureDateLength := len(req.DepartureDate); departureDateLength < 1 {
		return UpdateRequest{}, errorx.NewValidationError("departure_date", constant.ErrMustBeFilled)
	}
	if _, err := time.Parse("02/01/2006", req.DepartureDate); err != nil {
		return UpdateRequest{}, errorx.NewValidationError("departure_date", "Must be in format DD/MM/YYYY")
	}

	return req, nil
}
