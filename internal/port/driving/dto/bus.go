package dto

import (
	"time"

	"github.com/guregu/null/v6"
)

type BusRequest struct {
	Name string
	Seat int
}

type GetAllBusRequest struct {
	Page    int
	PerPage int
}

type BusResponse struct {
	Id   int64
	Name string
	Seat int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
