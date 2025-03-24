package dto

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type PackageRequest struct {
	Thumbnail     null.Int64
	Name          string
	Description   string
	IsActive      bool
	Category      entity.PackageCategory
	Type          entity.PackageType
	IsRecommended bool
	Images        []int64
}

type GetAllPackageRequest struct {
	Page    int
	PerPage int
}

type PackageResponse struct {
	Id            int64
	Thumbnail     null.Value[ImageResponse]
	Name          string
	Description   string
	IsActive      bool
	Category      entity.PackageCategory
	Type          entity.PackageType
	Slug          string
	IsRecommended bool
	Images        []ImageResponse

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type PackageListResponse struct {
	Id            int64
	Thumbnail     null.Value[ImageResponse]
	Name          string
	Description   string
	IsActive      bool
	Category      entity.PackageCategory
	Type          entity.PackageType
	Slug          string
	IsRecommended bool
	Images        []int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
