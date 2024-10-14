package middleware

import (
	_enum "github.com/thisissoroush/thunder.shared/primitive/enum"
	_primitive "github.com/thisissoroush/thunder.shared/primitive/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func BaseRequestMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	var requestUserId uuid.UUID

	if authHeader != "" {
		requestUserId = uuid.New()
	}

	culture := c.Get("Culture")
	if culture == "" {
		culture = "en"
	}

	baseRequest := _primitive.BaseRequest{
		RequestId:        uuid.New(),
		RequestIPAddress: c.IP(),
		RequestUserId:    requestUserId,
		RequestCulture:   _enum.Language(culture),
	}

	// Store the BaseRequest in the context for later use
	c.Locals("BaseReqeust", &baseRequest)

	// Continue to the next handler
	return c.Next()
}
