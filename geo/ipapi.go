package geo

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

const (
	IpAPITagKey = "ipapi"
)

type IPAPI struct {
	ip string
	jsoniter jsoniter.API
}

func NewIPAPI(ip string) *IPAPI {
	var j = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey: 				IpAPITagKey,
	}.Froze()

	return &IPAPI{
		ip: ip,
		jsoniter: j,
	}
}

func (i *IPAPI) Get() (Geo, error) {
	var g Geo

	agent := fiber.AcquireAgent()
	agent.Timeout(agentTimeout)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.SetContentType("application/json")
	req.Header.Set("Accept", "application/json")

	req.SetRequestURI(ipApiBaseURL + i.ip)

	if err := agent.Parse(); err != nil {
		panic(err)
	}
	code, body, errs := agent.Bytes()

	if code > 499 {
		return g, errors.New("Error: Unexpected response code from " + ipApiBaseURL + i.ip)
	}

	if len(errs) > 0 {
		return g, errors.New("Error: " + errs[0].Error())
	}

	err := i.jsoniter.Unmarshal(body, &g)
	if err != nil {
		return g, err
	}

	fiber.ReleaseAgent(agent)

	return g, nil
}