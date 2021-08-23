package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/lescactus/whoami-go/config"
	"github.com/lescactus/whoami-go/controller"
)

func main() {
	config := config.New(controller.DefaultErrorHandler)
	
	app := fiber.New(*config.GetFiberConfig())
		
	// Serve static assets
	app.Static("/static", config.GetString("VIEWS_STATIC_DIRECTORY"))

	// Middlewares registration
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: config.GetBool("MIDDLEWARE_RECOVER_ENABLE_STACK_TRACE"),
	}))
	app.Use(compress.New())

	app.Get("/", controller.IndexHandler)
	app.Get("/index", controller.IndexHandler)
	app.Get("/ip", controller.IPHandler)
	app.Get("/port", controller.PortHandler)
	app.Get("/lang", controller.LangHandler)
	app.Get("/ua", controller.UserAgentHandler)
	app.Get("/raw/go", controller.RawGoHandler)
	app.Get("/raw/json", controller.RawJSONHandler)
	app.Get("/raw/yaml", controller.RawYAMLHandler)

	// Start listening on the specified address
	err := app.Listen(config.GetString("APP_ADDR"))
	if err != nil {
		app.Shutdown()
	}
}