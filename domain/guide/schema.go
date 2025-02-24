package guide

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/domain/image"
)

type Request struct {
	Avatar      null.Int64 `json:"avatar"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
}

type Response struct {
	Id          int64                      `json:"id"`
	Avatar      null.Value[image.Response] `json:"avatar"`
	AvatarId    null.Int64                 `json:"-"`
	Name        string                     `json:"name"`
	Type        string                     `json:"type"`
	Description string                     `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
