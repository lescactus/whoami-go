package geo

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type FreeGeoIPApp struct {
	ip string
}

func NewFreeGeoIPApp(ip string) *FreeGeoIPApp {

	return &FreeGeoIPApp{
		ip: ip,
	}
}

func (f *FreeGeoIPApp) Get() (Geo, error) {
	var g Geo

	agent := fiber.AcquireAgent()
	agent.Timeout(agentTimeout)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.SetContentType("application/json")
	req.Header.Set("Accept", "application/json")

	req.SetRequestURI(freeGeoIPAppBaseURL + f.ip)

	if err := agent.Parse(); err != nil {
		return g, err
	}
	code, body, errs := agent.Bytes()

	if code > 499 {
		return g, errors.New("Error: Unexpected response code from " + freeGeoIPAppBaseURL + f.ip)
	}

	if len(errs) > 0 {
		return g, errors.New("Error: " + errs[0].Error())
	}

	err := json.Unmarshal(body, &g)
	if err != nil {
		return g, err
	}

	return g, nil
}
