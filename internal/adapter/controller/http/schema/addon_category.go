package schema

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type AddonCategoryRequest struct {
	Name string `json:"name"`
}

func (r AddonCategoryRequest) ToDtoRequest() dto.AddonCategoryRequest {
	return dto.AddonCategoryRequest{
		Name: r.Name,
	}
}

type GetAllAddonCategoryQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type AddonCategoryParams struct {
	Id int64 `params:"id"`
}

type AddonCategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewAddonCategoryResponse(dtoResponse dto.AddonCategoryResponse) AddonCategoryResponse {
	return AddonCategoryResponse{
		Id:        dtoResponse.Id,
		Name:      dtoResponse.Name,
		CreatedAt: dtoResponse.CreatedAt,
		UpdatedAt: dtoResponse.UpdatedAt,
		DeletedAt: dtoResponse.DeletedAt,
	}
}

func NewAddonCategoryResponses(dtoResponses []dto.AddonCategoryResponse) []AddonCategoryResponse {
	addonCategoryResponses := make([]AddonCategoryResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		addonCategoryResponses[i] = NewAddonCategoryResponse(dtoResponse)
	}

	return addonCategoryResponses
}
