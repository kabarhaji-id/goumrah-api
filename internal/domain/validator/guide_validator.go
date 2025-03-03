package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type GuideValidator struct {
}

func NewGuideValidator() GuideValidator {
	return GuideValidator{}
}

func (v GuideValidator) ValidateRequest(ctx context.Context, request dto.GuideRequest) error {
	if request.Avatar.Valid {
		avatar := request.Avatar.Int64
		if avatar < 1 {
			return newError("Avatar", mustBeGte(1))
		}
	}

	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	typeLength := len(request.Type)
	if typeLength < 1 {
		return newError("Type", mustBeNotEmpty)
	}
	switch request.Type {
	case entity.GuidePerjalanan, entity.GuideIbadah:
		break
	default:
		return newError("Type", mustBe(
			string(entity.GuidePerjalanan),
			string(entity.GuideIbadah),
		))
	}

	descriptionLength := len(request.Description)
	if descriptionLength < 1 {
		return newError("Description", mustBeNotEmpty)
	}
	if descriptionLength > 100 {
		return newError("Description", maxChars(500))
	}

	return nil
}

func (v GuideValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v GuideValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllGuideRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
