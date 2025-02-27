package city_tour

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Description string `json:"description"`
}

type UpdateRequest struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Description string `json:"description"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
