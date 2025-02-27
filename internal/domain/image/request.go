package image

import (
	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type ImageFile struct {
	Name string
	Data []byte
}

type CreateRequest struct {
	Alt      string      `form:"alt"`
	Category null.String `form:"category"`
	Title    string      `form:"title"`
}

type UpdateRequest struct {
	Alt      string      `form:"alt"`
	Category null.String `form:"category"`
	Title    string      `form:"title"`
}

type Params struct {
	ID int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
