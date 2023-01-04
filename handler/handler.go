package handler

import (
	"net/http"
	"strconv"

	"github.com/afolabiolayinka/contact-go/database/repository"
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

// GeneratePageable
func (h *BaseHandler) GeneratePageable(context *fiber.Ctx) (pageable repository.Pageable) {

	pageable.Page = 1
	pageable.Size = 20
	pageable.SortBy = "created_at"
	pageable.SortDirection = "asc"
	pageable.Search = ""

	size, err := strconv.Atoi(context.Query("size", "0"))
	if (size > 0) && err == nil {
		pageable.Size = size
	}

	page, err := strconv.Atoi(context.Query("page", "1"))
	if (page > 0) && err == nil {
		pageable.Page = page
	}

	orderBy := context.Query("sort_by", "")
	if orderBy != "" {
		pageable.SortBy = orderBy
	}

	sortDir := context.Query("sort_dir", "")
	if sortDir != "" {
		pageable.SortBy = sortDir
	}

	search := context.Query("search", "")
	if search != "" {
		pageable.Search = search
	}

	return pageable
}
