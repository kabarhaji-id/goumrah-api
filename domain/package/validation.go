package pkg

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/constant"
)

func validateRequest(c *fiber.Ctx) (Request, bool, error) {
	req := Request{}
	if err := c.BodyParser(&req); err != nil {
		return Request{}, false, api.ErrInvalidRequestBody(c, err)
	}

	// Validate thumbnail
	if req.Thumbnail.Valid {
		thumbnail := req.Thumbnail.Int64
		if thumbnail < 1 {
			return Request{}, false, api.ErrInvalidRequestField(c, "thumbnail", constant.ErrMin1)
		}
	}

	// Validate name
	req.Name = strings.TrimSpace(req.Name)

	if nameLength := len(req.Name); nameLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMax100Chars)
	}

	// Validate description
	req.Description = strings.TrimSpace(req.Description)

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return Request{}, false, api.ErrInvalidRequestField(c, "description", constant.ErrMax500Chars)
	}

	// Validate category
	if categoryLength := len(req.Category); categoryLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "category", constant.ErrMustBeFilled)
	} else if categoryLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "category", constant.ErrMax100Chars)
	}

	switch req.Category {
	case "Silver", "Gold", "Platinum", "Luxury":
		break
	default:
		return Request{}, false, api.ErrInvalidRequestField(c, "category", constant.ErrInvalidPackageCategory)
	}

	// Validate type
	if typeLength := len(req.Type); typeLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "type", constant.ErrMustBeFilled)
	} else if typeLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "type", constant.ErrMax100Chars)
	}

	switch req.Type {
	case "Reguler", "Plus Wisata":
		break
	default:
		return Request{}, false, api.ErrInvalidRequestField(c, "type", constant.ErrInvalidPackageType)
	}

	return req, true, nil
}
