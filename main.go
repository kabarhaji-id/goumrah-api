package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/config"
	"github.com/kabarhaji-id/goumrah-api/database"
	"github.com/kabarhaji-id/goumrah-api/domain/airline"
	"github.com/kabarhaji-id/goumrah-api/domain/embarkation"
	"github.com/kabarhaji-id/goumrah-api/domain/image"
	pkg "github.com/kabarhaji-id/goumrah-api/domain/package"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.InitPool(cfg); err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()
	app.Route("/images", image.Routing)
	app.Route("/airlines", airline.Routing)
	app.Route("/embarkations", embarkation.Routing)
	app.Route("/packages", pkg.Routing)

	if err := app.Listen(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
