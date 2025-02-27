package pkg

import (
	"context"
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/embarkation"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/package_session"
)

type Service struct {
	validator        Validator
	sessionValidator package_session.Validator

	uow *database.UnitOfWork
}

func NewService(validator Validator, sessionValidator package_session.Validator, uow *database.UnitOfWork) Service {
	return Service{validator, sessionValidator, uow}
}

func (s Service) Create(req CreateRequest) (Response, error) {
	req, err := s.validator.ValidateCreateRequest(req)
	if err != nil {
		return Response{}, err
	}

	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		imageRepository := image.NewRepository(db)

		entity, err := repository.Create(ctx, Entity{
			ThumbnailId:   req.Thumbnail,
			Name:          req.Name,
			Description:   req.Description,
			IsActive:      req.IsActive,
			Category:      Category(req.Category),
			Type:          Type(req.Type),
			IsRecommended: req.IsRecommended,
		})
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			Thumbnail:     null.NewValue(image.Response{}, false),
			Name:          entity.Name,
			Description:   entity.Description,
			IsActive:      entity.IsActive,
			Category:      entity.Category,
			Type:          entity.Type,
			Slug:          entity.Slug,
			IsRecommended: entity.IsRecommended,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.ThumbnailId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.ThumbnailId.Int64)
			if err != nil {
				return err
			}

			response.Thumbnail = null.NewValue(image.Response{
				Id:        imageEntity.Id,
				Src:       imageEntity.Src,
				Alt:       imageEntity.Alt,
				Category:  imageEntity.Category,
				Title:     imageEntity.Title,
				CreatedAt: imageEntity.CreatedAt,
				UpdatedAt: imageEntity.UpdatedAt,
				DeletedAt: imageEntity.DeletedAt,
			}, true)
		}

		return nil
	}); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (s Service) Get(params Params) (Response, error) {
	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		imageRepository := image.NewRepository(db)

		entity, err := repository.FindByID(ctx, params.ID)
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			Thumbnail:     null.NewValue(image.Response{}, false),
			Name:          entity.Name,
			Description:   entity.Description,
			IsActive:      entity.IsActive,
			Category:      entity.Category,
			Type:          entity.Type,
			Slug:          entity.Slug,
			IsRecommended: entity.IsRecommended,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.ThumbnailId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.ThumbnailId.Int64)
			if err != nil {
				return err
			}

			response.Thumbnail = null.NewValue(image.Response{
				Id:        imageEntity.Id,
				Src:       imageEntity.Src,
				Alt:       imageEntity.Alt,
				Category:  imageEntity.Category,
				Title:     imageEntity.Title,
				CreatedAt: imageEntity.CreatedAt,
				UpdatedAt: imageEntity.UpdatedAt,
				DeletedAt: imageEntity.DeletedAt,
			}, true)
		}

		return nil
	}); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (s Service) List(query Query) ([]Response, error) {
	responses := []Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		imageRepository := image.NewRepository(db)

		entities, err := repository.FindAll(ctx, RepositoryFindAllOption{
			Limit:  query.PerPage,
			Offset: null.NewInt((query.Page.Int64-1)*query.PerPage.Int64, query.Page.Valid),
		})
		if err != nil {
			return err
		}

		for _, entity := range entities {
			response := Response{
				Id:            entity.Id,
				Thumbnail:     null.NewValue(image.Response{}, false),
				Name:          entity.Name,
				Description:   entity.Description,
				IsActive:      entity.IsActive,
				Category:      entity.Category,
				Type:          entity.Type,
				Slug:          entity.Slug,
				IsRecommended: entity.IsRecommended,
				CreatedAt:     entity.CreatedAt,
				UpdatedAt:     entity.UpdatedAt,
				DeletedAt:     entity.DeletedAt,
			}

			if entity.ThumbnailId.Valid {
				imageEntity, err := imageRepository.FindByID(ctx, entity.ThumbnailId.Int64)
				if err != nil {
					return err
				}

				response.Thumbnail = null.NewValue(image.Response{
					Id:        imageEntity.Id,
					Src:       imageEntity.Src,
					Alt:       imageEntity.Alt,
					Category:  imageEntity.Category,
					Title:     imageEntity.Title,
					CreatedAt: imageEntity.CreatedAt,
					UpdatedAt: imageEntity.UpdatedAt,
					DeletedAt: imageEntity.DeletedAt,
				}, true)
			}

			responses = append(responses, response)
		}

		return nil
	}); err != nil {
		return []Response{}, err
	}

	return responses, nil
}

