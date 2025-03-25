package dto

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type BusRequest struct {
	Name  string
	Seat  int
	Class entity.BusClass
}

type GetAllBusRequest struct {
	Page    int
	PerPage int
}

type BusResponse struct {
	Id    int64
	Name  string
	Seat  int
	Class entity.BusClass

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
