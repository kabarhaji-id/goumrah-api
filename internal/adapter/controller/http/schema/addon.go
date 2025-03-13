package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonRequest struct {
	Category int64   `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

func (r AddonRequest) ToDtoRequest() dto.AddonRequest {
	return dto.AddonRequest{
		Category: r.Category,
		Name:     r.Name,
		Price:    r.Price,
	}
}

type GetAllAddonQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type AddonParams struct {
	Id int64 `params:"id"`
}

type AddonResponse struct {
	Id       int64                 `json:"id"`
	Category AddonCategoryResponse `json:"category"`
	Name     string                `json:"name"`
	Price    float64               `json:"price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewAddonResponse(dtoResponse dto.AddonResponse) AddonResponse {
	category := NewAddonCategoryResponse(dtoResponse.Category)

	return AddonResponse{
		Id:        dtoResponse.Id,
		Category:  category,
		Name:      dtoResponse.Name,
		Price:     dtoResponse.Price,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewAddonResponses(dtoResponses []dto.AddonResponse) []AddonResponse {
	addonResponses := make([]AddonResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		addonResponses[i] = NewAddonResponse(dtoResponse)
	}

	return addonResponses
}
