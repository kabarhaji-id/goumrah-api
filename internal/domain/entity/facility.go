package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type Facility struct {
	Id   int64
	Name string
	Icon string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
