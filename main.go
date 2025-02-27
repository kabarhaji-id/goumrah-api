package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/config"
	"github.com/kabarhaji-id/goumrah-api/internal/common/database"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/addon_category"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/airline"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/airport"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/bus"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/embarkation"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/facility"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/guide"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/hotel"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/image"
	pkg "github.com/kabarhaji-id/goumrah-api/internal/domain/package"
	"github.com/kabarhaji-id/goumrah-api/internal/domain/package_session"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewPostgres(cfg.PostgresDSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()

	uow := database.NewUnitOfWork(db)

	imageValidator := image.NewValidator()
	airlineValidator := airline.NewValidator()
	embarkationValidator := embarkation.NewValidator()
	packageValidator := pkg.NewValidator()
	addonCategoryValidator := addon_category.NewValidator()
	guideValidator := guide.NewValidator()
	airportValidator := airport.NewValidator()
	busValidator := bus.NewValidator()
	packageSessionValidator := package_session.NewValidator()
	hotelValidator := hotel.NewValidator()
	facilityValidator := facility.NewValidator()

	imageService := image.NewService(imageValidator, uow)
	airlineService := airline.NewService(airlineValidator, uow)
	embarkationService := embarkation.NewService(embarkationValidator, uow)
	packageService := pkg.NewService(packageValidator, packageSessionValidator, uow)
	addonCategoryService := addon_category.NewService(addonCategoryValidator, uow)
	guideService := guide.NewService(guideValidator, uow)
	airportService := airport.NewService(airportValidator, uow)
	busService := bus.NewService(busValidator, uow)
	packagesSessionService := package_session.NewService(packageSessionValidator, uow)
	hotelService := hotel.NewService(hotelValidator, uow)
	facilityService := facility.NewService(facilityValidator, uow)

	imageHandler := image.NewHandler(imageService)
	airlineHandler := airline.NewHandler(airlineService)
	embarkationHandler := embarkation.NewHandler(embarkationService)
	packageHandler := pkg.NewHandler(packageService)
	addonCategoryHandler := addon_category.NewHandler(addonCategoryService)
	guideHandler := guide.NewHandler(guideService)
	airportHandler := airport.NewHandler(airportService)
	busHandler := bus.NewHandler(busService)
	packageSessionHandler := package_session.NewHandler(packagesSessionService)
	hotelHandler := hotel.NewHandler(hotelService)
	facilityHandler := facility.NewHandler(facilityService)

	app := fiber.New()
	app.Route("/images", imageHandler.Routing)
	app.Route("/airlines", airlineHandler.Routing)
	app.Route("/embarkations", embarkationHandler.Routing)
	app.Route("/packages", packageHandler.Routing)
	app.Route("/addon-categories", addonCategoryHandler.Routing)
	app.Route("/guides", guideHandler.Routing)
	app.Route("/airports", airportHandler.Routing)
	app.Route("/buses", busHandler.Routing)
	app.Route("/package-sessions", packageSessionHandler.Routing)
	app.Route("/hotels", hotelHandler.Routing)
	app.Route("/facilities", facilityHandler.Routing)

	if err := app.Listen(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
