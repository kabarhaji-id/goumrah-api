package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type EmbarkationRequest struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (r EmbarkationRequest) ToDtoRequest() dto.EmbarkationRequest {
	return dto.EmbarkationRequest{
		Name:      r.Name,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}

type GetAllEmbarkationQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type EmbarkationParams struct {
	Id int64 `params:"id"`
}

type EmbarkationResponse struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Slug      string  `json:"slug"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewEmbarkationResponse(dtoResponse dto.EmbarkationResponse) EmbarkationResponse {
	return EmbarkationResponse{
		Id:        dtoResponse.Id,
		Name:      dtoResponse.Name,
		Latitude:  dtoResponse.Latitude,
		Longitude: dtoResponse.Longitude,
		Slug:      dtoResponse.Slug,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewEmbarkationResponses(dtoResponses []dto.EmbarkationResponse) []EmbarkationResponse {
	embarkationResponses := make([]EmbarkationResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		embarkationResponses[i] = NewEmbarkationResponse(dtoResponse)
	}

	return embarkationResponses
}
