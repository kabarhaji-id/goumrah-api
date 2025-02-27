package airport

import (
	"time"

	"github.com/guregu/null/v5"
)

type Response struct {
	Id   int64  `json:"id"`
	City string `json:"city"`
	Name string `json:"name"`
	Code string `json:"code"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
