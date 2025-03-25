package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type BusClass string

const (
	BusClassEconomy BusClass = "Economy"
	BusClassVIP     BusClass = "VIP"
)

type Bus struct {
	Id    int64
	Name  string
	Seat  int
	Class BusClass

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
