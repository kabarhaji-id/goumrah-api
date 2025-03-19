package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type FacilityRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (r FacilityRequest) ToDtoRequest() dto.FacilityRequest {
	return dto.FacilityRequest{
		Name: r.Name,
		Icon: r.Icon,
	}
}

type GetAllFacilityQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type FacilityParams struct {
	Id int64 `params:"id"`
}

type FacilityResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewFacilityResponse(dtoResponse dto.FacilityResponse) FacilityResponse {
	return FacilityResponse{
		Id:        dtoResponse.Id,
		Name:      dtoResponse.Name,
		Icon:      dtoResponse.Icon,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewFacilityResponses(dtoResponses []dto.FacilityResponse) []FacilityResponse {
	facilityResponses := make([]FacilityResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		facilityResponses[i] = NewFacilityResponse(dtoResponse)
	}

	return facilityResponses
}
