package mapper

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/domain/entity"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
)

type UserMapper struct {
}

func NewUserMapper() UserMapper {
	return UserMapper{}
}

func (UserMapper) MapRequestToEntity(ctx context.Context, request dto.UserRequest) entity.User {
	return entity.User{
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Address:     request.Address,
	}
}

func (m UserMapper) MapEntityToResponse(ctx context.Context, userEntity entity.User) dto.UserResponse {
	return dto.UserResponse{
		Id:          userEntity.Id,
		FullName:    userEntity.FullName,
		PhoneNumber: userEntity.PhoneNumber,
		Email:       userEntity.Email,
		Address:     userEntity.Address,
		CreatedAt:   userEntity.CreatedAt,
		UpdatedAt:   userEntity.UpdatedAt,
		DeletedAt:   userEntity.DeletedAt,
	}
}

func (m UserMapper) MapEntitiesToResponses(ctx context.Context, userEntities []entity.User) []dto.UserResponse {
	userResponses := make([]dto.UserResponse, len(userEntities))

	for i, userEntity := range userEntities {
		userResponses[i] = m.MapEntityToResponse(ctx, userEntity)
	}

	return userResponses
}
