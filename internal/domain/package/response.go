package pkg

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
)

type Response struct {
	Id            int64                      `json:"id"`
	Thumbnail     null.Value[image.Response] `json:"thumbnail"`
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	IsActive      bool                       `json:"is_active"`
	Category      Category                   `json:"category"`
	Type          Type                       `json:"type"`
	Slug          string                     `json:"slug"`
	IsRecommended bool                       `json:"is_recommended"`
	Images        []image.Response           `json:"images"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

type ListResponse struct {
	Id            int64                      `json:"id"`
	Thumbnail     null.Value[image.Response] `json:"thumbnail"`
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	IsActive      bool                       `json:"is_active"`
	Category      Category                   `json:"category"`
	Type          Type                       `json:"type"`
	Slug          string                     `json:"slug"`
	IsRecommended bool                       `json:"is_recommended"`
	Images        []int64                    `json:"images"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

type ListMeta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	FirstPage int `json:"first_page"`
	LastPage  int `json:"last_page"`
	Total     int `json:"total"`
}
