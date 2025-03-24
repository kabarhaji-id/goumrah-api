package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type PackageCategory string

const (
	PackageCategorySilver   PackageCategory = "Silver"
	PackageCategoryGold     PackageCategory = "Gold"
	PackageCategoryPlatinum PackageCategory = "Platinum"
	PackageCategoryLuxury   PackageCategory = "Luxury"
)

type PackageType string

const (
	PackageTypeReguler PackageType = "Reguler"
	PackageTypePlus    PackageType = "Plus"
)

type Package struct {
	Id            int64
	ThumbnailId   null.Int64
	Name          string
	Description   string
	IsActive      bool
	Category      PackageCategory
	Type          PackageType
	Slug          string
	IsRecommended bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
