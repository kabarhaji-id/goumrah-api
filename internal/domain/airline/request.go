package airline

import (
	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Name          string     `json:"name"`
	SkytraxType   string     `json:"skytrax_type"`
	SkytraxRating int        `json:"skytrax_rating"`
	Logo          null.Int64 `json:"logo"`
}

type UpdateRequest struct {
	Name          string     `json:"name"`
	SkytraxType   string     `json:"skytrax_type"`
	SkytraxRating int        `json:"skytrax_rating"`
	Logo          null.Int64 `json:"logo"`
}

type Params struct {
	ID int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
