package package_session

import (
	"time"

	"github.com/guregu/null/v5"
)

type Entity struct {
	Id            int64
	PackageId     int64
	EmbarkationId int64
	DepartureDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
