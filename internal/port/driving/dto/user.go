package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type UserRequest struct {
	FullName    string
	PhoneNumber string
	Email       string
	Address     string
}

type GetAllUserRequest struct {
	Page    int
	PerPage int
}

type UserResponse struct {
	Id          int64
	FullName    string
	PhoneNumber string
	Email       string
	Address     string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
