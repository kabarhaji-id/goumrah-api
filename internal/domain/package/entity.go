package pkg

import (
	"time"

	"github.com/guregu/null/v5"
)

type Category string

const (
	CategorySilver   Category = "Silver"
	CategoryGold     Category = "Gold"
	CategoryPlatinum Category = "Platinum"
	CategoryLuxury   Category = "Luxury"
)

type Type string

const (
	TypeReguler Type = "Reguler"
	TypePlus    Type = "Plus"
)

type Entity struct {
	Id            int64
	ThumbnailId   null.Int64
	Name          string
	Description   string
	IsActive      bool
	Category      Category
	Type          Type
	Slug          string
	IsRecommended bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
