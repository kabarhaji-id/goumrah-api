package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http/schema"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type LandingController struct {
	landingService service.LandingService
}

func NewLandingController(landingService service.LandingService) LandingController {
	return LandingController{
		landingService,
	}
}

func (c LandingController) CreateLanding(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.LandingRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Create landing with service
	dtoResponse, err := c.landingService.CreateLanding(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewLandingResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c LandingController) GetLanding(ctx *fiber.Ctx) error {
	// Get landing by id with service
	dtoResponse, err := c.landingService.GetLanding(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewLandingResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c LandingController) UpdateLanding(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.LandingRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Update landing with service
	dtoResponse, err := c.landingService.UpdateLanding(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewLandingResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}
