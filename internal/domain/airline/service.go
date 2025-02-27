package airline

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
)

type Service struct {
	validator Validator
	uow       *database.UnitOfWork
}

func NewService(validator Validator, uow *database.UnitOfWork) Service {
	return Service{validator, uow}
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
			Name:          req.Name,
			SkytraxType:   SkytraxType(req.SkytraxType),
			SkytraxRating: req.SkytraxRating,
			LogoId:        req.Logo,
		})
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			Name:          entity.Name,
			SkytraxType:   entity.SkytraxType,
			SkytraxRating: entity.SkytraxRating,
			Logo:          null.NewValue(image.Response{}, false),
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.LogoId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.LogoId.Int64)
			if err != nil {
				return err
			}

			response.Logo = null.NewValue(image.Response{
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
			Name:          entity.Name,
			SkytraxType:   entity.SkytraxType,
			SkytraxRating: entity.SkytraxRating,
			Logo:          null.NewValue(image.Response{}, false),
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.LogoId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.LogoId.Int64)
			if err != nil {
				return err
			}

			response.Logo = null.NewValue(image.Response{
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

func (s Service) List(query Query) ([]Response, ListMeta, error) {
	responses := []Response{}
	meta := ListMeta{}

	page := int(query.Page.Int64)
	if !query.Page.Valid {
		page = 1
	}

	perPage := int(query.PerPage.Int64)
	if !query.PerPage.Valid {
		perPage = 10
	}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		imageRepository := image.NewRepository(db)

		count, err := repository.Count(ctx)
		if err != nil {
			return err
		}

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
				Name:          entity.Name,
				SkytraxType:   entity.SkytraxType,
				SkytraxRating: entity.SkytraxRating,
				Logo:          null.NewValue(image.Response{}, false),
				CreatedAt:     entity.CreatedAt,
				UpdatedAt:     entity.UpdatedAt,
				DeletedAt:     entity.DeletedAt,
			}

			if entity.LogoId.Valid {
				imageEntity, err := imageRepository.FindByID(ctx, entity.LogoId.Int64)
				if err != nil {
					return err
				}

				response.Logo = null.NewValue(image.Response{
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

		meta = ListMeta{
			Page:      page,
			PerPage:   perPage,
			FirstPage: 1,
			LastPage:  count/perPage + 1,
			Total:     count,
		}

		return nil
	}); err != nil {
		return nil, ListMeta{}, err
	}

	return responses, meta, nil
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
			Name:          req.Name,
			SkytraxType:   SkytraxType(req.SkytraxType),
			SkytraxRating: req.SkytraxRating,
			LogoId:        req.Logo,
		})
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			Name:          entity.Name,
			SkytraxType:   entity.SkytraxType,
			SkytraxRating: entity.SkytraxRating,
			Logo:          null.NewValue(image.Response{}, false),
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.LogoId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.LogoId.Int64)
			if err != nil {
				return err
			}

			response.Logo = null.NewValue(image.Response{
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
			Name:          entity.Name,
			SkytraxType:   entity.SkytraxType,
			SkytraxRating: entity.SkytraxRating,
			Logo:          null.NewValue(image.Response{}, false),
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		if entity.LogoId.Valid {
			imageEntity, err := imageRepository.FindByID(ctx, entity.LogoId.Int64)
			if err != nil {
				return err
			}

			response.Logo = null.NewValue(image.Response{
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
