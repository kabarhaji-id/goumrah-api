package validator

import (
	"context"
	"fmt"
	"time"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageSessionValidator struct {
}

func NewPackageSessionValidator() PackageSessionValidator {
	return PackageSessionValidator{}
}

func (v PackageSessionValidator) ValidateRequest(ctx context.Context, request dto.PackageSessionRequest) error {
	if request.Package < 1 {
		return newError("Package", mustBeGte(1))
	}

	if request.Embarkation < 1 {
		return newError("Embarkation", mustBeGte(1))
	}

	if len(request.DepartureDate) < 1 {
		return newError("DepartureDate", mustBeNotEmpty)
	}
	if _, err := time.Parse("02/01/2006", request.DepartureDate); err != nil {
		return newError("DepartureDate", invalidDate("DD/MM/YYYY"))
	}

	for i, guide := range request.Guides {
		if guide < 1 {
			return newError(fmt.Sprintf("Guides.%d", i), mustBeGte(1))
		}
	}

	return nil
}

func (v PackageSessionValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v PackageSessionValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllPackageSessionRequest) error {
	if request.Package.Valid {
		if request.Package.Int64 < 1 {
			return newError("Package", mustBeGte(1))
		}
	}

	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
