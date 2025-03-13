package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http/schema"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type BusController struct {
	busService service.BusService
}

func NewBusController(busService service.BusService) BusController {
	return BusController{
		busService,
	}
}

func (c BusController) CreateBus(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.BusRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Create bus with service
	dtoResponse, err := c.busService.CreateBus(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewBusResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c BusController) GetAllBus(ctx *fiber.Ctx) error {
	// Parse request query
	query := schema.GetAllBusQuery{}
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
	dtoGetAllRequest := dto.GetAllBusRequest{
		Page:    page,
		PerPage: perPage,
	}

	// Get all bus with service
	dtoResponses, err := c.busService.GetAllBus(context.Background(), dtoGetAllRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema responses from dto responses
	schemaResponses := schema.NewBusResponses(dtoResponses)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponses))
}

func (c BusController) GetBusById(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.BusParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get bus by id with service
	dtoResponse, err := c.busService.GetBusById(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewBusResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c BusController) UpdateBus(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.BusParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request body
	schemaRequest := schema.BusRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Update bus with service
	dtoResponse, err := c.busService.UpdateBus(context.Background(), params.Id, dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewBusResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c BusController) DeleteBus(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.BusParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Delete bus with service
	dtoResponse, err := c.busService.DeleteBus(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewBusResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}
