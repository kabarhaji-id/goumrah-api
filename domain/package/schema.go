package pkg

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/domain/image"
)

type Request struct {
	Thumbnail     null.Int64 `json:"thumbnail"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"is_active"`
	Category      string     `json:"category"`
	Type          string     `json:"type"`
	IsRecommended bool       `json:"is_recommended"`
}

type Response struct {
	Id            int64                      `json:"id"`
	Thumbnail     null.Value[image.Response] `json:"thumbnail"`
	ThumbnailId   null.Int64                 `json:"-"`
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	IsActive      bool                       `json:"is_active"`
	Category      string                     `json:"category"`
	Type          string                     `json:"type"`
	IsRecommended bool                       `json:"is_recommended"`
	Slug          string                     `json:"slug"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
