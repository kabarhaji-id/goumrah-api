package package_session

import (
	"context"
	"time"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/embarkation"
)

type Service struct {
	validator Validator
	uow       *database.UnitOfWork
}

func NewService(validator Validator, uow *database.UnitOfWork) Service {
	return Service{validator, uow}
}

func (s Service) Get(params Params) (Response, error) {
	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		embarkationRepository := embarkation.NewRepository(db)

		entity, err := repository.FindById(ctx, params.Id)
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			PackageId:     entity.PackageId,
			Embarkation:   embarkation.Response{},
			DepartureDate: entity.DepartureDate,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		embarkationEntity, err := embarkationRepository.FindById(ctx, entity.EmbarkationId)
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
		embarkationRepository := embarkation.NewRepository(db)

		count, err := repository.Count(ctx, RepositoryCountOption{})
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
				PackageId:     entity.PackageId,
				Embarkation:   embarkation.Response{},
				DepartureDate: entity.DepartureDate,
				CreatedAt:     entity.CreatedAt,
				UpdatedAt:     entity.UpdatedAt,
				DeletedAt:     entity.DeletedAt,
			}

			embarkationEntity, err := embarkationRepository.FindById(ctx, entity.EmbarkationId)
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
		embarkationRepository := embarkation.NewRepository(db)

		departureDate, err := time.Parse("02/01/2006", req.DepartureDate)
		if err != nil {
			return err
		}

		entity, err := repository.Update(ctx, params.Id, Entity{
			EmbarkationId: req.Embarkation,
			DepartureDate: departureDate,
		})
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			PackageId:     entity.PackageId,
			Embarkation:   embarkation.Response{},
			DepartureDate: entity.DepartureDate,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		embarkationEntity, err := embarkationRepository.FindById(ctx, entity.EmbarkationId)
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
		return Response{}, err
	}

	return response, nil
}

func (s Service) Delete(params Params) (Response, error) {
	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)
		embarkationRepository := embarkation.NewRepository(db)

		entity, err := repository.Delete(ctx, params.Id)
		if err != nil {
			return err
		}

		response = Response{
			Id:            entity.Id,
			PackageId:     entity.PackageId,
			Embarkation:   embarkation.Response{},
			DepartureDate: entity.DepartureDate,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
			DeletedAt:     entity.DeletedAt,
		}

		embarkationEntity, err := embarkationRepository.FindById(ctx, entity.EmbarkationId)
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
		return Response{}, err
	}

	return response, nil
}
