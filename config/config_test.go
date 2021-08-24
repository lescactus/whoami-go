package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetZapConfig(t *testing.T) {
	t.Run("Ensure Zap Config is not null", func(t *testing.T) {
		var errorHandlerFunc func(*fiber.Ctx, error) error

		cfg := New(errorHandlerFunc)
		zapCfg := cfg.GetZapConfig()
	
		assert.NotEmpty(t, zapCfg)
	})
}

func TestGetFiberConfig(t *testing.T) {
	t.Run("Ensure Fiber Config is not null", func(t *testing.T) {
		var errorHandlerFunc func(*fiber.Ctx, error) error

		cfg := New(errorHandlerFunc)
		fiberCfg := cfg.GetFiberConfig()

		assert.NotEmpty(t, fiberCfg)
	})
}

func TestSetZapConfig(t *testing.T) {
	tests := []struct {
		desc string
		have string
		want string
	}{
		{
			desc: "LOGGER_ZAP_ENCODING set to json",
			have: "json",
			want: "json",
		},
		{
			desc: "LOGGER_ZAP_ENCODING set to console",
			have: "console",
			want: "console",
		},
		
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			os.Setenv("LOGGER_ZAP_ENCODING", test.have)

			var errorHandlerFunc func(*fiber.Ctx, error) error

			cfg := New(errorHandlerFunc)
			cfg.setZapConfig()
			zapCfg := cfg.GetZapConfig()

			assert.NotEmpty(t, zapCfg)
			assert.Equal(t, test.want, zapCfg.Encoding)
		})
	}

	
	testsLogLevel := []struct {
		desc string
		loglevel string
		want zap.AtomicLevel
	}{
		{
			desc: "Set log level to info",
			loglevel: "info",
			want: zap.NewAtomicLevelAt(zap.InfoLevel),
		},
		{
			desc: "Set log level to error",
			loglevel: "error",
			want: zap.NewAtomicLevelAt(zap.ErrorLevel),
		},
		{
			desc: "Set log level to debug",
			loglevel: "debug",
			want: zap.NewAtomicLevelAt(zap.DebugLevel),
		},
	}

	for _, test := range testsLogLevel {
		t.Run(test.desc, func(t *testing.T) {
			os.Setenv("LOGGER_ZAP_LOG_LEVEL", test.loglevel)

			var errorHandlerFunc func(*fiber.Ctx, error) error

			cfg := New(errorHandlerFunc)
			cfg.setZapConfig()
			zapCfg := cfg.GetZapConfig()

			assert.NotEmpty(t, zapCfg)
			assert.Equal(t, test.want, zapCfg.Level)
		})
	}
}

func TestSetFiberConfig(t *testing.T) {
	tests := []struct {
		desc string
		have bool
		want bool
	}{
		{
			desc: "SERVER_FIBER_PREFORK set to true",
			have: true,
			want: true,
		},
		{
			desc: "SERVER_FIBER_PREFORK set to false",
			have: false,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			os.Setenv("SERVER_FIBER_PREFORK", strconv.FormatBool(test.have))

			var errorHandlerFunc func(*fiber.Ctx, error) error

			cfg := New(errorHandlerFunc)
			cfg.setFiberConfig()
			fiberCfg := cfg.GetFiberConfig()

			assert.NotEmpty(t, fiberCfg)
			assert.Equal(t, test.want, fiberCfg.Prefork)
		})
	}
}