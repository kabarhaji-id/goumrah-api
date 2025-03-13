package dto

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
)

type GuideRequest struct {
	Avatar      null.Int64
	Name        string
	Type        entity.GuideType
	Description string
}

type GetAllGuideRequest struct {
	Page    int
	PerPage int
}

type GuideResponse struct {
	Id          int64
	Avatar      null.Value[ImageResponse]
	Name        string
	Type        entity.GuideType
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
