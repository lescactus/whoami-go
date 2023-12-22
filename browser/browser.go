package browser

import (
	"github.com/lescactus/whoami-go/geo"
)

type Browser struct {
	ip        string
	port      string
	host      string
	headers   map[string]string
	url       string
	lang      string
	userAgent string
	geo       *geo.Geo
}

func NewBrowser(ip string, port string, host string, headers map[string]string, url string, lang string, userAgent string, geo *geo.Geo) *Browser {
	return &Browser{
		ip:        ip,
		port:      port,
		host:      host,
		headers:   headers,
		url:       url,
		lang:      lang,
		userAgent: userAgent,
		geo:       geo,
	}
}

func (b *Browser) GetIP() string {
	return b.ip
}

func (b *Browser) GetPort() string {
	return b.port
}

func (b *Browser) GetHost() string {
	return b.host
}

func (b *Browser) GetHeaders() map[string]string {
	return b.headers
}

func (b *Browser) GetURL() string {
	return b.url
}

func (b *Browser) GetLang() string {
	return b.lang
}

func (b *Browser) GetUserAgent() string {
	return b.userAgent
}

func (b *Browser) GetLongitude() float64 {
	return b.geo.Longitude
}

func (b *Browser) GetLatitude() float64 {
	return b.geo.Latitude
}

func (b *Browser) GetCountryName() string {
	return b.geo.CountryName
}

func (b *Browser) GetCountryCode() string {
	return b.geo.CountryCode
}

func (b *Browser) GetRegionCode() string {
	return b.geo.RegionCode
}

func (b *Browser) GetRegionName() string {
	return b.geo.RegionName
}

func (b *Browser) GetCity() string {
	return b.geo.City
}

func (b *Browser) GetZipCode() string {
	return b.geo.ZipCode
}

func (b *Browser) GetTimezone() string {
	return b.geo.Timezone
}

func (b *Browser) GetMetroCode() int {
	return b.geo.MetroCode
}
