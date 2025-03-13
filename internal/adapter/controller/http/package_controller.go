package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http/schema"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type PackageController struct {
	packageService        service.PackageService
	packageSessionService service.PackageSessionService
}

func NewPackageController(packageService service.PackageService, packageSessionService service.PackageSessionService) PackageController {
	return PackageController{
		packageService, packageSessionService,
	}
}

func (c PackageController) CreatePackage(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.PackageRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Create package with service
	dtoResponse, err := c.packageService.CreatePackage(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewPackageResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c PackageController) GetAllPackage(ctx *fiber.Ctx) error {
	// Parse request query
	query := schema.GetAllPackageQuery{}
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
	dtoGetAllRequest := dto.GetAllPackageRequest{
		Page:    page,
		PerPage: perPage,
	}

	// Get all package with service
	dtoResponses, err := c.packageService.GetAllPackage(context.Background(), dtoGetAllRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema responses from dto responses
	schemaResponses := schema.NewPackageListResponses(dtoResponses)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponses))
}

func (c PackageController) GetPackageById(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.PackageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get package by id with service
	dtoResponse, err := c.packageService.GetPackageById(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewPackageResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c PackageController) UpdatePackage(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.PackageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request body
	schemaRequest := schema.PackageRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()

	// Update package with service
	dtoResponse, err := c.packageService.UpdatePackage(context.Background(), params.Id, dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewPackageResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c PackageController) DeletePackage(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.PackageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Delete package with service
	dtoResponse, err := c.packageService.DeletePackage(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewPackageResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c PackageController) CreatePackageSession(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.PackageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request body
	schemaRequest := schema.PackageSessionRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()
	dtoRequest.Package = params.Id

	// Create package session with service
	dtoResponse, err := c.packageSessionService.CreatePackageSession(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewPackageSessionResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c PackageController) GetAllPackageSession(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.PackageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request query
	query := schema.GetAllPackageQuery{}
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
	dtoGetAllRequest := dto.GetAllPackageSessionRequest{
		Package: null.IntFrom(params.Id),
		Page:    page,
		PerPage: perPage,
	}

	// Get all package session with service
	dtoResponses, err := c.packageSessionService.GetAllPackageSession(context.Background(), dtoGetAllRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema responses from dto response
	schemaResponses := schema.NewPackageSessionListResponses(dtoResponses)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponses))
}
