package browser

import (
	"testing"

	"github.com/lescactus/whoami-go/geo"
	"github.com/stretchr/testify/assert"
)

var (
	ip         = "127.0.0.1"
	port       = "40000"
	host       = "example.com"
	headers    = map[string]string{"header1": "value1", "header2": "value2"}
	url        = "https://example.com/"
	lang       = "fr-FR,fr;q=0.5"
	userAgent  = "Fake User Agent"
	browserGeo = &geo.Geo{
		CountryCode: "FR",
		CountryName: "France",
		RegionCode:  "IDF",
		RegionName:  "Île-de-France",
		City:        "Paris",
		ZipCode:     "75008",
		Timezone:    "Europe/Paris",
		Latitude:    48.8714,
		Longitude:   2.32141,
		MetroCode:   8,
	}

	browser = &Browser{
		ip:        ip,
		port:      port,
		host:      host,
		headers:   headers,
		url:       url,
		lang:      lang,
		userAgent: userAgent,
		geo:       browserGeo,
	}
)

func TestNewBrowser(t *testing.T) {
	t.Run("NewBrowser()", func(t *testing.T) {
		b := NewBrowser(ip, port, host, headers, url, lang, userAgent, browserGeo)

		assert.Equal(t, b, browser)
		assert.NotEmpty(t, b)
	})

}

func TestGets(t *testing.T) {
	testsStrings := []struct {
		desc string
		have string
		want string
	}{
		{
			desc: "Ensure Browser IP is not null",
			have: browser.GetIP(),
			want: ip,
		},
		{
			desc: "Ensure Browser Port is not null",
			have: browser.GetPort(),
			want: port,
		},
		{
			desc: "Ensure Browser Host is not null",
			have: browser.GetHost(),
			want: host,
		},
		{
			desc: "Ensure Browser URL is not null",
			have: browser.GetURL(),
			want: url,
		},
		{
			desc: "Ensure Browser Lang is not null",
			have: browser.GetLang(),
			want: lang,
		},
		{
			desc: "Ensure Browser User Agent is not null",
			have: browser.GetUserAgent(),
			want: userAgent,
		},
		{
			desc: "Ensure Browser Country Name is not null",
			have: browser.GetCountryName(),
			want: browserGeo.CountryName,
		},
		{
			desc: "Ensure Browser Country Code is not null",
			have: browser.GetCountryCode(),
			want: browserGeo.CountryCode,
		},
		{
			desc: "Ensure Browser Region Code is not null",
			have: browser.GetRegionCode(),
			want: browserGeo.RegionCode,
		},
		{
			desc: "Ensure Browser Region Name is not null",
			have: browser.GetRegionName(),
			want: browserGeo.RegionName,
		},
		{
			desc: "Ensure Browser City is not null",
			have: browser.GetCity(),
			want: browserGeo.City,
		},
		{
			desc: "Ensure Browser ZipCode is not null",
			have: browser.GetZipCode(),
			want: browserGeo.ZipCode,
		},
		{
			desc: "Ensure Browser Timezone is not null",
			have: browser.GetTimezone(),
			want: browserGeo.Timezone,
		},
	}

	testsFloats := []struct {
		desc string
		have float64
		want float64
	}{
		{
			desc: "Ensure Browser Longitude is not null",
			have: browser.GetLongitude(),
			want: browserGeo.Longitude,
		},
		{
			desc: "Ensure Browser Latitude is not null",
			have: browser.GetLatitude(),
			want: browserGeo.Latitude,
		},
	}

	testsInts := []struct {
		desc string
		have int
		want int
	}{
		{
			desc: "Ensure Browser Metro Code is not null",
			have: browser.GetMetroCode(),
			want: browserGeo.MetroCode,
		},
	}

	testsMapStrings := []struct {
		desc string
		have map[string]string
		want map[string]string
	}{
		{
			desc: "Ensure Browser Headers Code is not null",
			have: browser.GetHeaders(),
			want: headers,
		},
	}

	for _, test := range testsStrings {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.want, test.have)
			assert.NotEmpty(t, test.have)
		})
	}

	for _, test := range testsFloats {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.want, test.have)
			assert.NotEmpty(t, test.have)
		})
	}

	for _, test := range testsMapStrings {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.want, test.have)
			assert.NotEmpty(t, test.have)
		})
	}

	for _, test := range testsInts {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.want, test.have)
			assert.NotEmpty(t, test.have)
		})
	}
}
