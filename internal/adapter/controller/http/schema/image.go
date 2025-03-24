package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type ImageRequest struct {
	Alt      string      `form:"alt"`
	Category null.String `form:"category"`
	Title    string      `form:"title"`
}

func (r ImageRequest) ToDtoRequest() dto.ImageRequest {
	return dto.ImageRequest{
		Alt:      r.Alt,
		Category: r.Category,
		Title:    r.Title,
	}
}

type GetAllImageQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type ImageParams struct {
	Id int64 `params:"id"`
}

type ImageResponse struct {
	Id       int64       `json:"id"`
	Src      string      `json:"src"`
	Alt      string      `json:"alt"`
	Category null.String `json:"category"`
	Title    string      `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewImageResponse(dtoResponse dto.ImageResponse) ImageResponse {
	return ImageResponse{
		Id:        dtoResponse.Id,
		Src:       dtoResponse.Src,
		Alt:       dtoResponse.Alt,
		Category:  dtoResponse.Category,
		Title:     dtoResponse.Title,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewImageResponses(dtoResponses []dto.ImageResponse) []ImageResponse {
	imageResponses := make([]ImageResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		imageResponses[i] = NewImageResponse(dtoResponse)
	}

	return imageResponses
}
