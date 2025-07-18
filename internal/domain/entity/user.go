package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type User struct {
	Id          int64
	FullName    string
	PhoneNumber string
	Email       string
	Address     string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
