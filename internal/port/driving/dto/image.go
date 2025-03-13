package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type ImageRequest struct {
	FileData []byte
	FileType string

	Alt      string
	Category null.String
	Title    string
}

type GetAllImageRequest struct {
	Page    int
	PerPage int
}

type ImageResponse struct {
	Id       int64
	Src      string
	Alt      string
	Category null.String
	Title    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
