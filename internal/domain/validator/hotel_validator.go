package validator

import (
	"context"
	"fmt"
	"net/url"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type HotelValidator struct {
}

func NewHotelValidator() HotelValidator {
	return HotelValidator{}
}

func (v HotelValidator) ValidateRequest(ctx context.Context, request dto.HotelRequest) error {
	nameLength := len(request.Name)
	if nameLength < 1 {
		return newError("Name", mustBeNotEmpty)
	}
	if nameLength > 100 {
		return newError("Name", maxChars(100))
	}

	if request.Rating < 1 || request.Rating > 5 {
		return newError("Rating", mustBetween(1, 5))
	}

	if len(request.Map) < 1 {
		return newError("Map", mustBeNotEmpty)
	}
	if _, err := url.Parse(request.Map); err != nil {
		return newError("Map", invalidUrl)
	}

	addressLength := len(request.Address)
	if addressLength < 1 {
		return newError("Address", mustBeNotEmpty)
	}
	if addressLength > 500 {
		return newError("Address", maxChars(500))
	}

	if request.Distance < 1 {
		return newError("Distance", mustBeGte(1))
	}

	distanceLandmarkLength := len(request.DistanceLandmark)
	if distanceLandmarkLength < 1 {
		return newError("DistanceLandmark", mustBeNotEmpty)
	}
	if distanceLandmarkLength > 100 {
		return newError("DistanceLandmark", maxChars(100))
	}

	if len(request.Review) < 1 {
		return newError("Review", mustBeNotEmpty)
	}
	if _, err := url.Parse(request.Review); err != nil {
		return newError("Review", invalidUrl)
	}

	descriptionLength := len(request.Description)
	if descriptionLength < 1 {
		return newError("Description", mustBeNotEmpty)
	}
	if descriptionLength > 500 {
		return newError("Description", maxChars(500))
	}

	locationLength := len(request.Location)
	if locationLength < 1 {
		return newError("Location", mustBeNotEmpty)
	}
	if locationLength > 100 {
		return newError("Location", maxChars(100))
	}

	for i, image := range request.Images {
		if image < 1 {
			return newError(fmt.Sprintf("Images.%d", i), mustBeGte(1))
		}
	}

	return nil
}

func (v HotelValidator) ValidateId(ctx context.Context, id int64) error {
	if id < 1 {
		return newError("id", mustBeGte(1))
	}

	return nil
}

func (v HotelValidator) ValidateGetAllRequest(ctx context.Context, request dto.GetAllHotelRequest) error {
	if request.Page < 1 {
		return newError("Page", mustBeGte(1))
	}

	if request.PerPage < 1 {
		return newError("PerPage", mustBeGte(1))
	}

	return nil
}
