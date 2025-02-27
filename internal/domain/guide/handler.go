package guide

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/internal/common/api"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) Create(c *fiber.Ctx) error {
	req := CreateRequest{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	response, err := h.service.Create(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.ResponseData(response))
}

func (h Handler) Get(c *fiber.Ctx) error {
	params := Params{}
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	response, err := h.service.Get(params)
	if err != nil {
		return err
	}

	return c.JSON(api.ResponseData(response))
}

func (h Handler) GetAll(c *fiber.Ctx) error {
	query := Query{}
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	response, meta, err := h.service.List(query)
	if err != nil {
		return err
	}

	return c.JSON(api.ResponseData(response, meta))
}

func (h Handler) Update(c *fiber.Ctx) error {
	params := Params{}
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	req := UpdateRequest{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	response, err := h.service.Update(params, req)
	if err != nil {
		return err
	}

	return c.JSON(api.ResponseData(response))
}

func (h Handler) Delete(c *fiber.Ctx) error {
	params := Params{}
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	response, err := h.service.Delete(params)
	if err != nil {
		return err
	}

	return c.JSON(api.ResponseData(response))
}

func (h Handler) Routing(router fiber.Router) {
	router.Post("/", h.Create)
	router.Get("/:id", h.Get)
	router.Get("/", h.GetAll)
	router.Put("/:id", h.Update)
	router.Delete("/:id", h.Delete)
}
