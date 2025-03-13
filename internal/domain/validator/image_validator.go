package validator

import (
	"context"
	"strings"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ImageValidator struct {
}

func NewImageValidator() ImageValidator {
	return ImageValidator{}
}

func (v ImageValidator) ValidateRequest(ctx context.Context, request dto.ImageRequest) error {
	if request.FileData == nil {
		return newError("FileData", mustBeNotEmpty)
	}

	if !strings.HasPrefix(request.FileType, "image/") {
		return newError("FileType", "Must be image")
	}

	altLength := len(request.Alt)
	if altLength < 1 {
		return newError("Alt", mustBeNotEmpty)
	}
	if altLength > 100 {
		return newError("Alt", maxChars(100))
	}

	if request.Category.Valid {
		category := request.Category.String
		categoryLength := len(category)
		if categoryLength < 1 {
			return newError("Category", mustBeNotEmpty)
		}
		if categoryLength > 100 {
			return newError("Category", maxChars(100))
		}
	}

	titleLength := len(request.Title)
	if titleLength < 1 {
		return newError("Title", mustBeNotEmpty)
	}
	if titleLength > 100 {
		return newError("Title", maxChars(100))
	}

	return nil
}

func (v ImageValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v ImageValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllImageRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
