package api

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func logErrs(errs ...error) {
	for _, err := range errs {
		if err != nil {
			now := time.Now()

			message := fmt.Sprintf("date:%s time:%s", now.Format("2006-01-02"), now.Format("15:04:05.000"))

			_, file, line, ok := runtime.Caller(2)
			if ok {
				message += fmt.Sprintf(" file:%s line:%d", file, line)
			}

			fmt.Println(message, " error: ", err)
		}
	}
}

func ErrInvalidRequestBody(c *fiber.Ctx, errs ...error) error {
	logErrs(errs...)

	return c.Status(fiber.StatusBadRequest).JSON(ResponseError("Invalid request body"))
}

func ErrInvalidRequestQuery(c *fiber.Ctx, errs ...error) error {
	logErrs(errs...)

	return c.Status(fiber.StatusBadRequest).JSON(ResponseError("Invalid request query"))
}

func ErrInternalServer(c *fiber.Ctx, errs ...error) error {
	logErrs(errs...)

	return c.Status(fiber.StatusInternalServerError).JSON(ResponseError("Internal Server Error"))
}

func ErrNotFound(c *fiber.Ctx, errs ...error) error {
	logErrs(errs...)

	return c.Status(fiber.StatusNotFound).JSON(ResponseError("Not Found"))
}

func ErrInvalidRequestField(c *fiber.Ctx, field, message string, errs ...error) error {
	logErrs(errs...)

	return c.Status(fiber.StatusUnprocessableEntity).JSON(ResponseError(FieldError{
		Field:   field,
		Message: message,
	}))
}
