package dto

import (
	"time"

	"github.com/guregu/null/v6"
)

type HotelRequest struct {
	Name        string
	Rating      int
	Map         string
	Address     string
	Distance    float64
	Review      string
	Description string
	Location    string
	Images      []int64
}

type GetAllHotelRequest struct {
	Page    int
	PerPage int
}

type HotelResponse struct {
	Id          int64
	Name        string
	Rating      int
	Map         string
	Address     string
	Distance    float64
	Review      string
	Description string
	Location    string
	Slug        string
	Images      []ImageResponse

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
