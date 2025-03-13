package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type AddonCategory struct {
	Id   int64
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
