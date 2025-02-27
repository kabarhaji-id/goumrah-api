package guide

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
)

type Response struct {
	Id          int64                      `json:"id"`
	Avatar      null.Value[image.Response] `json:"avatar"`
	Name        string                     `json:"name"`
	Type        GuideType                  `json:"type"`
	Description string                     `json:"description"`

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
