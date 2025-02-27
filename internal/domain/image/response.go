package image

import (
	"time"

	"github.com/guregu/null/v5"
)

type Response struct {
	Id       int64       `json:"id"`
	Src      string      `json:"src"`
	Alt      string      `json:"alt"`
	Category null.String `json:"category"`
	Title    string      `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
