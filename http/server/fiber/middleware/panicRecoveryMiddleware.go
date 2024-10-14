package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PanicRecoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic information
			fmt.Printf("Panic recovered: %v\n", r)

			// Send an error response to the client
			c.Status(http.StatusInternalServerError).SendString("An error occurred")
		}
	}()

	return c.Next()
}
