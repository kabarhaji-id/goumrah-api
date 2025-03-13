package entity

import (
	"time"

	"github.com/guregu/null/v5"
)

type Addon struct {
	Id         int64
	CategoryId int64
	Name       string
	Price      float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
