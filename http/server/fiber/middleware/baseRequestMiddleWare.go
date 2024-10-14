package middleware

import (
	"github.com/gofiber/fiber/v2"
	_enum "github.com/thisissoroush/thunder.shared/primitive/enum"
	_model "github.com/thisissoroush/thunder.shared/primitive/model"
)

func BaseRequestMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	var requestUserId _model.UUID

	if authHeader != "" {
		requestUserId = _model.NewUid()
	}

	culture := c.Get("Culture")
	if culture == "" {
		culture = "en"
	}

	baseRequest := _model.BaseRequest{
		RequestId:        _model.NewUid(),
		RequestIPAddress: c.IP(),
		RequestUserId:    requestUserId,
		RequestCulture:   _enum.Language(culture),
	}

	// Store the BaseRequest in the context for later use
	c.Locals("BaseReqeust", &baseRequest)

	// Continue to the next handler
	return c.Next()
}
