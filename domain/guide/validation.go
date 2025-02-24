package guide

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

	// Validate avatar
	if req.Avatar.Valid {
		avatar := req.Avatar.Int64
		if avatar < 1 {
			return Request{}, false, api.ErrInvalidRequestField(c, "logo", constant.ErrMin1)
		}
	}

	// Validate name
	req.Name = strings.TrimSpace(req.Name)

	if nameLength := len(req.Name); nameLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMax100Chars)
	}

	// Validate type
	req.Type = strings.TrimSpace(req.Type)

	if typeLength := len(req.Type); typeLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "type", constant.ErrMustBeFilled)
	} else if typeLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "type", constant.ErrMax100Chars)
	}

	switch req.Type {
	case "Perjalanan", "Ibadah":
		break
	default:
		return Request{}, false, api.ErrInvalidRequestField(
			c, "type", "Must be 'Perjalanan' or 'Ibadah'",
		)
	}

	// Validate description
	req.Description = strings.TrimSpace(req.Description)

	if descriptionLength := len(req.Description); descriptionLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "description", constant.ErrMustBeFilled)
	} else if descriptionLength > 500 {
		return Request{}, false, api.ErrInvalidRequestField(c, "description", constant.ErrMax500Chars)
	}

	return req, true, nil
}
