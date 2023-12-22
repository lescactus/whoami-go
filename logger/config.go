package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// zap defines the zap logger.
	//
	// Optional. Default: zap.NewProductionConfig()
	Zap *zap.Logger
}

var configDefault = Config{
	Next: nil,
	Zap:  defaultZapLogger(),
}

// Helper function to set default values
func newConfig(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return configDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = configDefault.Next
	}

	return cfg
}

func defaultZapLogger() *zap.Logger {
	zap, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return zap
}
