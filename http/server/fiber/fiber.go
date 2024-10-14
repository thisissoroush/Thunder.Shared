package fiber

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	return fiber.New()
}

type Ctx = fiber.Ctx
type Router = fiber.Router
