package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type BusRequest struct {
	Name string `json:"name"`
	Seat int    `json:"seat"`
}

func (r BusRequest) ToDtoRequest() dto.BusRequest {
	return dto.BusRequest{
		Name: r.Name,
		Seat: r.Seat,
	}
}

type GetAllBusQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type BusParams struct {
	Id int64 `params:"id"`
}

type BusResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Seat int    `json:"seat"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewBusResponse(dtoResponse dto.BusResponse) BusResponse {
	return BusResponse{
		Id:        dtoResponse.Id,
		Name:      dtoResponse.Name,
		Seat:      dtoResponse.Seat,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewBusResponses(dtoResponses []dto.BusResponse) []BusResponse {
	busResponses := make([]BusResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		busResponses[i] = NewBusResponse(dtoResponse)
	}

	return busResponses
}
