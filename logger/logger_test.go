package logger

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	capturer "github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logLine struct {
	Level     string        `json:"level"`
	Timestamp string        `json:"ts"`
	Message   string        `json:"msg"`
	IP        string        `json:"ip"`
	Method    string        `json:"method"`
	Path      string        `json:"path"`
	Code      int           `json:"code"`
	UserAgent string        `json:"user-agent"`
	Requestid string        `json:"requestid"`
	Duration  time.Duration `json:"duration"`
}

func TestNew(t *testing.T) {

	errorMessage := "ERROR"

	cfgZapJson := zap.Config{
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
	}
	cfgZapJson.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	// Capture both stdin and stderr
	//
	// Scenario:
	// Create a	route responding with an error
	// Ensure the middleware logger is showing the error in the log
	out := capturer.CaptureOutput(func() {

		zap, err := cfgZapJson.Build()
		if err != nil {
			panic(err)
		}

		app := fiber.New()

		app.Use(New(Config{
			Next: nil,
			Zap:  zap,
		}))

		app.Get("/error", func(c *fiber.Ctx) error {
			return errors.New(errorMessage)
		})

		resp, err := app.Test(httptest.NewRequest("GET", "/error", nil))
		assert.Equal(t, nil, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	var logline logLine
	err := json.Unmarshal([]byte(out), &logline)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, errorMessage, logline.Message)

	// Capture both stdin and stderr
	//
	// Scenario:
	// Create a	route responding without error
	// Ensure the middleware logger is showing no error in the log
	out = capturer.CaptureOutput(func() {

		zap, err := cfgZapJson.Build()
		if err != nil {
			panic(err)
		}

		app := fiber.New()

		app.Use(New(Config{
			Next: nil,
			Zap:  zap,
		}))

		// Middleware always returning no error
		app.Use(func(c *fiber.Ctx) (err error) {
			return nil
		})

		app.Get("/", func(c *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
		assert.Equal(t, nil, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	err = json.Unmarshal([]byte(out), &logline)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, "", logline.Message)
}
