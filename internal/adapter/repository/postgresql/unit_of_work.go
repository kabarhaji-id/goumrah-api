package postgresql

import (
	"context"

	"github.com/kabarhaji-id/goumrah-api/internal/port/driven/repository"
)

type factoryPostgresql struct {
	db DB
}

func (factory factoryPostgresql) NewImageRepository() repository.ImageRepository {
	return NewImageRepository(factory.db)
}

func (factory factoryPostgresql) NewAirlineRepository() repository.AirlineRepository {
	return NewAirlineRepository(factory.db)
}

func (factory factoryPostgresql) NewEmbarkationRepository() repository.EmbarkationRepository {
	return NewEmbarkationRepository(factory.db)
}

func (factory factoryPostgresql) NewPackageRepository() repository.PackageRepository {
	return NewPackageRepository(factory.db)
}

func (factory factoryPostgresql) NewAddonCategoryRepository() repository.AddonCategoryRepository {
	return NewAddonCategoryRepository(factory.db)
}

func (factory factoryPostgresql) NewGuideRepository() repository.GuideRepository {
	return NewGuideRepository(factory.db)
}

func (factory factoryPostgresql) NewAirportRepository() repository.AirportRepository {
	return NewAirportRepository(factory.db)
}

func (factory factoryPostgresql) NewBusRepository() repository.BusRepository {
	return NewBusRepository(factory.db)
}

func (factory factoryPostgresql) NewPackageSessionRepository() repository.PackageSessionRepository {
	return NewPackageSessionRepository(factory.db)
}

func (factory factoryPostgresql) NewHotelRepository() repository.HotelRepository {
	return NewHotelRepository(factory.db)
}

func (factory factoryPostgresql) NewFacilityRepository() repository.FacilityRepository {
	return NewFacilityRepository(factory.db)
}

func (factory factoryPostgresql) NewAddonRepository() repository.AddonRepository {
	return NewAddonRepository(factory.db)
}

func (factory factoryPostgresql) NewCityTourRepository() repository.CityTourRepository {
	return NewCityTourRepository(factory.db)
}

type unitOfWorkPostgresql struct {
	db DB
}

func NewUnitOfWork(db DB) repository.UnitOfWork {
	unitOfWork := unitOfWorkPostgresql{db}

	return unitOfWork
}

func (uow unitOfWorkPostgresql) Do(
	ctx context.Context,
	fn func(ctx context.Context, factory repository.Factory) error,
) error {
	tx, err := uow.db.Begin(ctx)
	if err != nil {
		return err
	}

	factory := factoryPostgresql{tx}

	if err = fn(ctx, factory); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			err = rollbackErr
		}

		return err
	}

	return tx.Commit(ctx)
}
