package app

import (
	"net/http"

	dtos "github.com/afolabiolayinka/contact-go/database/dto/app"
	services "github.com/afolabiolayinka/contact-go/database/service/app"
	"github.com/afolabiolayinka/contact-go/handler"
	"github.com/afolabiolayinka/contact-go/payload/response"
	validators "github.com/afolabiolayinka/contact-go/validator/app"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ContactHandlerInterface interface {
	IndexHandle(context *fiber.Ctx) error
	CreateHandle(context *fiber.Ctx) error
	ReadHandle(context *fiber.Ctx) error
	UpdateHandle(context *fiber.Ctx) error
	DeleteHandle(context *fiber.Ctx) error
}

type contactHandler struct {
	handler.BaseHandler
	contactService   services.ContactServiceInterface
	contactValidator validators.ContactValidator
}

func NewContactHandler(
	contactService services.ContactServiceInterface,
) ContactHandlerInterface {
	return &contactHandler{
		contactService: contactService,
	}
}

// IndexHandle : handle
func (handler *contactHandler) IndexHandle(context *fiber.Ctx) (err error) {

	var resp response.Response

	pageable := handler.GeneratePageable(context)

	contacts, pagination, queryError := handler.contactService.ReadAll(pageable)

	if queryError != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = queryError.Error()

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"contacts": contacts, "pagination": pagination}

	return context.Status(http.StatusOK).JSON(resp)
}

// CreateHandle : handle
func (handler *contactHandler) CreateHandle(context *fiber.Ctx) (err error) {

	var resp response.Response

	userID, _ := handler.GetUserID(context)

	contactDto := new(dtos.ContactDTO)

	if err := context.BodyParser(contactDto); err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = err.Error()

		return context.Status(http.StatusBadRequest).JSON(resp)
	}

	contactDto.CreatedByID = userID

	vEs, err := handler.contactValidator.Validate(*contactDto)

	if err != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = "Validation error"
		resp.Data = vEs

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	if _, err := handler.contactService.Create(*contactDto); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Status = http.StatusCreated
	resp.Message = http.StatusText(http.StatusCreated)
	resp.Data = map[string]interface{}{"contact": contactDto}

	return context.Status(http.StatusOK).JSON(resp)

}

// ReadHandle : handle
func (handler *contactHandler) ReadHandle(context *fiber.Ctx) (err error) {

	var resp response.Response

	uid, err := uuid.Parse(context.Params("uid"))
	if err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Exception Error: " + err.Error()
		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	contactDto, err := handler.contactService.Read(uid)

	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Record not found"
		return context.Status(http.StatusNotFound).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"contact": contactDto}

	return context.Status(http.StatusOK).JSON(resp)
}

// UpdateHandle : handle
func (handler *contactHandler) UpdateHandle(context *fiber.Ctx) (err error) {

	var resp response.Response
	userID, _ := handler.GetUserID(context)

	uid, err := uuid.Parse(context.Params("uid"))
	if err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Exception Error: " + err.Error()
		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	contactDto, err := handler.contactService.Read(uid)

	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Record not found"
		return context.Status(http.StatusNotFound).JSON(resp)
	}

	if err := context.BodyParser(&contactDto); err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = err.Error()

		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	// assign ID must be place before validation
	contactDto.UUID = uid
	contactDto.UpdatedByID = userID

	//validate
	vEs, err := handler.contactValidator.Validate(contactDto)

	if err != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = "Validation error"
		resp.Data = vEs

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	if _, err := handler.contactService.Update(contactDto); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"contact": contactDto}

	return context.Status(http.StatusOK).JSON(resp)
}

// DeleteHandle : handle
func (handler *contactHandler) DeleteHandle(context *fiber.Ctx) (err error) {

	var resp response.Response
	userID, _ := handler.GetUserID(context)

	uid, err := uuid.Parse(context.Params("uid"))
	if err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Exception Error: " + err.Error()
		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	contactDto, err := handler.contactService.Read(uid)

	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Record not found"
		return context.Status(http.StatusNotFound).JSON(resp)
	}

	if err := handler.contactService.Delete(contactDto.UUID, userID); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)

	return context.Status(http.StatusOK).JSON(resp)
}
