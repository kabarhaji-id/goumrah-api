package airport

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	City string `json:"city"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type UpdateRequest struct {
	City string `json:"city"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
