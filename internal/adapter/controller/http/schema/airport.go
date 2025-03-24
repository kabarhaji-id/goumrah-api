package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirportRequest struct {
	City string `json:"city"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (r AirportRequest) ToDtoRequest() dto.AirportRequest {
	return dto.AirportRequest{
		City: r.City,
		Name: r.Name,
		Code: r.Code,
	}
}

type GetAllAirportQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type AirportParams struct {
	Id int64 `params:"id"`
}

type AirportResponse struct {
	Id   int64  `json:"id"`
	City string `json:"city"`
	Name string `json:"name"`
	Code string `json:"code"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewAirportResponse(dtoResponse dto.AirportResponse) AirportResponse {
	return AirportResponse{
		Id:        dtoResponse.Id,
		City:      dtoResponse.City,
		Name:      dtoResponse.Name,
		Code:      dtoResponse.Code,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewAirportResponses(dtoResponses []dto.AirportResponse) []AirportResponse {
	airportResponses := make([]AirportResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		airportResponses[i] = NewAirportResponse(dtoResponse)
	}

	return airportResponses
}
