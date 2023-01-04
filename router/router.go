package router

import (
	"github.com/afolabiolayinka/contact-go/database"
	"github.com/afolabiolayinka/contact-go/handler"
	"github.com/afolabiolayinka/contact-go/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func InitializeRouter(router *fiber.App, dbConn database.DatabaseInterface) {
	router.Get("/monitor", monitor.New())

	router.Get("/", handler.Index)

	// Authentication
	auth := router.Group("/auth")

	// Application
	app := router.Group("/app", middleware.Protected())

	InitializeAuthRouter(auth, app, dbConn)
	InitializeAppRouter(app, dbConn)
}
