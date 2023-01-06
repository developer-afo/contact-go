package auth

import (
	"net/http"

	dtos "github.com/afolabiolayinka/contact-go/database/dto/auth"
	services "github.com/afolabiolayinka/contact-go/database/service/auth"
	"github.com/afolabiolayinka/contact-go/handler"
	"github.com/afolabiolayinka/contact-go/helper"
	authRequest "github.com/afolabiolayinka/contact-go/payload/request/auth"
	"github.com/afolabiolayinka/contact-go/payload/response"
	"github.com/afolabiolayinka/contact-go/security"
	validators "github.com/afolabiolayinka/contact-go/validator/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlerInterface interface {
	LoginHandle(context *fiber.Ctx) error
	RegisterHandle(context *fiber.Ctx) error
}

type authHandler struct {
	handler.BaseHandler
	userService   services.UserServiceInterface
	userValidator validators.UserValidator
	hashUtils     helper.Argon2
}

func NewAuthHandler(
	userService services.UserServiceInterface,
) AuthHandlerInterface {
	return &authHandler{
		userService: userService,
	}
}

func (handler *authHandler) LoginHandle(c *fiber.Ctx) error {
	var resp response.Response
	var signinRequest authRequest.LoginRequest

	if err := c.BodyParser(&signinRequest); err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Error on signin request"
		resp.Data = map[string]interface{}{"Error": err.Error()}

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	user, err := handler.userService.Authenticate(signinRequest.Email, signinRequest.Password)

	if err != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = err.Error()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
	}

	resp.Status = http.StatusAccepted
	resp.Message = "Login Successful"

	accessToken, jwtErr := security.CreateToken(user.UUID.String())
	if jwtErr != nil {
		return jwtErr
	}

	if err != nil {
		return err
	}

	resp.Data = map[string]interface{}{"user": user, "tokens": map[string]interface{}{"accessToken": accessToken}}

	return c.Status(http.StatusOK).JSON(resp)
}

func (handler *authHandler) RegisterHandle(context *fiber.Ctx) error {
	var resp response.Response
	var signupRequest authRequest.RegisterRequest

	if err := context.BodyParser(&signupRequest); err != nil {
		resp.Status = http.StatusExpectationFailed
		resp.Message = "Error on signup request"
		resp.Data = map[string]interface{}{"Error": err.Error()}

		return context.Status(fiber.StatusBadRequest).JSON(resp)
	}

	var userDto dtos.UserDTO
	userDto.Email = signupRequest.Email
	userDto.Password, _ = handler.hashUtils.HashPassword(signupRequest.Password)

	vEs, verr := handler.userValidator.Validate(userDto)

	if verr != nil {
		resp.Status = http.StatusUnprocessableEntity
		resp.Message = "Validation error"
		resp.Data = vEs

		return context.Status(http.StatusUnprocessableEntity).JSON(resp)
	}

	userDto, err := handler.userService.Create(userDto)
	if err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = err.Error()

		return context.Status(http.StatusInternalServerError).JSON(resp)
	}

	handler.userService.Create(userDto)

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)

	return context.Status(http.StatusOK).JSON(resp)
}
