package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type PackageRequest struct {
	Thumbnail     null.Int64 `json:"thumbnail"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"is_active"`
	Category      string     `json:"category"`
	Type          string     `json:"type"`
	IsRecommended bool       `json:"is_recommended"`
	Images        []int64    `json:"images"`
}

func (r PackageRequest) ToDtoRequest() dto.PackageRequest {
	return dto.PackageRequest{
		Thumbnail:     r.Thumbnail,
		Name:          r.Name,
		Description:   r.Description,
		IsActive:      r.IsActive,
		Category:      entity.PackageCategory(r.Category),
		Type:          entity.PackageType(r.Type),
		IsRecommended: r.IsRecommended,
		Images:        r.Images,
	}
}

type GetAllPackageQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type PackageParams struct {
	Id int64 `params:"id"`
}

type PackageResponse struct {
	Id            int64                     `json:"id"`
	Thumbnail     null.Value[ImageResponse] `json:"thumbnail"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	IsActive      bool                      `json:"is_active"`
	Category      string                    `json:"category"`
	Type          string                    `json:"type"`
	Slug          string                    `json:"slug"`
	IsRecommended bool                      `json:"is_recommended"`
	Images        []ImageResponse           `json:"images"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewPackageResponse(dtoResponse dto.PackageResponse) PackageResponse {
	thumbnail := null.NewValue(ImageResponse{}, false)
	if dtoResponse.Thumbnail.Valid {
		imageResponse := NewImageResponse(dtoResponse.Thumbnail.V)

		thumbnail = null.NewValue(imageResponse, true)
	}

	images := NewImageResponses(dtoResponse.Images)

	return PackageResponse{
		Id:            dtoResponse.Id,
		Thumbnail:     thumbnail,
		Name:          dtoResponse.Name,
		Description:   dtoResponse.Description,
		IsActive:      dtoResponse.IsActive,
		Category:      string(dtoResponse.Category),
		Type:          string(dtoResponse.Type),
		Slug:          dtoResponse.Slug,
		IsRecommended: dtoResponse.IsRecommended,
		Images:        images,
		CreatedAt:     dtoResponse.CreatedAt,
		UpdatedAt:     dtoResponse.UpdatedAt,
		DeletedAt:     dtoResponse.DeletedAt,
	}
}

func NewPackageResponses(dtoResponses []dto.PackageResponse) []PackageResponse {
	imageResponses := make([]PackageResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		imageResponses[i] = NewPackageResponse(dtoResponse)
	}

	return imageResponses
}

type PackageListResponse struct {
	Id            int64                     `json:"id"`
	Thumbnail     null.Value[ImageResponse] `json:"thumbnail"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	IsActive      bool                      `json:"is_active"`
	Category      string                    `json:"category"`
	Type          string                    `json:"type"`
	Slug          string                    `json:"slug"`
	IsRecommended bool                      `json:"is_recommended"`
	Images        []int64                   `json:"images"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewPackageListResponse(dtoListResponse dto.PackageListResponse) PackageListResponse {
	thumbnail := null.NewValue(ImageResponse{}, false)
	if dtoListResponse.Thumbnail.Valid {
		imageResponse := NewImageResponse(dtoListResponse.Thumbnail.V)

		thumbnail = null.NewValue(imageResponse, true)
	}

	return PackageListResponse{
		Id:            dtoListResponse.Id,
		Thumbnail:     thumbnail,
		Name:          dtoListResponse.Name,
		Description:   dtoListResponse.Description,
		IsActive:      dtoListResponse.IsActive,
		Category:      string(dtoListResponse.Category),
		Type:          string(dtoListResponse.Type),
		Slug:          dtoListResponse.Slug,
		IsRecommended: dtoListResponse.IsRecommended,
		Images:        dtoListResponse.Images,
		CreatedAt:     dtoListResponse.CreatedAt,
		UpdatedAt:     dtoListResponse.UpdatedAt,
		DeletedAt:     dtoListResponse.DeletedAt,
	}
}

func NewPackageListResponses(dtoListResponses []dto.PackageListResponse) []PackageListResponse {
	packageListResponses := make([]PackageListResponse, len(dtoListResponses))

	for i, dtoResponse := range dtoListResponses {
		packageListResponses[i] = NewPackageListResponse(dtoResponse)
	}

	return packageListResponses
}
