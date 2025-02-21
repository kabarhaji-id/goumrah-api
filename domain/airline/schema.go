package airline

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/domain/image"
)

type Request struct {
	Name          string     `json:"name"`
	SkytraxType   string     `json:"skytrax_type"`
	SkytraxRating int        `json:"skytrax_rating"`
	Logo          null.Int64 `json:"logo"`
}

type Response struct {
	Id            int64                      `json:"id"`
	Name          string                     `json:"name"`
	SkytraxType   string                     `json:"skytrax_type"`
	SkytraxRating int                        `json:"skytrax_rating"`
	Logo          null.Value[image.Response] `json:"logo"`
	LogoId        null.Int64                 `json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
