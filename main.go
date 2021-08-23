package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"

	"github.com/lescactus/whoami-go/controller"
)

const (
	appName = "whoami-go"

	listen = "0.0.0.0:8080"

	templateDir = "./views/templates"

	templateExt = ".html"

	staticDir = "./views/static"

	GCPTrustedProxyIP = "169.254.8.129"
)

func main() {
	engine := html.New(templateDir, templateExt)

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   75 * time.Second,
		EnableTrustedProxyCheck: true,
		ProxyHeader: "X-Forwarded-For",
		TrustedProxies: []string{GCPTrustedProxyIP},
		AppName: appName,
		Views:   engine,
	})

	app.Static("/static", staticDir)

	app.Use(requestid.New())
	app.Use(logger.New())

	app.Get("/", controller.IndexHandler)
	app.Get("/index", controller.IndexHandler)
	app.Get("/ip", controller.IPHandler)
	app.Get("/port", controller.PortHandler)
	app.Get("/lang", controller.LangHandler)
	app.Get("/ua", controller.UserAgentHandler)
	app.Get("/raw/go", controller.RawGoHandler)
	app.Get("/raw/json", controller.RawJSONHandler)
	app.Get("/raw/yaml", controller.RawYAMLHandler)

	log.Fatal(app.Listen(listen))
}