func (s Service) Update(params Params, req UpdateRequest) (Response, error) {
	req, err := s.validator.ValidateUpdateRequest(req)
	if err != nil {
		return Response{}, err
	}

	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		imageRepository := image.NewRepository(db)

		entity, err := repository.Update(ctx, params.ID, Entity{
			ThumbnailId:   req.Thumbnail,
			Name:          req.Name,
			Description:   req.Description,
			IsActive:      req.IsActive,
			Category:      Category(req.Category),
			Type:          Type(req.Type),
			IsRecommended: req.IsRecommended,
		})
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			Thumbnail:     null.NewValue(image.Response{}, false),
			Name:          entity.Name,
			Description:   entity.Description,
			IsActive:      entity.IsActive,
			Category:      entity.Category,
			Type:          entity.Type,
			Slug:          entity.Slug,
			IsRecommended: entity.IsRecommended,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.ThumbnailId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.ThumbnailId.Int64)
			if err != nil {
				return err
			}

			response.Thumbnail = null.NewValue(image.Response{
				Id:        imageEntity.Id,
				Src:       imageEntity.Src,
				Alt:       imageEntity.Alt,
				Category:  imageEntity.Category,
				Title:     imageEntity.Title,
				CreatedAt: imageEntity.CreatedAt,
				UpdatedAt: imageEntity.UpdatedAt,
				DeletedAt: imageEntity.DeletedAt,
			}, true)
		}

		return nil
	}); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (s Service) Delete(params Params) (Response, error) {
	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		imageRepository := image.NewRepository(db)

		entity, err := repository.Delete(ctx, params.ID)
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			Thumbnail:     null.NewValue(image.Response{}, false),
			Name:          entity.Name,
			Description:   entity.Description,
			IsActive:      entity.IsActive,
			Category:      entity.Category,
			Type:          entity.Type,
			Slug:          entity.Slug,
			IsRecommended: entity.IsRecommended,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.ThumbnailId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.ThumbnailId.Int64)
			if err != nil {
				return err
			}

			response.Thumbnail = null.NewValue(image.Response{
				Id:        imageEntity.Id,
				Src:       imageEntity.Src,
				Alt:       imageEntity.Alt,
				Category:  imageEntity.Category,
				Title:     imageEntity.Title,
				CreatedAt: imageEntity.CreatedAt,
				UpdatedAt: imageEntity.UpdatedAt,
				DeletedAt: imageEntity.DeletedAt,
			}, true)
		}

		return nil
	}); err != nil {
		return Response{}, err
	}

	return response, nil
}

func (s Service) CreateSession(params Params, req package_session.CreateRequest) (package_session.Response, error) {
	req, err := s.sessionValidator.ValidateCreateRequest(req)
	if err != nil {
		return package_session.Response{}, err
	}

	response := package_session.Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := package_session.NewRepository(db)
		embarkationRepository := embarkation.NewRepository(db)

		departureDate, err := time.Parse("02/01/2006", req.DepartureDate)
		if err != nil {
			return err
		}

		entity, err := repository.Create(ctx, package_session.Entity{
			PackageId:     params.ID,
			EmbarkationId: req.Embarkation,
			DepartureDate: departureDate,
		})
		if err != nil {
			return err
		}

		response = package_session.Response{
			Id:            entity.Id,
			PackageId:     entity.PackageId,
			Embarkation:   embarkation.Response{},
			DepartureDate: entity.DepartureDate,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		embarkationEntity, err := embarkationRepository.FindByID(ctx, entity.EmbarkationId)
		if err != nil {
			return err
		}
		response.Embarkation = embarkation.Response{
			Id:        embarkationEntity.Id,
			Name:      embarkationEntity.Name,
			Latitude:  embarkationEntity.Latitude,
			Longitude: embarkationEntity.Longitude,
			Slug:      embarkationEntity.Slug,
			CreatedAt: embarkationEntity.CreatedAt,
			UpdatedAt: embarkationEntity.UpdatedAt,
			DeletedAt: embarkationEntity.DeletedAt,
		}

		return nil
	}); err != nil {
		return package_session.Response{}, err
	}

	return response, nil
}

func (s Service) ListSession(params Params, query package_session.Query) ([]package_session.Response, error) {
	responses := []package_session.Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := package_session.NewRepository(db)
		embarkationRepository := embarkation.NewRepository(db)

		entities, err := repository.FindAll(ctx, package_session.RepositoryFindAllOption{
			PackageID: null.IntFrom(params.ID),
			Limit:     query.PerPage,
			Offset:    null.NewInt((query.Page.Int64-1)*query.PerPage.Int64, query.Page.Valid),
		})
		if err != nil {
			return err
		}

		for _, entity := range entities {
			response := package_session.Response{
				Id:            entity.Id,
				PackageId:     entity.PackageId,
				Embarkation:   embarkation.Response{},
				DepartureDate: entity.DepartureDate,
				CreatedAt:     entity.CreatedAt,
				UpdatedAt:     entity.UpdatedAt,
				DeletedAt:     entity.DeletedAt,
			}

			embarkationEntity, err := embarkationRepository.FindByID(ctx, entity.EmbarkationId)
			if err != nil {
				return err
			}

			response.Embarkation = embarkation.Response{
				Id:        embarkationEntity.Id,
				Name:      embarkationEntity.Name,
				Latitude:  embarkationEntity.Latitude,
				Longitude: embarkationEntity.Longitude,
				Slug:      embarkationEntity.Slug,
				CreatedAt: embarkationEntity.CreatedAt,
				UpdatedAt: embarkationEntity.UpdatedAt,
				DeletedAt: embarkationEntity.DeletedAt,
			}

			responses = append(responses, response)
		}

		return nil
	}); err != nil {
		return []package_session.Response{}, err
	}

	return responses, nil
}
