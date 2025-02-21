package image

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/constant"
)

func validateRequest(c *fiber.Ctx) (Request, bool, error) {
	req := Request{}
	if err := c.BodyParser(&req); err != nil {
		return Request{}, false, api.ErrInvalidRequestBody(c, err)
	}

	// Validate image
	imageFileHeader, err := c.FormFile("image")
	if err != nil {
		return Request{}, false, api.ErrInvalidRequestField(c, "image", constant.ErrMustBeFilled, err)
	}

	if !strings.HasPrefix(imageFileHeader.Header.Get("Content-Type"), "image/") {
		return Request{}, false, api.ErrInvalidRequestField(c, "image", "Must be an image")
	}

	req.Image = imageFileHeader

	// Validate alt
	req.Alt = strings.TrimSpace(req.Alt)

	if altLength := len(req.Alt); altLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "alt", constant.ErrMustBeFilled)
	} else if altLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "alt", constant.ErrMax100Chars)
	}

	// Validate category
	if req.Category.Valid {
		category := strings.TrimSpace(req.Category.String)

		if categoryLength := len(category); categoryLength < 1 {
			return Request{}, false, api.ErrInvalidRequestField(c, "category", constant.ErrMustBeFilled)
		} else if categoryLength > 100 {
			return Request{}, false, api.ErrInvalidRequestField(c, "category", constant.ErrMax100Chars)
		}

		req.Category = null.StringFrom(category)
	}

	// Validate title
	req.Title = strings.TrimSpace(req.Title)

	if titleLength := len(req.Title); titleLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "title", constant.ErrMustBeFilled)
	} else if titleLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "title", constant.ErrMax100Chars)
	}

	return req, true, nil
}
