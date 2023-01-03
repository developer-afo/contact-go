package handler

import (
	"net/http"

	"github.com/afolabiolayinka/contact-go/payload/response"
	"github.com/afolabiolayinka/contact-go/security"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// BaseHandler : handler
type BaseHandler struct {
}

// GetUserID handler
func (baseHandler *BaseHandler) GetUserID(context *fiber.Ctx) (uuid.UUID, error) {
	userID, err := security.ExtractUserID(context.Request())
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

// Index hanlde
func (handler *BaseHandler) Index(c *fiber.Ctx) error {

	var resp response.Response

	var about struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Build   int8   `json:"build"`
		Author  string `json:"author"`
	}

	about.Name = "Contact App API"
	about.Version = "0.1.1"
	about.Build = 1
	about.Author = "Afolabi Olayinka"

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"about": about}

	return c.JSON(resp)
}
