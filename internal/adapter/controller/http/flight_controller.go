package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http/schema"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type FlightController struct {
	flightService service.FlightService
}

func NewFlightController(flightService service.FlightService) FlightController {
	return FlightController{
		flightService,
	}
}

func (c FlightController) CreateFlight(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.FlightRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Create flight with service
	dtoResponse, err := c.flightService.CreateFlight(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewFlightResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c FlightController) GetAllFlight(ctx *fiber.Ctx) error {
	// Parse request query
	query := schema.GetAllFlightQuery{}
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get page from query
	page := 1
	if query.Page.Valid {
		page = int(query.Page.Int64)
	}

	// Get per page from query
	perPage := 10
	if query.PerPage.Valid {
		perPage = int(query.PerPage.Int64)
	}

	// Create dto get all request
	dtoGetAllRequest := dto.GetAllFlightRequest{
		Page:    page,
		PerPage: perPage,
	}

	// Get all flight with service
	dtoResponses, err := c.flightService.GetAllFlight(context.Background(), dtoGetAllRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema responses from dto responses
	schemaResponses := schema.NewFlightResponses(dtoResponses)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponses))
}

func (c FlightController) GetFlightById(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.FlightParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get flight by id with service
	dtoResponse, err := c.flightService.GetFlightById(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewFlightResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c FlightController) UpdateFlight(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.FlightParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request body
	schemaRequest := schema.FlightRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Update flight with service
	dtoResponse, err := c.flightService.UpdateFlight(context.Background(), params.Id, dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewFlightResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c FlightController) DeleteFlight(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.FlightParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Delete flight with service
	dtoResponse, err := c.flightService.DeleteFlight(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewFlightResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}
