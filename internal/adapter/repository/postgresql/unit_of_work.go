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

func (factory factoryPostgresql) NewFlightRepository() repository.FlightRepository {
	return NewFlightRepository(factory.db)
}

func (factory factoryPostgresql) NewFlightRouteRepository() repository.FlightRouteRepository {
	return NewFlightRouteRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryRepository() repository.ItineraryRepository {
	return NewItineraryRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryDayRepository() repository.ItineraryDayRepository {
	return NewItineraryDayRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryWidgetRepository() repository.ItineraryWidgetRepository {
	return NewItineraryWidgetRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryWidgetActivityRepository() repository.ItineraryWidgetActivityRepository {
	return NewItineraryWidgetActivityRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryWidgetHotelRepository() repository.ItineraryWidgetHotelRepository {
	return NewItineraryWidgetHotelRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryWidgetInformationRepository() repository.ItineraryWidgetInformationRepository {
	return NewItineraryWidgetInformationRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryWidgetTransportRepository() repository.ItineraryWidgetTransportRepository {
	return NewItineraryWidgetTransportRepository(factory.db)
}

func (factory factoryPostgresql) NewItineraryWidgetRecommendationRepository() repository.ItineraryWidgetRecommendationRepository {
	return NewItineraryWidgetRecommendationRepository(factory.db)
}

func (factory factoryPostgresql) NewLandingHeroContentRepository() repository.LandingHeroContentRepository {
	return NewLandingHeroContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingSectionHeaderRepository() repository.LandingSectionHeaderRepository {
	return NewLandingSectionHeaderRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingPackageItemRepository() repository.LandingPackageItemRepository {
	return NewLandingPackageItemRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingSinglePackageContentRepository() repository.LandingSinglePackageContentRepository {
	return NewLandingSinglePackageContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingPackageDetailRepository() repository.LandingPackageDetailRepository {
	return NewLandingPackageDetailRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingPackageDetailItemRepository() repository.LandingPackageDetailItemRepository {
	return NewLandingPackageDetailItemRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingPackagesContentRepository() repository.LandingPackagesContentRepository {
	return NewLandingPackagesContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingFeaturesContentRepository() repository.LandingFeaturesContentRepository {
	return NewLandingFeaturesContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingFeaturesContentBenefitRepository() repository.LandingFeaturesContentBenefitRepository {
	return NewLandingFeaturesContentBenefitRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingMomentsContentRepository() repository.LandingMomentsContentRepository {
	return NewLandingMomentsContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingMomentsContentImageRepository() repository.LandingMomentsContentImageRepository {
	return NewLandingMomentsContentImageRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingAffiliatesContentRepository() repository.LandingAffiliatesContentRepository {
	return NewLandingAffiliatesContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingAffiliatesContentAffiliateRepository() repository.LandingAffiliatesContentAffiliateRepository {
	return NewLandingAffiliatesContentAffiliateRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingFaqContentRepository() repository.LandingFaqContentRepository {
	return NewLandingFaqContentRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingFaqContentFaqRepository() repository.LandingFaqContentFaqRepository {
	return NewLandingFaqContentFaqRepository(factory.db)

}
func (factory factoryPostgresql) NewLandingMenuRepository() repository.LandingMenuRepository {
	return NewLandingMenuRepository(factory.db)

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
