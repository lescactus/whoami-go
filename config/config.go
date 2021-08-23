package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

const (
	// Name of the application
	appName = "whoami-go"

	// Internal GCP Load Balancer IP
	GCPTrustedProxyIP = "169.254.8.129"
)

type Config struct {
	*viper.Viper

	errorHandler fiber.ErrorHandler
	fiber *fiber.Config
}

func New(errorHandler fiber.ErrorHandler) *Config {
	config := &Config{
		Viper: viper.New(),
	}

	// Set default configurations
	config.setDefaults()

	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	// Automatically refresh environment variables
	config.AutomaticEnv()

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	// Set the default error handler
	config.setErrorHandler(errorHandler)

	config.setFiberConfig()

	return config
}

func (config *Config) setErrorHandler(errorHandler fiber.ErrorHandler) {
	config.errorHandler = errorHandler
}

func (config *Config) setDefaults()  {
	// Set default App configuration
	config.SetDefault("APP_ADDR", ":8080")
	config.SetDefault("APP_ENV", "local")
	
	// Set default server app configuration
	config.SetDefault("SERVER_FIBER_PREFORK", false)
	config.SetDefault("SERVER_FIBER_READ_TIMEOUT", 10 * time.Second)
	config.SetDefault("SERVER_FIBER_WRITE_TIMEOUT", 10 * time.Second)
	config.SetDefault("SERVER_FIBER_IDLE_TIMEOUT", 75 * time.Second)
	config.SetDefault("SERVER_FIBER_ENABLE_TRUSTED_PROXY_CHECK", true)
	config.SetDefault("SERVER_FIBER_PROXY_HEADER", "X-Forwarded-For")
	config.SetDefault("SERVER_FIBER_TRUSTED_PROXIES", []string{GCPTrustedProxyIP})
	config.SetDefault("SERVER_FIBER_DISABLE_KEEPALIVE", "false")

	// Set default views configuration
	config.SetDefault("VIEWS_TEMPLATE_DIRECTORY", "./views/templates")
	config.SetDefault("VIEWS_TEMPLATE_EXTENSIONS", ".html")
	config.SetDefault("VIEWS_STATIC_DIRECTORY", "./views/static")

	// Set default middlewares configuration
	config.SetDefault("MIDDLEWARE_RECOVER_ENABLE_STACK_TRACE", true)
}

func (config *Config) setFiberConfig() {
	engine := html.New(config.GetString("VIEWS_TEMPLATE_DIRECTORY"), config.GetString("VIEWS_TEMPLATE_EXTENSIONS"))
	
	config.fiber = &fiber.Config{
		AppName:			appName,
		Prefork:			config.GetBool("SERVER_FIBER_PREFORK"),
		ServerHeader:		config.GetString("FIBER_SERVERHEADER"),
		Views:				engine,
		ReadTimeout:		config.GetDuration("SERVER_FIBER_READ_TIMEOUT"),
		WriteTimeout:		config.GetDuration("SERVER_FIBER_WRITE_TIMEOUT"),
		IdleTimeout:		config.GetDuration("SERVER_FIBER_IDLE_TIMEOUT"),
		ProxyHeader:		config.GetString("SERVER_FIBER_PROXY_HEADER"),
		ErrorHandler:		config.errorHandler,
		DisableKeepalive:	config.GetBool("SERVER_FIBER_DISABLE_KEEPALIVE"),
	}
}

func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}