package validator

import (
	"context"
	"fmt"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageValidator struct {
}

func NewPackageValidator() PackageValidator {
	return PackageValidator{}
}

func (v PackageValidator) ValidateRequest(ctx context.Context, request dto.PackageRequest) error {
	if request.Thumbnail.Valid {
		logo := request.Thumbnail.Int64
		if logo < 1 {
			return newError("Thumbnail", mustBeGte(1))
		}
	}

	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	descriptionLength := len(request.Description)
	if descriptionLength < 1 {
		return newError("Description", mustBeNotEmpty)
	}
	if descriptionLength > 500 {
		return newError("Description", maxChars(500))
	}

	categoryLength := len(request.Category)
	if categoryLength < 1 {
		return newError("Category", mustBeNotEmpty)
	}
	switch request.Category {
	case entity.PackageCategorySilver, entity.PackageCategoryGold, entity.PackageCategoryPlatinum, entity.PackageCategoryLuxury:
		break
	default:
		return newError("Category", mustBe(
			string(entity.PackageCategorySilver),
			string(entity.PackageCategoryGold),
			string(entity.PackageCategoryPlatinum),
			string(entity.PackageCategoryLuxury),
		))
	}

	typeLength := len(request.Type)
	if typeLength < 1 {
		return newError("Type", mustBeNotEmpty)
	}
	switch request.Type {
	case entity.PackageTypeReguler, entity.PackageTypePlus:
		break
	default:
		return newError("Type", mustBe(
			string(entity.PackageTypeReguler),
			string(entity.PackageTypePlus),
		))
	}

	for i, image := range request.Images {
		if image < 1 {
			return newError(fmt.Sprintf("Images.%d", i), mustBeGte(1))
		}
	}

	return nil
}

func (v PackageValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v PackageValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllPackageRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
