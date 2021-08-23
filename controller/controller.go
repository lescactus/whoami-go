package controller

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/gofiber/fiber/v2"
	"github.com/lescactus/whoami-go/browser"
	"github.com/lescactus/whoami-go/geo"
)

const (
	// String respone when a client http header is not present
	undefinedResponse = "<undefined>"
)

var (
	// Define which IP Geolocation API will be used to fetch browser informations
	geoLocationAPI string
)

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Set error message
	message := err.Error()

	// Check if it's a fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// Return JSON internal server error
	c.Status(code)
	return c.JSON(message)
}

func IndexHandler(c *fiber.Ctx) error {
	m := make(map[string]string)

	c.Request().Header.VisitAll(func(key []byte, value []byte) {
		m[string(key)] = string(value)
	})

	m = removeCustomHeaders(m)

	clientIP := c.IP()

	var n geo.GeoLocation
	switch geoLocationAPI {
	case "freegeoip":
		n = geo.NewFreeGeoIPApp(clientIP)
	case "ipapi":
		n = geo.NewIPAPI(clientIP)
	}
	g := geo.GetBrowserLocation(n)

	b := browser.NewBrowser(
		clientIP,
		c.Port(),
		string(c.Context().Host()),
		m,
		c.OriginalURL(),
		c.Get("Accept-Language"),
		c.Get("User-Agent"),
		&g,
	)

	j, err := json.Marshal(m)
	if err != nil {
		j = []byte("")
	}

	y, err := yaml.Marshal(m)
	if err != nil {
		y = []byte("")
	}

	return c.Render("index", fiber.Map{
		"IP": b.GetIP(),
		"Port": b.GetPort(),
		"Headers": b.GetHeaders(),
		"Host": b.GetHost(),
		"Lang": b.GetLang(),
		"UserAgent": b.GetUserAgent(),
		"Longitude": b.GetLongitude(),
		"Latitude": b.GetLatitude(),
		"CountryName": b.GetCountryName(),
		"CountryCode": b.GetCountryCode(),
		"RegionCode": b.GetRegionCode(),
		"RegionName": b.GetRegionName(),
		"City": b.GetCity(),
		"ZipCode": b.GetZipCode(),
		"Timezone": b.GetTimezone(),
		"MetroCode": b.GetMetroCode(),
		"BaseURL": c.BaseURL(),
		"JSON": string(j),
		"YAML": string(y),
	})
}

func IPHandler(c *fiber.Ctx) error {
	return c.SendString(c.IP())
}

func PortHandler(c *fiber.Ctx) error {
	return c.SendString(c.Port())
}

func LangHandler(c *fiber.Ctx) error {
	acceptLanguage := c.Get("Accept-Language")
	if acceptLanguage == "" {
		return c.SendString(undefinedResponse)
	}
	return c.SendString(acceptLanguage)
}

func UserAgentHandler(c *fiber.Ctx) error {
	ua := c.Get("User-Agent")
	if ua == "" {
		return c.SendString(undefinedResponse)
	}
	return c.SendString(ua)
}

func RawGoHandler(c *fiber.Ctx) error {
	m := make(map[string]string)

	c.Request().Header.VisitAll(func(key []byte, value []byte) {
		m[string(key)] = string(value)
	})

	m = removeCustomHeaders(m)
	return c.SendString(fmt.Sprintf("%v", m))
}

func RawJSONHandler(c *fiber.Ctx) error {
	m := make(map[string]string)

	c.Request().Header.VisitAll(func(key []byte, value []byte) {
		m[string(key)] = string(value)
	})

	m = removeCustomHeaders(m)
	return c.JSON(m)
}

func RawYAMLHandler(c *fiber.Ctx) error {
	m := make(map[string]string)

	c.Request().Header.VisitAll(func(key []byte, value []byte) {
		m[string(key)] = string(value)
	})
	m = removeCustomHeaders(m)

	y, err := yaml.Marshal(m)
	if err != nil {
		return fiber.ErrBadRequest
	}

	return c.Send(y)
}

func SetGeoLocationAPI(api string) {
	geoLocationAPI = api
}