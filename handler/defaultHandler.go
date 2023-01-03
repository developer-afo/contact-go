package handler

import (
	"github.com/gofiber/fiber/v2"
)

// DefaultHandler : handler
type DefaultHandler struct {
	BaseHandler
}

// Index hanlde api status
func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "200", "message": "Success", "data": map[string]interface{}{"name": "Contact App - Go"}})
}
