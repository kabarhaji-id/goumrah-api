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

	unitOfWork repository.UnitOfWork
}

func NewHotelService(
	hotelRepository repository.HotelRepository,
	hotelValidator validator.HotelValidator,
	hotelMapper mapper.HotelMapper,
	unitOfWork repository.UnitOfWork,
) serviceport.HotelService {
	return hotelServiceImpl{
		hotelRepository,
		hotelValidator,
		hotelMapper,
		unitOfWork,
	}
}

func (s hotelServiceImpl) CreateHotel(ctx context.Context, request dto.HotelRequest) (dto.HotelResponse, error) {
	// Validate request
	if err := s.hotelValidator.ValidateRequest(ctx, request); err != nil {
		return dto.HotelResponse{}, err
	}

	// Map request into entity
	hotelEntity := s.hotelMapper.MapRequestToEntity(ctx, request)

	// Create response
	response := dto.HotelResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create hotel repository
		hotelRepository := factory.NewHotelRepository()

		// Create entity with repository
		hotelEntity, err := hotelRepository.Create(ctx, hotelEntity)
		if err != nil {
			return err
		}

		// Create images with repository
		if _, err = hotelRepository.AttachImages(ctx, hotelEntity.Id, request.Images); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.hotelMapper.MapEntityToResponse(ctx, hotelRepository, hotelEntity)

		return err
	})

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
	response, err := s.hotelMapper.MapEntityToResponse(ctx, s.hotelRepository, hotelEntity)
	if err != nil {
		return dto.HotelResponse{}, err
	}

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
	responses, err := s.hotelMapper.MapEntitiesToResponses(ctx, s.hotelRepository, hotelEntities)
	if err != nil {
		return nil, err
	}

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

	// Create response
	response := dto.HotelResponse{}

	// Start transaction with unit of work
	err := s.unitOfWork.Do(ctx, func(ctx context.Context, factory repository.Factory) error {
		// Create hotel repository
		hotelRepository := factory.NewHotelRepository()

		// Update entity with repository
		hotelEntity, err := hotelRepository.Update(ctx, id, hotelEntity)
		if err != nil {
			return err
		}

		// Delete images with repository
		if _, err := hotelRepository.DetachImages(ctx, hotelEntity.Id); err != nil {
			return err
		}

		// Create images with repository
		if _, err = hotelRepository.AttachImages(ctx, hotelEntity.Id, request.Images); err != nil {
			return err
		}

		// Map entity into response
		response, err = s.hotelMapper.MapEntityToResponse(ctx, hotelRepository, hotelEntity)

		return err
	})

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
	response, err := s.hotelMapper.MapEntityToResponse(ctx, s.hotelRepository, hotelEntity)
	if err != nil {
		return dto.HotelResponse{}, err
	}

	return response, err
}
