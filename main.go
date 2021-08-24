package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/ansrivas/fiberprometheus/v2"

	"github.com/lescactus/whoami-go/config"
	"github.com/lescactus/whoami-go/controller"
	"github.com/lescactus/whoami-go/geo"
	whoamilogger "github.com/lescactus/whoami-go/logger"
)

func main() {
	cfg := config.New(controller.DefaultErrorHandler)

	// Set timeout of http client
	geo.SetHTTPClientTimeout(cfg.GetDuration("HTTP_CLIENT_TIMEOUT"))

	// Set the IP Getlocation API used to fetch browser IP location informations
	controller.SetGeoLocationAPI(cfg.GetString("GEOLOCATION_API"))
	
	// New fiber app using custom configuration
	app := fiber.New(*cfg.GetFiberConfig())
		
	// Serve static assets
	app.Static("/static", cfg.GetString("VIEWS_STATIC_DIRECTORY"))


	// Middlewares and logger registration
	//
	// RequestID middlewares - non optional
	app.Use(requestid.New())

	// Logger registration
	if cfg.GetString("LOGGER_TYPE") == "zap" {
		ZapLogger, e := cfg.GetZapConfig().Build() 
		if e != nil {
			panic(e)
		}
		app.Use(whoamilogger.New(whoamilogger.Config{
			Zap: ZapLogger,
		}))
	} else { // using default gofiber logger
		app.Use(logger.New())
	}
	
	
	// Prometheus middleware - optional
	if cfg.GetBool("MIDDLEWARE_PROMETHEUS_ENABLED") {
		prometheus := fiberprometheus.New(config.AppName)
		prometheus.RegisterAt(app, cfg.GetString("MIDDLEWARE_PROMETHEUS_METRICS_PATH"))
		app.Use(prometheus.Middleware)
	}

	// Recover middleware - non optional
	app.Use(recover.New(recover.Config{
		EnableStackTrace: cfg.GetBool("MIDDLEWARE_RECOVER_ENABLE_STACK_TRACE"),
	}))

	// Compress middleware - non optional
	app.Use(compress.New())

	// pprof middleware - optional
	if cfg.GetString("APP_ENV") != "production" {
		app.Use(pprof.New())
	}

	
	// Routes registration
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
	err := app.Listen(cfg.GetString("APP_ADDR"))
	if err != nil {
		app.Shutdown()
	}
}
