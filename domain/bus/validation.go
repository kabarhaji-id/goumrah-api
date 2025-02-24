package bus

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

	// Validate seat
	if req.Seat < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "seat", constant.ErrMin1)
	}

	return req, true, nil
}
