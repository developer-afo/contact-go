package main

import (
	"log"
	"time"

	"github.com/afolabiolayinka/contact-go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	/*app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200, https://equalscloud.com",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))*/
	app.Use(limiter.New(limiter.Config{
		Max:               50,
		Expiration:        60 * time.Second,
		LimiterMiddleware: limiter.FixedWindow{},
	}))

	router.InitializeRouter(app)

	log.Fatal(app.Listen(":8080"))
}
