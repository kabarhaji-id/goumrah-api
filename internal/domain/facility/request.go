package facility

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type UpdateRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
