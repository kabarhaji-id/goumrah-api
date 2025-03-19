package schema

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type UserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func (r UserRequest) ToDtoRequest() dto.UserRequest {
	return dto.UserRequest{
		FullName:    r.FullName,
		PhoneNumber: r.PhoneNumber,
		Email:       r.Email,
		Address:     r.Address,
	}
}

type GetAllUserQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}

type UserParams struct {
	Id int64 `params:"id"`
}

type UserResponse struct {
	Id          int64  `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewUserResponse(dtoResponse dto.UserResponse) UserResponse {
	return UserResponse{
		Id:          dtoResponse.Id,
		FullName:    dtoResponse.FullName,
		PhoneNumber: dtoResponse.PhoneNumber,
		Email:       dtoResponse.Email,
		Address:     dtoResponse.Address,
		CreatedAt:   dtoResponse.CreatedAt,
		UpdatedAt:   dtoResponse.UpdatedAt,
		DeletedAt:   dtoResponse.DeletedAt,
	}
}

func NewUserResponses(dtoResponses []dto.UserResponse) []UserResponse {
	userResponses := make([]UserResponse, len(dtoResponses))

	for i, dtoResponse := range dtoResponses {
		userResponses[i] = NewUserResponse(dtoResponse)
	}

	return userResponses
}
