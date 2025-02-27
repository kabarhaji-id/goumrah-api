package image

import (
	"context"
	"crypto/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
)

type Service struct {
	validator Validator
	uow       *database.UnitOfWork
}

func NewService(validator Validator, uow *database.UnitOfWork) Service {
	return Service{validator, uow}
}

func (s Service) Create(imageFile ImageFile, req CreateRequest) (Response, error) {
	req, err := s.validator.ValidateCreateRequest(req)
	if err != nil {
		return Response{}, err
	}

	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)

		imageFileName := strings.ToLower(rand.Text()) + filepath.Ext(imageFile.Name)

		entity, err := repository.Create(ctx, Entity{
			Src:      imageFileName,
			Alt:      req.Alt,
			Category: req.Category,
			Title:    req.Title,
		})
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join("public", imageFileName), imageFile.Data, 0644); err != nil {
			return err
		}

		response = Response{
			Id:        entity.Id,
			Src:       entity.Src,
			Alt:       entity.Alt,
			Category:  entity.Category,
			Title:     entity.Title,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
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

		entity, err := repository.FindByID(ctx, params.ID)
		if err != nil {
			return err
		}

		response = Response{
			Id:        entity.Id,
			Src:       entity.Src,
			Alt:       entity.Alt,
			Category:  entity.Category,
			Title:     entity.Title,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
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

		entities, err := repository.FindAll(ctx, RepositoryFindAllOption{
			Limit:  query.PerPage,
			Offset: null.NewInt((query.Page.Int64-1)*query.PerPage.Int64, query.Page.Valid),
		})
		if err != nil {
			return err
		}

		for _, entity := range entities {
			responses = append(responses, Response{
				Id:        entity.Id,
				Src:       entity.Src,
				Alt:       entity.Alt,
				Category:  entity.Category,
				Title:     entity.Title,
				CreatedAt: entity.CreatedAt,
				UpdatedAt: entity.UpdatedAt,
				DeletedAt: entity.DeletedAt,
			})
		}

		return nil
	}); err != nil {
		return []Response{}, err
	}

	return responses, nil
}

func (s Service) Update(params Params, imageFile ImageFile, req UpdateRequest) (Response, error) {
	req, err := s.validator.ValidateUpdateRequest(req)
	if err != nil {
		return Response{}, err
	}

	response := Response{}

	if err := s.uow.Do(context.Background(), func(ctx context.Context, db database.DB) error {
		repository := NewRepository(db)

		entity, err := repository.Update(ctx, params.ID, Entity{
			Alt:      req.Alt,
			Category: req.Category,
			Title:    req.Title,
		})
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join("public", entity.Src), imageFile.Data, 0644); err != nil {
			return err
		}

		response = Response{
			Id:        entity.Id,
			Src:       entity.Src,
			Alt:       entity.Alt,
			Category:  entity.Category,
			Title:     entity.Title,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
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

		entity, err := repository.Delete(ctx, params.ID)
		if err != nil {
			return err
		}

		if err := os.Remove(filepath.Join("public", entity.Src)); err != nil {
			return err
		}

		response = Response{
			Id:        entity.Id,
			Src:       entity.Src,
			Alt:       entity.Alt,
			Category:  entity.Category,
			Title:     entity.Title,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
		}

		return nil
	}); err != nil {
		return Response{}, err
	}

	return response, nil
}
