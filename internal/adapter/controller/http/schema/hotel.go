package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type HotelRequest struct {
	Name             string  `json:"name"`
	Rating           int     `json:"rating"`
	Map              string  `json:"map"`
	Address          string  `json:"address"`
	Distance         float64 `json:"distance"`
	DistanceLandmark string  `json:"distance_landmark"`
	Review           string  `json:"review"`
	Description      string  `json:"description"`
	Location         string  `json:"location"`
	Images           []int64 `json:"images"`
}

func (r HotelRequest) ToDtoRequest() dto.HotelRequest {
	return dto.HotelRequest{
		Name:             r.Name,
		Rating:           r.Rating,
		Map:              r.Map,
		Address:          r.Address,
		Distance:         r.Distance,
		DistanceLandmark: r.DistanceLandmark,
		Review:           r.Review,
		Description:      r.Description,
		Location:         r.Location,
		Images:           r.Images,
	}
}

type GetAllHotelQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type HotelParams struct {
	Id int64 `params:"id"`
}

type HotelResponse struct {
	Id               int64           `json:"id"`
	Name             string          `json:"name"`
	Rating           int             `json:"rating"`
	Map              string          `json:"map"`
	Address          string          `json:"address"`
	Distance         float64         `json:"distance"`
	DistanceLandmark string          `json:"distance_landmark"`
	Review           string          `json:"review"`
	Description      string          `json:"description"`
	Location         string          `json:"location"`
	Images           []ImageResponse `json:"images"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewHotelResponse(dtoResponse dto.HotelResponse) HotelResponse {
	return HotelResponse{
		Id:               dtoResponse.Id,
		Name:             dtoResponse.Name,
		Rating:           dtoResponse.Rating,
		Map:              dtoResponse.Map,
		Address:          dtoResponse.Address,
		Distance:         dtoResponse.Distance,
		DistanceLandmark: dtoResponse.DistanceLandmark,
		Review:           dtoResponse.Review,
		Description:      dtoResponse.Description,
		Location:         dtoResponse.Location,
		Images:           NewImageResponses(dtoResponse.Images),
		CreatedAt:        dtoResponse.CreatedAt,
		UpdatedAt:        dtoResponse.UpdatedAt,
		DeletedAt:        dtoResponse.DeletedAt,
	}
}

func NewHotelResponses(dtoResponses []dto.HotelResponse) []HotelResponse {
	hotelResponses := make([]HotelResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		hotelResponses[i] = NewHotelResponse(dtoResponse)
	}

	return hotelResponses
}
