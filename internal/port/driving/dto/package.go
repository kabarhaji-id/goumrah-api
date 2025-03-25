package dto

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type PackageRequest struct {
	Thumbnail null.Int64
	Name      string
	Category  entity.PackageCategory
	Type      entity.PackageType
	Images    []int64
}

type GetAllPackageRequest struct {
	Page    int
	PerPage int
}

type PackageResponse struct {
	Id        int64
	Thumbnail null.Value[ImageResponse]
	Name      string
	Category  entity.PackageCategory
	Type      entity.PackageType
	Slug      string
	Images    []ImageResponse

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type PackageListResponse struct {
	Id        int64
	Thumbnail null.Value[ImageResponse]
	Name      string
	Category  entity.PackageCategory
	Type      entity.PackageType
	Slug      string
	Images    []int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
