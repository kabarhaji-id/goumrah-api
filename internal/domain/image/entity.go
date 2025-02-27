package image

import (
	"time"

	"github.com/guregu/null/v5"
)

type Entity struct {
	Id       int64
	Src      string
	Alt      string
	Category null.String
	Title    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
