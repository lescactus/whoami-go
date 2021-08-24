package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	zapConfig *zap.Config
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
	config.setZapConfig()

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

	// Set default IP Geolocation API
	config.SetDefault("GEOLOCATION_API", "freegeoip") // Availables: "freegeoip" or "ipapi"

	// Set default http client configuration
	config.SetDefault("HTTP_CLIENT_TIMEOUT", 5 * time.Second)

	// Set default middlewares configuration
	config.SetDefault("MIDDLEWARE_RECOVER_ENABLE_STACK_TRACE", true)

	// Set default loggers configuration
	config.SetDefault("LOGGER_TYPE", "gofiber") // "gofiber" or "zap"
	config.SetDefault("LOGGER_ZAP_LOG_LEVEL", "info") // "info", "debug" or "error"
	config.SetDefault("LOGGER_ZAP_DEVELOPMENT_MODE", false)
	config.SetDefault("LOGGER_ZAP_DISABLE_CALLER", true)
	config.SetDefault("LOGGER_ZAP_DISABLE_STACK_TRACE", true)
	config.SetDefault("LOGGER_ZAP_ENCODING", "json") // "json" or "console"

}

func (config *Config) setFiberConfig() {
	engine := html.New(config.GetString("VIEWS_TEMPLATE_DIRECTORY"), config.GetString("VIEWS_TEMPLATE_EXTENSIONS"))
	
	config.fiber = &fiber.Config{
		AppName:					appName,
		Prefork:					config.GetBool("SERVER_FIBER_PREFORK"),
		Views:						engine,
		ReadTimeout:				config.GetDuration("SERVER_FIBER_READ_TIMEOUT"),
		WriteTimeout:				config.GetDuration("SERVER_FIBER_WRITE_TIMEOUT"),
		IdleTimeout:				config.GetDuration("SERVER_FIBER_IDLE_TIMEOUT"),
		EnableTrustedProxyCheck:	config.GetBool("SERVER_FIBER_ENABLE_TRUSTED_PROXY_CHECK"),
		ProxyHeader:				config.GetString("SERVER_FIBER_PROXY_HEADER"),
		TrustedProxies: 			config.GetStringSlice("SERVER_FIBER_TRUSTED_PROXIES"),
		ErrorHandler:				config.errorHandler,
		DisableKeepalive:			config.GetBool("SERVER_FIBER_DISABLE_KEEPALIVE"),
	}
}

func (config *Config) setZapConfig() {
	logLevel := config.GetString("LOGGER_ZAP_LOG_LEVEL")

	cfg := zap.Config{
		Development:       config.GetBool("LOGGER_ZAP_DEVELOPMENT_MODE"),
		DisableCaller:     config.GetBool("LOGGER_ZAP_DISABLE_CALLER"),
		DisableStacktrace: config.GetBool("LOGGER_ZAP_DISABLE_STACK_TRACE"),
		Encoding:         config.GetString("LOGGER_ZAP_ENCODING"),
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	if logLevel == "error" {
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}
	if logLevel == "info" {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	if logLevel == "debug" {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.DisableCaller = false
	}
	config.zapConfig = &cfg
}

func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}

func (config *Config) GetZapConfig() *zap.Config {
	return config.zapConfig
}