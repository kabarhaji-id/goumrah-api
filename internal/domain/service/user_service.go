package service

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	serviceport "github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
	userValidator  validator.UserValidator
	userMapper     mapper.UserMapper
}

func NewUserService(
	userRepository repository.UserRepository,
	userValidator validator.UserValidator,
	userMapper mapper.UserMapper,
) serviceport.UserService {
	return userServiceImpl{
		userRepository,
		userValidator,
		userMapper,
	}
}

func (s userServiceImpl) CreateUser(ctx context.Context, request dto.UserRequest) (dto.UserResponse, error) {
	// Validate request
	if err := s.userValidator.ValidateRequest(ctx, request); err != nil {
		return dto.UserResponse{}, err
	}

	// Map request into entity
	userEntity := s.userMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	userEntity, err := s.userRepository.Create(ctx, userEntity)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Map entity into response
	response := s.userMapper.MapEntityToResponse(ctx, userEntity)

	return response, err
}

func (s userServiceImpl) GetUserById(ctx context.Context, id int64) (dto.UserResponse, error) {
	// Validate id
	if err := s.userValidator.ValidateId(ctx, id); err != nil {
		return dto.UserResponse{}, err
	}

	// Find entity by id with repository
	userEntity, err := s.userRepository.FindById(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Map entity into response
	response := s.userMapper.MapEntityToResponse(ctx, userEntity)

	return response, nil
}

func (s userServiceImpl) GetAllUser(ctx context.Context, request dto.GetAllUserRequest) ([]dto.UserResponse, error) {
	// Validate request
	if err := s.userValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	userEntities, err := s.userRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.userMapper.MapEntitiesToResponses(ctx, userEntities)

	return responses, nil
}

func (s userServiceImpl) UpdateUser(ctx context.Context, id int64, request dto.UserRequest) (dto.UserResponse, error) {
	// Validate id
	if err := s.userValidator.ValidateId(ctx, id); err != nil {
		return dto.UserResponse{}, err
	}

	// Validate request
	if err := s.userValidator.ValidateRequest(ctx, request); err != nil {
		return dto.UserResponse{}, err
	}

	// Map request into entity
	userEntity := s.userMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	userEntity, err := s.userRepository.Update(ctx, id, userEntity)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Map entity into response
	response := s.userMapper.MapEntityToResponse(ctx, userEntity)

	return response, err
}

func (s userServiceImpl) DeleteUser(ctx context.Context, id int64) (dto.UserResponse, error) {
	// Validate id
	if err := s.userValidator.ValidateId(ctx, id); err != nil {
		return dto.UserResponse{}, err
	}

	// Delete entity with repository
	userEntity, err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Map entity into response
	response := s.userMapper.MapEntityToResponse(ctx, userEntity)

	return response, err
}
