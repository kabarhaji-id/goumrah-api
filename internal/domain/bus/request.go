package bus

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Name string `json:"name"`
	Seat int    `json:"seat"`
}

type UpdateRequest struct {
	Name string `json:"name"`
	Seat int    `json:"seat"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
