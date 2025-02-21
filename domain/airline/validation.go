package airline

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

	// Validate name
	req.Name = strings.TrimSpace(req.Name)

	if nameLength := len(req.Name); nameLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMax100Chars)
	}

	// Validate skytrax type
	req.SkytraxType = strings.TrimSpace(req.SkytraxType)

	if skytraxTypeLength := len(req.SkytraxType); skytraxTypeLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "skytrax_type", constant.ErrMustBeFilled)
	} else if skytraxTypeLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "skytrax_type", constant.ErrMax100Chars)
	}

	switch req.SkytraxType {
	case "Full Service", "Low Cost":
		break
	default:
		return Request{}, false, api.ErrInvalidRequestField(
			c, "skytrax_type", "Must be 'Full Service' or 'Low Cost'",
		)
	}

	// Validate skytrax rating
	if req.SkytraxRating < 1 || req.SkytraxRating > 5 {
		return Request{}, false, api.ErrInvalidRequestField(
			c, "skytrax_rating", "Must be between 1 and 5",
		)
	}

	// Validate logo
	if req.Logo.Valid {
		logo := req.Logo.Int64
		if logo < 1 {
			return Request{}, false, api.ErrInvalidRequestField(c, "logo", "Must be greater than 0")
		}
	}

	return req, true, nil
}
