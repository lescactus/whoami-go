package geo

import (
	"time"
)

const (
	freeGeoIPAppBaseURL = "https://freegeoip.app/json/"

	agentTimeout = 5 * time.Second
)

type Geo struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	Timezone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

type GeoLocation interface {
	Get() (Geo, error)
}

func GetBrowserLocation(g GeoLocation) Geo {
	var geo Geo
	geo, err := g.Get()
	if err != nil {
		return Geo{}
	}
	return geo
}