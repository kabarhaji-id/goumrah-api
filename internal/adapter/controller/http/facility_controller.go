package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http/schema"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type FacilityController struct {
	facilityService service.FacilityService
}

func NewFacilityController(facilityService service.FacilityService) FacilityController {
	return FacilityController{
		facilityService,
	}
}

func (c FacilityController) CreateFacility(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.FacilityRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Create facility with service
	dtoResponse, err := c.facilityService.CreateFacility(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewFacilityResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c FacilityController) GetAllFacility(ctx *fiber.Ctx) error {
	// Parse request query
	query := schema.GetAllFacilityQuery{}
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
	dtoGetAllRequest := dto.GetAllFacilityRequest{
		Page:    page,
		PerPage: perPage,
	}

	// Get all facility with service
	dtoResponses, err := c.facilityService.GetAllFacility(context.Background(), dtoGetAllRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema responses from dto responses
	schemaResponses := schema.NewFacilityResponses(dtoResponses)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponses))
}

func (c FacilityController) GetFacilityById(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.FacilityParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get facility by id with service
	dtoResponse, err := c.facilityService.GetFacilityById(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewFacilityResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c FacilityController) UpdateFacility(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.FacilityParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request body
	schemaRequest := schema.FacilityRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Update facility with service
	dtoResponse, err := c.facilityService.UpdateFacility(context.Background(), params.Id, dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewFacilityResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c FacilityController) DeleteFacility(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.FacilityParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Delete facility with service
	dtoResponse, err := c.facilityService.DeleteFacility(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewFacilityResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}
