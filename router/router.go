package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func InitializeRouter(router *fiber.App) {
	router.Get("/monitor", monitor.New())

	InitializeAuthRouter(router)
	InitializeAppRouter(router)
}
