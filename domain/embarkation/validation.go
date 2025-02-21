package embarkation

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

	// Validate latitude
	if req.Latitude < -90 || req.Latitude > 90 {
		return Request{}, false, api.ErrInvalidRequestField(
			c, "latitude", "Must be between -90 and 90",
		)
	}

	// Validate longitude
	if req.Longitude < -180 || req.Longitude > 180 {
		return Request{}, false, api.ErrInvalidRequestField(
			c, "longitude", "Must be between -180 and 180",
		)
	}

	return req, true, nil
}
