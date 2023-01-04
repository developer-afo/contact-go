package router

import (
	"github.com/afolabiolayinka/contact-go/database"
	repositories "github.com/afolabiolayinka/contact-go/database/repository/auth"
	services "github.com/afolabiolayinka/contact-go/database/service/auth"
	handlers "github.com/afolabiolayinka/contact-go/handler/auth"
	"github.com/gofiber/fiber/v2"
)

func InitializeAuthRouter(router fiber.Router, protectedRouter fiber.Router, dbConn database.DatabaseInterface) {

	userRepository := repositories.NewUserRepostiory(dbConn)
	userService := services.NewUserService(userRepository)

	userHandler := handlers.NewUserHandler(userService)

	authHandler := handlers.NewAuthHandler(userService)

	router.Post("/login", authHandler.LoginHandle)
	router.Post("/register", authHandler.RegisterHandle)

	userRoutes := protectedRouter.Group("/users")
	userRoutes.Get("/", userHandler.IndexHandle)
	userRoutes.Post("/", userHandler.CreateHandle)
	userRoutes.Get("/:uid", userHandler.ReadHandle)
	userRoutes.Patch("/:uid", userHandler.UpdateHandle)
	userRoutes.Delete("/:uid", userHandler.DeleteHandle)
}
