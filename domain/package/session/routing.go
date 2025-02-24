package pkgsession

import "github.com/gofiber/fiber/v2"

func Routing(router fiber.Router) {
	router.Get("/", GetAllHandler)
	router.Get("/:id", GetOneHandler)
	router.Put("/:id", UpdateHandler)
	router.Delete("/:id", DeleteHandler)
}
