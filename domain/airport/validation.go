package airport

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

	// Validate city
	req.City = strings.TrimSpace(req.City)

	if cityLength := len(req.City); cityLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "city", constant.ErrMustBeFilled)
	} else if cityLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "city", constant.ErrMax100Chars)
	}

	// Validate name
	req.Name = strings.TrimSpace(req.Name)

	if nameLength := len(req.Name); nameLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMustBeFilled)
	} else if nameLength > 100 {
		return Request{}, false, api.ErrInvalidRequestField(c, "name", constant.ErrMax100Chars)
	}

	// Validate code
	req.Code = strings.ToUpper(strings.TrimSpace(req.Code))

	if codeLength := len(req.Code); codeLength < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "code", constant.ErrMustBeFilled)
	} else if codeLength != 3 {
		return Request{}, false, api.ErrInvalidRequestField(c, "code", constant.ErrInvalidAirportCode)
	}

	for _, char := range req.Code {
		if char < 'A' || char > 'Z' {
			return Request{}, false, api.ErrInvalidRequestField(c, "code", constant.ErrInvalidAirportCode)
		}
	}

	return req, true, nil
}
