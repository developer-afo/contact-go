package router

import (
	"github.com/afolabiolayinka/contact-go/database"
	repositories "github.com/afolabiolayinka/contact-go/database/repository/app"
	services "github.com/afolabiolayinka/contact-go/database/service/app"
	handlers "github.com/afolabiolayinka/contact-go/handler/app"
	"github.com/gofiber/fiber/v2"
)

func InitializeAppRouter(router fiber.Router, dbConn database.DatabaseInterface) {

	contactRepository := repositories.NewContactRepostiory(dbConn)
	contactService := services.NewContactService(contactRepository)

	contactHandler := handlers.NewContactHandler(contactService)

	contactRoutes := router.Group("/contacts")
	contactRoutes.Get("/", contactHandler.IndexHandle)
	contactRoutes.Post("/", contactHandler.CreateHandle)
	contactRoutes.Get("/:uid", contactHandler.ReadHandle)
	contactRoutes.Patch("/:uid", contactHandler.UpdateHandle)
	contactRoutes.Delete("/:uid", contactHandler.DeleteHandle)
}
