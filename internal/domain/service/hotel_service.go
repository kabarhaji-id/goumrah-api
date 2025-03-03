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

type hotelServiceImpl struct {
	hotelRepository repository.HotelRepository
	hotelValidator  validator.HotelValidator
	hotelMapper     mapper.HotelMapper
}

func NewHotelService(
	hotelRepository repository.HotelRepository,
	hotelValidator validator.HotelValidator,
	hotelMapper mapper.HotelMapper,
) serviceport.HotelService {
	return hotelServiceImpl{
		hotelRepository,
		hotelValidator,
		hotelMapper,
	}
}

func (s hotelServiceImpl) CreateHotel(ctx context.Context, request dto.HotelRequest) (dto.HotelResponse, error) {
	// Validate request
	if err := s.hotelValidator.ValidateRequest(ctx, request); err != nil {
		return dto.HotelResponse{}, err
	}

	// Map request into entity
	hotelEntity := s.hotelMapper.MapRequestToEntity(ctx, request)

	// Create entity with repository
	hotelEntity, err := s.hotelRepository.Create(ctx, hotelEntity)
	if err != nil {
		return dto.HotelResponse{}, err
	}

	// Map entity into response
	response := s.hotelMapper.MapEntityToResponse(ctx, hotelEntity)

	return response, err
}

func (s hotelServiceImpl) GetHotelById(ctx context.Context, id int64) (dto.HotelResponse, error) {
	// Validate id
	if err := s.hotelValidator.ValidateId(ctx, id); err != nil {
		return dto.HotelResponse{}, err
	}

	// Find entity by id with repository
	hotelEntity, err := s.hotelRepository.FindById(ctx, id)
	if err != nil {
		return dto.HotelResponse{}, err
	}

	// Map entity into response
	response := s.hotelMapper.MapEntityToResponse(ctx, hotelEntity)

	return response, nil
}

func (s hotelServiceImpl) GetAllHotel(ctx context.Context, request dto.GetAllHotelRequest) ([]dto.HotelResponse, error) {
	// Validate request
	if err := s.hotelValidator.ValidateGetAllRequest(ctx, request); err != nil {
		return nil, err
	}

	// Find all entities with repository
	hotelEntities, err := s.hotelRepository.FindAll(ctx, repository.FindAllOptions{
		Limit:  null.IntFrom(int64(request.PerPage)),
		Offset: null.IntFrom(int64((request.Page - 1) * request.PerPage)),
	})
	if err != nil {
		return nil, err
	}

	// Map entities into responses
	responses := s.hotelMapper.MapEntitiesToResponses(ctx, hotelEntities)

	return responses, nil
}

func (s hotelServiceImpl) UpdateHotel(ctx context.Context, id int64, request dto.HotelRequest) (dto.HotelResponse, error) {
	// Validate id
	if err := s.hotelValidator.ValidateId(ctx, id); err != nil {
		return dto.HotelResponse{}, err
	}

	// Validate request
	if err := s.hotelValidator.ValidateRequest(ctx, request); err != nil {
		return dto.HotelResponse{}, err
	}

	// Map request into entity
	hotelEntity := s.hotelMapper.MapRequestToEntity(ctx, request)

	// Update entity with repository
	hotelEntity, err := s.hotelRepository.Update(ctx, id, hotelEntity)
	if err != nil {
		return dto.HotelResponse{}, err
	}

	// Map entity into response
	response := s.hotelMapper.MapEntityToResponse(ctx, hotelEntity)

	return response, err
}

func (s hotelServiceImpl) DeleteHotel(ctx context.Context, id int64) (dto.HotelResponse, error) {
	// Validate id
	if err := s.hotelValidator.ValidateId(ctx, id); err != nil {
		return dto.HotelResponse{}, err
	}

	// Delete entity with repository
	hotelEntity, err := s.hotelRepository.Delete(ctx, id)
	if err != nil {
		return dto.HotelResponse{}, err
	}

	// Map entity into response
	response := s.hotelMapper.MapEntityToResponse(ctx, hotelEntity)

	return response, err
}
