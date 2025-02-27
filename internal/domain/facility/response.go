package facility

import (
	"time"

	"github.com/guregu/null/v5"
)

type Response struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

type ListMeta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	FirstPage int `json:"first_page"`
	LastPage  int `json:"last_page"`
	Total     int `json:"total"`
}
