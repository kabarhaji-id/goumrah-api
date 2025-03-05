package repository

import "context"

type Factory interface {
	NewImageRepository() ImageRepository
	NewAirlineRepository() AirlineRepository
	NewEmbarkationRepository() EmbarkationRepository
	NewPackageRepository() PackageRepository
	NewAddonCategoryRepository() AddonCategoryRepository
	NewGuideRepository() GuideRepository
	NewAirportRepository() AirportRepository
	NewBusRepository() BusRepository
	NewPackageSessionRepository() PackageSessionRepository
	NewHotelRepository() HotelRepository
	NewFacilityRepository() FacilityRepository
	NewAddonRepository() AddonRepository
	NewCityTourRepository() CityTourRepository
	NewFlightRepository() FlightRepository
	NewFlightRouteRepository() FlightRouteRepository
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(ctx context.Context, factory Factory) error) error
}
