package guide

import (
	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Avatar      null.Int64 `json:"avatar"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
}

type UpdateRequest struct {
	Avatar      null.Int64 `json:"avatar"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
}

type Params struct {
	ID int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
