package geo

import (
	"time"
)

const (
	freeGeoIPAppBaseURL = "https://freegeoip.app/json/"

	ipApiBaseURL = "http://ip-api.com/json/" // HTTPS not available with the free version
)

var (
	agentTimeout time.Duration
)

type Geo struct {
	IP          string  `json:"ip" ipapi:"query"`
	CountryCode string  `json:"country_code" ipapi:"countryCode"`
	CountryName string  `json:"country_name" ipapi:"country"`
	RegionCode  string  `json:"region_code" ipapi:"region"`
	RegionName  string  `json:"region_name" ipapi:"regionName"`
	City        string  `json:"city" ipapi:"city"`
	ZipCode     string  `json:"zip_code" ipapi:"zip"`
	Timezone    string  `json:"time_zone" ipapi:"timezone"`
	Latitude    float64 `json:"latitude" ipapi:"lat"`
	Longitude   float64 `json:"longitude" ipapi:"lon"`
	MetroCode   int     `json:"metro_code" ipapi:"metro"`
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
