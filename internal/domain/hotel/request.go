package hotel

import (
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type CreateRequest struct {
	Name        string  `json:"name"`
	Rating      int     `json:"rating"`
	Map         string  `json:"map"`
	Address     string  `json:"address"`
	Distance    float64 `json:"distance"`
	Review      string  `json:"review"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
}

type UpdateRequest struct {
	Name        string  `json:"name"`
	Rating      int     `json:"rating"`
	Map         string  `json:"map"`
	Address     string  `json:"address"`
	Distance    float64 `json:"distance"`
	Review      string  `json:"review"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
}

type Params struct {
	Id int64 `params:"id"`
}

type Query struct {
	api.PaginationQuery
}
