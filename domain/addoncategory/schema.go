package addoncategory

import (
	"time"

	"github.com/guregu/null/v5"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
