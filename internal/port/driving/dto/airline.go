package dto

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type AirlineRequest struct {
	Name          string
	SkytraxType   entity.SkytraxType
	SkytraxRating int
	Logo          null.Int64
}

type GetAllAirlineRequest struct {
	Page    int
	PerPage int
}

type AirlineResponse struct {
	Id            int64
	Name          string
	SkytraxType   entity.SkytraxType
	SkytraxRating int
	Logo          null.Value[ImageResponse]

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
