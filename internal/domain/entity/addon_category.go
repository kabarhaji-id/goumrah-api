package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type AddonCategory struct {
	Id   int64
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
