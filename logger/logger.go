package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func New(config ...Config) fiber.Handler {
	cfg := newConfig(config...)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// override error handler
		errHandler := c.App().Config().ErrorHandler

		var start, stop time.Time

		// Start request duration timer
		start = time.Now()

		// Continue stack
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		// Stop request duration timer
		stop = time.Now()

		if chainErr != nil {
			cfg.Zap.Error(chainErr.Error(),
				zap.String("ip", c.IP()),
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.Int("code", c.Response().StatusCode()),
				zap.String("user-agent", c.Get("User-Agent")),
				zap.String("requestid", string(c.Response().Header.Peek("X-Request-Id"))),
				zap.Duration("duration", stop.Sub(start).Round(time.Millisecond)))
		} else {
			cfg.Zap.Info("",
				zap.String("ip", c.IP()),
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.Int("code", c.Response().StatusCode()),
				zap.String("user-agent", c.Get("User-Agent")),
				zap.String("requestid", string(c.Response().Header.Peek("X-Request-Id"))),
				zap.Duration("duration", stop.Sub(start).Round(time.Millisecond)))
		}
		return nil
	}
}
