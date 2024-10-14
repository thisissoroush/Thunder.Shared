package middleware

import (
	_exc "github.com/thisissoroush/thunder.shared/primitive/model/exception"

	"github.com/gofiber/fiber/v2"
)

func ResponseWrapper(c *fiber.Ctx) error {
	// Call the next middleware/handler
	res := c.Next()

	// Check if there was an error
	if res != nil {
		if err, ok := res.(*_exc.ThunderException); ok {

			c.Status(err.Code).SendString(err.Message)
		} else {

			c.Status(500).SendString(err.Message)
		}
	}

	return nil
}
