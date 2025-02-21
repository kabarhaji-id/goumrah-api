package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/config"
	"github.com/kabarhaji-id/goumrah-api/database"
	"github.com/kabarhaji-id/goumrah-api/domain/image"
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

	if err := app.Listen(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
