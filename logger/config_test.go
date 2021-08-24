package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewConfig(t *testing.T) {
	zap, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	tests := []struct {
		desc string
		have Config
		want Config
	}{
		{
			desc: "Pass nil Next() function",
			have: newConfig(Config{
				Next: nil,
			}),
			want: Config{
				Next: nil,
			},
		},
		{
			desc: "Pass zap logger",
			have: newConfig(Config{
				Zap: zap,
			}),
			want: Config{
				Zap: zap,
			},
		},
		{
			desc: "Pass nil Next() function and zap logger",
			have : newConfig(Config{
				Next: nil,
				Zap: zap,
			}),
			want: Config{
				Next: nil,
				Zap: zap,
			},
		},
		{
			desc: "Pass no parameter",
			have: newConfig(),
			want: configDefault,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.want, test.have)
		})
	}
}

func TestDefaultZapLogger(t *testing.T) {
	pZap, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	tests := []struct {
		desc string
		have *zap.Logger
		want *zap.Logger
	}{
		{
			desc: "Default Zap logger has Info log level",
			have: pZap,
			want: &zap.Logger{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			test.want = defaultZapLogger()
			assert.True(t, test.want.Core().Enabled(zapcore.InfoLevel))
			assert.NotPanics(t, func(){
				test.want = defaultZapLogger()
			})

		})
	}
}