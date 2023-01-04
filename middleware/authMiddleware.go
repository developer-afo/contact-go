package middleware

import (
	"net/http"

	"github.com/afolabiolayinka/contact-go/security"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	err := jwtware.New(jwtware.Config{
		SigningKey:   []byte(security.GetSecret()),
		ErrorHandler: jwtError,
	})

	if err != nil {
		return err
	}

	return checkUser()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"status": http.StatusUnauthorized, "message": "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusForbidden).
		JSON(fiber.Map{"status": http.StatusForbidden, "message": "Invalid or expired JWT"})
}

func checkUser() fiber.Handler {
	return func(context *fiber.Ctx) (err error) {
		_, err = security.ExtractUserID(context.Request())
		if err != nil {
			context.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"status": http.StatusForbidden, "message": "Unable to extract User"})
		}

		return context.Next()
	}
}
