package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type CityTourRequest struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Description string `json:"description"`
}

func (r CityTourRequest) ToDtoRequest() dto.CityTourRequest {
	return dto.CityTourRequest{
		Name:        r.Name,
		City:        r.City,
		Description: r.Description,
	}
}

type GetAllCityTourQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type CityTourParams struct {
	Id int64 `params:"id"`
}

type CityTourResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewCityTourResponse(dtoResponse dto.CityTourResponse) CityTourResponse {
	return CityTourResponse{
		Id:          dtoResponse.Id,
		Name:        dtoResponse.Name,
		City:        dtoResponse.City,
		Description: dtoResponse.Description,
		CreatedAt:   dtoResponse.CreatedAt,
		UpdatedAt:   dtoResponse.UpdatedAt,
		DeletedAt:   dtoResponse.DeletedAt,
	}
}

func NewCityTourResponses(dtoResponses []dto.CityTourResponse) []CityTourResponse {
	cityTourResponses := make([]CityTourResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		cityTourResponses[i] = NewCityTourResponse(dtoResponse)
	}

	return cityTourResponses
}
