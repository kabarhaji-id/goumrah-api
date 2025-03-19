package service

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, request dto.UserRequest) (dto.UserResponse, error)
	GetUserById(ctx context.Context, id int64) (dto.UserResponse, error)
	GetAllUser(ctx context.Context, request dto.GetAllUserRequest) ([]dto.UserResponse, error)
	UpdateUser(ctx context.Context, id int64, request dto.UserRequest) (dto.UserResponse, error)
	DeleteUser(ctx context.Context, id int64) (dto.UserResponse, error)
}
