package package_session

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/embarkation"
)

type Response struct {
	Id            int64                `json:"id"`
	PackageId     int64                `json:"package_id"`
	Embarkation   embarkation.Response `json:"embarkation"`
	DepartureDate time.Time            `json:"departure_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
