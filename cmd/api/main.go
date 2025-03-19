package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/config"
	httpcontroller "github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http"
	postgresqlrepository "github.com/kabarhaji-id/goumrah-api/internal/adapter/repository/postgresql"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/mapper"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/service"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/validator"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := postgresqlrepository.NewDatabase(cfg.PostgresDSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	unitOfWork := postgresqlrepository.NewUnitOfWork(db)

	imageRepository := postgresqlrepository.NewImageRepository(db)
	airlineRepository := postgresqlrepository.NewAirlineRepository(db)
	embarkationRepository := postgresqlrepository.NewEmbarkationRepository(db)
	packageRepository := postgresqlrepository.NewPackageRepository(db)
	addonCategoryRepository := postgresqlrepository.NewAddonCategoryRepository(db)
	guideRepository := postgresqlrepository.NewGuideRepository(db)
	airportRepository := postgresqlrepository.NewAirportRepository(db)
	busRepository := postgresqlrepository.NewBusRepository(db)
	packageSessionRepository := postgresqlrepository.NewPackageSessionRepository(db)
	hotelRepository := postgresqlrepository.NewHotelRepository(db)
	facilityRepository := postgresqlrepository.NewFacilityRepository(db)
	addonRepository := postgresqlrepository.NewAddonRepository(db)
	cityTourRepository := postgresqlrepository.NewCityTourRepository(db)
	flightRepository := postgresqlrepository.NewFlightRepository(db)
	flightRouteRepository := postgresqlrepository.NewFlightRouteRepository(db)
	itineraryRepository := postgresqlrepository.NewItineraryRepository(db)
	itineraryDayRepository := postgresqlrepository.NewItineraryDayRepository(db)
	itineraryWidgetRepository := postgresqlrepository.NewItineraryWidgetRepository(db)
	itineraryWidgetActivityRepository := postgresqlrepository.NewItineraryWidgetActivityRepository(db)
	itineraryWidgetHotelRepository := postgresqlrepository.NewItineraryWidgetHotelRepository(db)
	itineraryWidgetInformationRepository := postgresqlrepository.NewItineraryWidgetInformationRepository(db)
	itineraryWidgetTransportRepository := postgresqlrepository.NewItineraryWidgetTransportRepository(db)
	itineraryWidgetRecommendationRepository := postgresqlrepository.NewItineraryWidgetRecommendationRepository(db)
	userRepository := postgresqlrepository.NewUserRepository(db)

	imageValidator := validator.NewImageValidator()
	airlineValidator := validator.NewAirlineValidator()
	embarkationValidator := validator.NewEmbarkationValidator()
	packageValidator := validator.NewPackageValidator()
	addonCategoryValidator := validator.NewAddonCategoryValidator()
	guideValidator := validator.NewGuideValidator()
	airportValidator := validator.NewAirportValidator()
	busValidator := validator.NewBusValidator()
	packageSessionValidator := validator.NewPackageSessionValidator()
	hotelValidator := validator.NewHotelValidator()
	facilityValidator := validator.NewFacilityValidator()
	addonValidator := validator.NewAddonValidator()
	cityTourValidator := validator.NewCityTourValidator()
	flightValidator := validator.NewFlightValidator()
	userValidator := validator.NewUserValidator()

	imageMapper := mapper.NewImageMapper()
	airlineMapper := mapper.NewAirlineMapper(imageMapper)
	embarkationMapper := mapper.NewEmbarkationMapper()
	packageMapper := mapper.NewPackageMapper(imageMapper)
	addonCategoryMapper := mapper.NewAddonCategoryMapper()
	guideMapper := mapper.NewGuideMapper(imageMapper)
	airportMapper := mapper.NewAirportMapper()
	busMapper := mapper.NewBusMapper()
	flightMapper := mapper.NewFlightMapper(airlineMapper, airportMapper)
	hotelMapper := mapper.NewHotelMapper(imageMapper)
	itineraryWidgetMapper := mapper.NewItineraryWidgetMapper(imageMapper, hotelMapper)
	itineraryDayMapper := mapper.NewItineraryDayMapper(imageMapper, itineraryWidgetMapper)
	itineraryMapper := mapper.NewItineraryMapper(imageMapper, itineraryDayMapper)
	packageSessionMapper := mapper.NewPackageSessionMapper(embarkationMapper, guideMapper, flightMapper, busMapper, itineraryMapper)
	facilityMapper := mapper.NewFacilityMapper()
	addonMapper := mapper.NewAddonMapper(addonCategoryMapper)
	cityTourMapper := mapper.NewCityTourMapper()
	userMapper := mapper.NewUserMapper()

	imageService := service.NewImageService(imageRepository, imageValidator, imageMapper, unitOfWork)
	airlineService := service.NewAirlineService(airlineRepository, airlineValidator, airlineMapper, imageRepository)
	embarkationService := service.NewEmbarkationService(embarkationRepository, embarkationValidator, embarkationMapper)
	packageService := service.NewPackageService(packageRepository, packageValidator, packageMapper, imageRepository, unitOfWork)
	addonCategoryService := service.NewAddonCategoryService(addonCategoryRepository, addonCategoryValidator, addonCategoryMapper)
	guideService := service.NewGuideService(guideRepository, guideValidator, guideMapper, imageRepository)
	airportService := service.NewAirportService(airportRepository, airportValidator, airportMapper)
	busService := service.NewBusService(busRepository, busValidator, busMapper)
	packageSessionService := service.NewPackageSessionService(
		packageSessionRepository, packageSessionValidator, packageSessionMapper,
		embarkationRepository,
		imageRepository,
		flightRouteRepository,
		flightRepository,
		airlineRepository,
		airportRepository,
		busRepository,
		hotelRepository,
		itineraryRepository,
		itineraryMapper,
		itineraryDayRepository,
		itineraryDayMapper,
		itineraryWidgetRepository,
		itineraryWidgetActivityRepository,
		itineraryWidgetHotelRepository,
		itineraryWidgetInformationRepository,
		itineraryWidgetTransportRepository,
		itineraryWidgetRecommendationRepository,
		itineraryWidgetMapper,
		unitOfWork,
	)
	hotelService := service.NewHotelService(hotelRepository, hotelValidator, hotelMapper, unitOfWork)
	facilityService := service.NewFacilityService(facilityRepository, facilityValidator, facilityMapper)
	addonService := service.NewAddonService(addonRepository, addonValidator, addonMapper, addonCategoryRepository)
	cityTourService := service.NewCityTourService(cityTourRepository, cityTourValidator, cityTourMapper)
	flightService := service.NewFlightService(flightRepository, flightValidator, flightMapper, imageRepository, airlineRepository, airportRepository)
	userService := service.NewUserService(userRepository, userValidator, userMapper)

	imageController := httpcontroller.NewImageController(imageService)
	airlineController := httpcontroller.NewAirlineController(airlineService)
	embarkationController := httpcontroller.NewEmbarkationController(embarkationService)
	packageController := httpcontroller.NewPackageController(packageService, packageSessionService)
	addonCategoryController := httpcontroller.NewAddonCategoryController(addonCategoryService)
	guideController := httpcontroller.NewGuideController(guideService)
	airportController := httpcontroller.NewAirportController(airportService)
	busController := httpcontroller.NewBusController(busService)
	packageSessionController := httpcontroller.NewPackageSessionController(packageSessionService)
	hotelController := httpcontroller.NewHotelController(hotelService)
	facilityController := httpcontroller.NewFacilityController(facilityService)
	addonController := httpcontroller.NewAddonController(addonService)
	cityTourController := httpcontroller.NewCityTourController(cityTourService)
	flightController := httpcontroller.NewFlightController(flightService)
	userController := httpcontroller.NewUserController(userService)

	app := fiber.New()
	app.Route("/images", func(router fiber.Router) {
		router.Post("", imageController.CreateImage)
		router.Get("", imageController.GetAllImage)
		router.Get("/:id", imageController.GetImageById)
		router.Put("/:id", imageController.UpdateImage)
		router.Delete("/:id", imageController.DeleteImage)
	})
	app.Route("/airlines", func(router fiber.Router) {
		router.Post("", airlineController.CreateAirline)
		router.Get("", airlineController.GetAllAirline)
		router.Get("/:id", airlineController.GetAirlineById)
		router.Put("/:id", airlineController.UpdateAirline)
		router.Delete("/:id", airlineController.DeleteAirline)
	})
	app.Route("/embarkations", func(router fiber.Router) {
		router.Post("", embarkationController.CreateEmbarkation)
		router.Get("", embarkationController.GetAllEmbarkation)
		router.Get("/:id", embarkationController.GetEmbarkationById)
		router.Put("/:id", embarkationController.UpdateEmbarkation)
		router.Delete("/:id", embarkationController.DeleteEmbarkation)
	})
	app.Route("/packages", func(router fiber.Router) {
		router.Post("", packageController.CreatePackage)
		router.Get("", packageController.GetAllPackage)
		router.Get("/:id", packageController.GetPackageById)
		router.Put("/:id", packageController.UpdatePackage)
		router.Delete("/:id", packageController.DeletePackage)
		router.Post("/:id/sessions", packageController.CreatePackageSession)
		router.Get("/:id/sessions", packageController.GetAllPackageSession)
	})
	app.Route("/addon-categories", func(router fiber.Router) {
		router.Post("", addonCategoryController.CreateAddonCategory)
		router.Get("", addonCategoryController.GetAllAddonCategory)
		router.Get("/:id", addonCategoryController.GetAddonCategoryById)
		router.Put("/:id", addonCategoryController.UpdateAddonCategory)
		router.Delete("/:id", addonCategoryController.DeleteAddonCategory)
	})
	app.Route("/guides", func(router fiber.Router) {
		router.Post("", guideController.CreateGuide)
		router.Get("", guideController.GetAllGuide)
		router.Get("/:id", guideController.GetGuideById)
		router.Put("/:id", guideController.UpdateGuide)
		router.Delete("/:id", guideController.DeleteGuide)
	})
	app.Route("/airports", func(router fiber.Router) {
		router.Post("", airportController.CreateAirport)
		router.Get("", airportController.GetAllAirport)
		router.Get("/:id", airportController.GetAirportById)
		router.Put("/:id", airportController.UpdateAirport)
		router.Delete("/:id", airportController.DeleteAirport)
	})
	app.Route("/buses", func(router fiber.Router) {
		router.Post("", busController.CreateBus)
		router.Get("", busController.GetAllBus)
		router.Get("/:id", busController.GetBusById)
		router.Put("/:id", busController.UpdateBus)
		router.Delete("/:id", busController.DeleteBus)
	})
	app.Route("/package-sessions", func(router fiber.Router) {
		router.Get("", packageSessionController.GetAllPackageSession)
		router.Get("/:id", packageSessionController.GetPackageSessionById)
		router.Put("/:id", packageSessionController.UpdatePackageSession)
		router.Delete("/:id", packageSessionController.DeletePackageSession)
	})
	app.Route("/hotels", func(router fiber.Router) {
		router.Post("", hotelController.CreateHotel)
		router.Get("", hotelController.GetAllHotel)
		router.Get("/:id", hotelController.GetHotelById)
		router.Put("/:id", hotelController.UpdateHotel)
		router.Delete("/:id", hotelController.DeleteHotel)
	})
	app.Route("/facilities", func(router fiber.Router) {
		router.Post("", facilityController.CreateFacility)
		router.Get("", facilityController.GetAllFacility)
		router.Get("/:id", facilityController.GetFacilityById)
		router.Put("/:id", facilityController.UpdateFacility)
		router.Delete("/:id", facilityController.DeleteFacility)
	})
	app.Route("/addons", func(router fiber.Router) {
		router.Post("", addonController.CreateAddon)
		router.Get("", addonController.GetAllAddon)
		router.Get("/:id", addonController.GetAddonById)
		router.Put("/:id", addonController.UpdateAddon)
		router.Delete("/:id", addonController.DeleteAddon)
	})
	app.Route("/city-tours", func(router fiber.Router) {
		router.Post("", cityTourController.CreateCityTour)
		router.Get("", cityTourController.GetAllCityTour)
		router.Get("/:id", cityTourController.GetCityTourById)
		router.Put("/:id", cityTourController.UpdateCityTour)
		router.Delete("/:id", cityTourController.DeleteCityTour)
	})
	app.Route("/flights", func(router fiber.Router) {
		router.Post("", flightController.CreateFlight)
		router.Get("", flightController.GetAllFlight)
		router.Get("/:id", flightController.GetFlightById)
		router.Put("/:id", flightController.UpdateFlight)
		router.Delete("/:id", flightController.DeleteFlight)
	})
	app.Route("/users", func(router fiber.Router) {
		router.Post("", userController.CreateUser)
		router.Get("", userController.GetAllUser)
		router.Get("/:id", userController.GetUserById)
		router.Put("/:id", userController.UpdateUser)
		router.Delete("/:id", userController.DeleteUser)
	})

	if err := app.Listen(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
