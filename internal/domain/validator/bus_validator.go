package validator

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type BusValidator struct {
}

func NewBusValidator() BusValidator {
	return BusValidator{}
}

func (v BusValidator) ValidateRequest(ctx context.Context, request dto.BusRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	if request.Seat < 1 {
		return newError("Latitude", mustBeGte(1))
	}

	switch request.Class {
	case entity.BusClassEconomy, entity.BusClassVIP:
		break
	default:
		return newError("Class", mustBe(
			string(entity.BusClassEconomy),
			string(entity.BusClassVIP),
		))
	}

	return nil
}

func (v BusValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v BusValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllBusRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
