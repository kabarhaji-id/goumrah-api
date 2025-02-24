package pkg

import (
	"github.com/gofiber/fiber/v2"
	pkgsession "github.com/kabarhaji-id/goumrah-api/domain/package/session"
)

func Routing(router fiber.Router) {
	router.Post("/", CreateHandler)
	router.Get("/", GetAllHandler)
	router.Get("/:id", GetOneHandler)
	router.Put("/:id", UpdateHandler)
	router.Delete("/:id", DeleteHandler)

	router.Post("/:package_id/sessions", pkgsession.CreateHandler)
	router.Get("/:package_id/sessions", pkgsession.GetAllByPackageHandler)
}
