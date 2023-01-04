package auth

import (
	"net/http"

	dtos "github.com/afolabiolayinka/contact-go/database/dto/auth"
	services "github.com/afolabiolayinka/contact-go/database/service/auth"
	"github.com/afolabiolayinka/contact-go/handler"
	"github.com/afolabiolayinka/contact-go/payload/response"
	validators "github.com/afolabiolayinka/contact-go/validator/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	IndexHandle(context *fiber.Ctx) error
	CreateHandle(context *fiber.Ctx) error
	ReadHandle(context *fiber.Ctx) error
	UpdateHandle(context *fiber.Ctx) error
	DeleteHandle(context *fiber.Ctx) error
}

type userHandler struct {
	handler.BaseHandler
	userService   services.UserServiceInterface
	userValidator validators.UserValidator
}

func NewUserHandler(
	userService services.UserServiceInterface,
) UserHandlerInterface {
	return &userHandler{
		userService: userService,
	}
}

// IndexHandle : handle
func (handler *userHandler) IndexHandle(context *fiber.Ctx) (err error) {

	var resp response.Response

	pageable := handler.GeneratePageable(context)

	users, pagination, queryError := handler.userService.ReadAll(pageable)

	if queryError != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = queryError.Error()

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"users": users, "pagination": pagination}

	return context.Status(http.StatusOK).JSON(resp)
}

// CreateHandle : handle
func (handler *userHandler) CreateHandle(context *fiber.Ctx) (err error) {

	var resp response.Response

	userID, _ := handler.GetUserID(context)

	userDto := new(dtos.UserDTO)

	if err := context.BodyParser(userDto); err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = err.Error()

		return context.Status(http.StatusBadRequest).JSON(resp)
	}

	userDto.CreatedByID = userID

	vEs, err := handler.userValidator.Validate(*userDto)

	if err != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = "Validation error"
		resp.Data = vEs

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	if _, err := handler.userService.Create(*userDto); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Status = http.StatusCreated
	resp.Message = http.StatusText(http.StatusCreated)
	resp.Data = map[string]interface{}{"user": userDto}

	return context.Status(http.StatusOK).JSON(resp)

}

// ReadHandle : handle
func (handler *userHandler) ReadHandle(context *fiber.Ctx) (err error) {

	var resp response.Response

	uid, err := uuid.Parse(context.Params("uid"))
	if err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Exception Error: " + err.Error()
		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	userDto, err := handler.userService.Read(uid)

	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Record not found"
		return context.Status(http.StatusNotFound).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"user": userDto}

	return context.Status(http.StatusOK).JSON(resp)
}

// UpdateHandle : handle
func (handler *userHandler) UpdateHandle(context *fiber.Ctx) (err error) {

	var resp response.Response
	userID, _ := handler.GetUserID(context)

	uid, err := uuid.Parse(context.Params("uid"))
	if err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Exception Error: " + err.Error()
		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	userDto, err := handler.userService.Read(uid)

	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Record not found"
		return context.Status(http.StatusNotFound).JSON(resp)
	}

	if err := context.BodyParser(&userDto); err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = err.Error()

		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	// assign ID must be place before validation
	userDto.UUID = uid
	userDto.UpdatedByID = userID

	//validate
	vEs, err := handler.userValidator.Validate(userDto)

	if err != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = "Validation error"
		resp.Data = vEs

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	if _, err := handler.userService.Update(userDto); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]interface{}{"user": userDto}

	return context.Status(http.StatusOK).JSON(resp)
}

// DeleteHandle : handle
func (handler *userHandler) DeleteHandle(context *fiber.Ctx) (err error) {

	var resp response.Response
	userID, _ := handler.GetUserID(context)

	uid, err := uuid.Parse(context.Params("uid"))
	if err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Exception Error: " + err.Error()
		return context.Status(http.StatusExpectationFailed).JSON(resp)
	}

	userDto, err := handler.userService.Read(uid)

	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Record not found"
		return context.Status(http.StatusNotFound).JSON(resp)
	}

	if err := handler.userService.Delete(userDto.UUID, userID); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)

	return context.Status(http.StatusOK).JSON(resp)
}
