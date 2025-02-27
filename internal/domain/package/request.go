package pkg

import (
	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Thumbnail     null.Int64 `json:"thumbnail"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"is_active"`
	Category      string     `json:"category"`
	Type          string     `json:"type"`
	IsRecommended bool       `json:"is_recommended"`
}

type UpdateRequest struct {
	Thumbnail     null.Int64 `json:"thumbnail"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"is_active"`
	Category      string     `json:"category"`
	Type          string     `json:"type"`
	IsRecommended bool       `json:"is_recommended"`
}

type Params struct {
	ID int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
