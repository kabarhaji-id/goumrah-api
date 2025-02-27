package package_session

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Embarkation   int64  `json:"embarkation"`
	DepartureDate string `json:"departure"`
}

type UpdateRequest struct {
	Embarkation   int64  `json:"embarkation"`
	DepartureDate string `json:"departure"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
