package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type Bus struct {
	Id   int64
	Name string
	Seat int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
