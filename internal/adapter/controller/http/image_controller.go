package http

import (
	"context"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/internal/adapter/controller/http/schema"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/dto"
	"github.com/kabarhaji-id/goumrah-api/internal/port/driving/service"
)

type ImageController struct {
	imageService service.ImageService
}

func NewImageController(imageService service.ImageService) ImageController {
	return ImageController{
		imageService,
	}
}

func (c ImageController) CreateImage(ctx *fiber.Ctx) error {
	// Parse request body
	schemaRequest := schema.ImageRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get image form file
	imageFileHeader, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Open image form file
	imageFile, err := imageFileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}
	defer imageFile.Close()

	// Read all image form file content
	imageFileData, err := io.ReadAll(imageFile)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()
	dtoRequest.FileData = imageFileData
	dtoRequest.FileType = imageFileHeader.Header.Get("Content-Type")

	// Create image with service
	dtoResponse, err := c.imageService.CreateImage(context.Background(), dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewImageResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c ImageController) GetAllImage(ctx *fiber.Ctx) error {
	// Parse request query
	query := schema.GetAllImageQuery{}
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
	dtoGetAllRequest := dto.GetAllImageRequest{
		Page:    page,
		PerPage: perPage,
	}

	// Get all image with service
	dtoResponses, err := c.imageService.GetAllImage(context.Background(), dtoGetAllRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema responses from dto responses
	schemaResponses := schema.NewImageResponses(dtoResponses)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponses))
}

func (c ImageController) GetImageById(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.ImageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get image by id with service
	dtoResponse, err := c.imageService.GetImageById(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Create schema response from dto response
	schemaResponse := schema.NewImageResponse(dtoResponse)

	return ctx.Status(fiber.StatusCreated).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c ImageController) UpdateImage(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.ImageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Parse request body
	schemaRequest := schema.ImageRequest{}
	if err := ctx.BodyParser(&schemaRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Get image form file
	imageFileHeader, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Open image form file
	imageFile, err := imageFileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}
	defer imageFile.Close()

	// Read all image form file content
	imageFileData, err := io.ReadAll(imageFile)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Create dto request
	dtoRequest := schemaRequest.ToDtoRequest()
	dtoRequest.FileData = imageFileData
	dtoRequest.FileType = imageFileHeader.Header.Get("Content-Type")

	// Update image with service
	dtoResponse, err := c.imageService.UpdateImage(context.Background(), params.Id, dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewImageResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}

func (c ImageController) DeleteImage(ctx *fiber.Ctx) error {
	// Parse request path params
	params := schema.ImageParams{}
	if err := ctx.ParamsParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schema.NewErrorResponse(err))
	}

	// Delete image with service
	dtoResponse, err := c.imageService.DeleteImage(context.Background(), params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schema.NewErrorResponse(err))
	}

	// Update schema response from dto response
	schemaResponse := schema.NewImageResponse(dtoResponse)

	return ctx.Status(fiber.StatusOK).JSON(schema.NewSuccessResponse(schemaResponse))
}
