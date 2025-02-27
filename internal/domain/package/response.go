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

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
