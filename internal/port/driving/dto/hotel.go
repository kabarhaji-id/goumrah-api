package dto

import (
	"time"

	"github.com/guregu/null/v5"
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

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
