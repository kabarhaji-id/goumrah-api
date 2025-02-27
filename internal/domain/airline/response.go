package airline

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
)

type Response struct {
	Id            int64                      `json:"id"`
	Name          string                     `json:"name"`
	SkytraxType   SkytraxType                `json:"skytrax_type"`
	SkytraxRating int                        `json:"skytrax_rating"`
	Logo          null.Value[image.Response] `json:"logo"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
