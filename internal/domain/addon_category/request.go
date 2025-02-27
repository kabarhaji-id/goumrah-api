package addon_category

import "github.com/kabarhaji-id/goumrah-api/internal/common/api"

type CreateRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	Name string `json:"name"`
}

type Params struct {
	ID int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
