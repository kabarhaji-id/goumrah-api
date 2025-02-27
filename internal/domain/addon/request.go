package addon

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Category int64   `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

type UpdateRequest struct {
	Category int64   `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
