package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type PackageSessionRequest struct {
	Package       int64
	Embarkation   int64
	DepartureDate string
}

type GetAllPackageSessionRequest struct {
	Page    int
	PerPage int
}

type PackageSessionResponse struct {
	Id            int64
	Package       int64
	Embarkation   EmbarkationResponse
	DepartureDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
