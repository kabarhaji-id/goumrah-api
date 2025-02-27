package hotel

import (
	"time"

	"github.com/guregu/null/v5"
)

type Response struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Rating      int     `json:"rating"`
	Map         string  `json:"map"`
	Address     string  `json:"address"`
	Distance    float64 `json:"distance"`
	Review      string  `json:"review"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Slug        string  `json:"slug"`

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
