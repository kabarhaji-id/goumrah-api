package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type GuideType string

const (
	GuidePerjalanan GuideType = "Perjalanan"
	GuideIbadah     GuideType = "Ibadah"
)

type Guide struct {
	Id          int64
	AvatarId    null.Int64
	Name        string
	Type        GuideType
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
