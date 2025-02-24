package pkgsession

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kabarhaji-id/goumrah-api/api"
	"github.com/kabarhaji-id/goumrah-api/constant"
)

func validateRequest(c *fiber.Ctx) (Request, bool, error) {
	req := Request{}
	if err := c.BodyParser(&req); err != nil {
		return Request{}, false, api.ErrInvalidRequestBody(c, err)
	}

	// Validate embarkation
	if req.Embarkation < 1 {
		return Request{}, false, api.ErrInvalidRequestField(c, "embarkation", constant.ErrMin1)
	}

	// Validate departure date
	if req.DepartureDate.IsZero() {
		return Request{}, false, api.ErrInvalidRequestField(c, "departure_date", constant.ErrNotZeroDate)
	}
	if req.DepartureDate.Before(time.Now()) {
		return Request{}, false, api.ErrInvalidRequestField(c, "departure_date", constant.ErrNotBeforeNow)
	}

	return req, true, nil
}
