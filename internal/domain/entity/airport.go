package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type Airport struct {
	Id   int64
	City string
	Name string
	Code string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
