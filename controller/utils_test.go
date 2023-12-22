package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	XForwardedForHeaderPascalCase  = "X-Forwarded-For"
	XForwardedForHeaderLowerCase   = "x-forwarded-for"
	XForwardedForHeaderUpperCase   = "X-FORWARDED-FOR"
	XForwardedHostHeaderPascalCase = "X-Forwarded-Host"
	XForwardedHostHeaderLowerCase  = "x-forwarded-host"
	XForwardedHostHeaderUpperCase  = "X-FORWARDED-HOST"
	XCustomHeaderPascalCase        = "X-Custom"
	XCustomHeaderLowerCase         = "x-custom"
	XCustomHeaderUpperCase         = "X-CUSTOM"
)

func TestRemoveCustomHeaders(t *testing.T) {
	tests := []struct {
		desc string
		have map[string]string
		want map[string]string
	}{
		{
			desc: "Headers map contains " + XForwardedForHeaderPascalCase + " header",
			have: map[string]string{
				"Host":                        "127.0.0.1",
				"Accept":                      "*/*",
				XForwardedForHeaderPascalCase: "1.1.1.1",
				"Accept-Encoding":             "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XForwardedForHeaderLowerCase + " header",
			have: map[string]string{
				"Host":                       "127.0.0.1",
				"Accept":                     "*/*",
				XForwardedForHeaderLowerCase: "1.1.1.1",
				"Accept-Encoding":            "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XForwardedForHeaderUpperCase + " header",
			have: map[string]string{
				"Host":                       "127.0.0.1",
				"Accept":                     "*/*",
				XForwardedForHeaderUpperCase: "1.1.1.1",
				"Accept-Encoding":            "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XForwardedHostHeaderPascalCase + " header",
			have: map[string]string{
				"Host":                         "127.0.0.1",
				"Accept":                       "*/*",
				XForwardedHostHeaderPascalCase: "app.run",
				"Accept-Encoding":              "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XForwardedHostHeaderLowerCase + " header",
			have: map[string]string{
				"Host":                        "127.0.0.1",
				"Accept":                      "*/*",
				XForwardedHostHeaderLowerCase: "app.run",
				"Accept-Encoding":             "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XForwardedHostHeaderLowerCase + " header",
			have: map[string]string{
				"Host":                        "127.0.0.1",
				"Accept":                      "*/*",
				XForwardedHostHeaderLowerCase: "app.run",
				"Accept-Encoding":             "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XCustomHeaderPascalCase + " header",
			have: map[string]string{
				"Host":                  "127.0.0.1",
				"Accept":                "*/*",
				XCustomHeaderPascalCase: "Custom",
				"Accept-Encoding":       "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XCustomHeaderLowerCase + " header",
			have: map[string]string{
				"Host":                 "127.0.0.1",
				"Accept":               "*/*",
				XCustomHeaderLowerCase: "custom",
				"Accept-Encoding":      "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
		{
			desc: "Headers map contains " + XCustomHeaderUpperCase + " header",
			have: map[string]string{
				"Host":                 "127.0.0.1",
				"Accept":               "*/*",
				XCustomHeaderUpperCase: "CUSTOM",
				"Accept-Encoding":      "gzip, deflate",
			},
			want: map[string]string{
				"Host":            "127.0.0.1",
				"Accept":          "*/*",
				"Accept-Encoding": "gzip, deflate",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			m := removeCustomHeaders(test.have)
			assert.Equal(t, test.want, m)
		})
	}
}

func BenchmarkRemoveCustomHeaders(b *testing.B) {
	m := map[string]string{
		"Host":                 "127.0.0.1",
		"Accept":               "*/*",
		XCustomHeaderUpperCase: "CUSTOM",
		"Accept-Encoding":      "gzip, deflate",
	}

	for n := 0; n < b.N; n++ {
		_ = removeCustomHeaders(m)
	}
}
