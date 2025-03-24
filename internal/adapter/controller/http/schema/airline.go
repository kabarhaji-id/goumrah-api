package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AirlineRequest struct {
	Name          string     `json:"name"`
	SkytraxType   string     `json:"skytrax_type"`
	SkytraxRating int        `json:"skytrax_rating"`
	Logo          null.Int64 `json:"logo"`
}

func (r AirlineRequest) ToDtoRequest() dto.AirlineRequest {
	return dto.AirlineRequest{
		Name:          r.Name,
		SkytraxType:   entity.SkytraxType(r.SkytraxType),
		SkytraxRating: r.SkytraxRating,
		Logo:          r.Logo,
	}
}

type GetAllAirlineQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type AirlineParams struct {
	Id int64 `params:"id"`
}

type AirlineResponse struct {
	Id            int64                     `json:"id"`
	Name          string                    `json:"name"`
	SkytraxType   string                    `json:"skytrax_type"`
	SkytraxRating int                       `json:"skytrax_rating"`
	Logo          null.Value[ImageResponse] `json:"logo"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewAirlineResponse(dtoResponse dto.AirlineResponse) AirlineResponse {
	logo := null.NewValue(ImageResponse{}, false)
	if dtoResponse.Logo.Valid {
		imageResponse := NewImageResponse(dtoResponse.Logo.V)

		logo = null.NewValue(imageResponse, true)
	}

	return AirlineResponse{
		Id:            dtoResponse.Id,
		Name:          dtoResponse.Name,
		SkytraxType:   string(dtoResponse.SkytraxType),
		SkytraxRating: dtoResponse.SkytraxRating,
		Logo:          logo,
		CreatedAt:     dtoResponse.CreatedAt,
		UpdatedAt:     dtoResponse.UpdatedAt,
		DeletedAt:     dtoResponse.DeletedAt,
	}
}

func NewAirlineResponses(dtoResponses []dto.AirlineResponse) []AirlineResponse {
	airlineResponses := make([]AirlineResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		airlineResponses[i] = NewAirlineResponse(dtoResponse)
	}

	return airlineResponses
}
