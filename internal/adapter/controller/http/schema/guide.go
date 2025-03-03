package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type GuideRequest struct {
	Avatar      null.Int64 `json:"avatar"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
}

func (r GuideRequest) ToDtoRequest() dto.GuideRequest {
	return dto.GuideRequest{
		Avatar:      r.Avatar,
		Name:        r.Name,
		Type:        entity.GuideType(r.Type),
		Description: r.Description,
	}
}

type GetAllGuideQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type GuideParams struct {
	Id int64 `params:"id"`
}

type GuideResponse struct {
	Id          int64                     `json:"id"`
	Avatar      null.Value[ImageResponse] `json:"avatar"`
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Description string                    `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewGuideResponse(dtoResponse dto.GuideResponse) GuideResponse {
	avatar := null.NewValue(ImageResponse{}, false)
	if dtoResponse.Avatar.Valid {
		imageResponse := NewImageResponse(dtoResponse.Avatar.V)

		avatar = null.NewValue(imageResponse, true)
	}

	return GuideResponse{
		Id:          dtoResponse.Id,
		Avatar:      avatar,
		Name:        dtoResponse.Name,
		Type:        string(dtoResponse.Type),
		Description: dtoResponse.Description,
		CreatedAt:   dtoResponse.CreatedAt,
		UpdatedAt:   dtoResponse.UpdatedAt,
		DeletedAt:   dtoResponse.DeletedAt,
	}
}

func NewGuideResponses(dtoResponses []dto.GuideResponse) []GuideResponse {
	guideResponses := make([]GuideResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		guideResponses[i] = NewGuideResponse(dtoResponse)
	}

	return guideResponses
}
