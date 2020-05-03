package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
)

var (
	// Directory where the htlm templates are stored
	templatesDir = "./templates/"

	// Index html template name
	indexName = "index.html"

	// Parse all templates into a single *Template
	templates = template.Must(template.ParseGlob(templatesDir + "*"))

	// Remove all 'x-' or 'X-' HTTP headers (Ex: 'X-Forwarded-For')
	headersToRemoveRegex = "(^((forwarded|Forwarded).*)|((x|X)-(.*))$)"

	// GeoIP API URL (https://freegeoip.app/)
	urlGeoIP = "https://freegeoip.app/json/"

	// In most Cloud environments, reverse proxies add HTTP headers to read the real client IP from
	// These slices contains the most commons
	httpIPHeaders = []string{"X-Real-IP", "x-real-ip"}
	xForwardedFor = []string{"X-Forwarded-For", "x-forwarded-for"}
)

// Browser is a structure containing browser infos
type Browser struct {
	IP        string
	Port      int
	Host      string
	Headers   map[string][]string
	URL       *url.URL
	Lang      string
	UserAgent string
	Location  *Location
	JSON      string
	YAML      string
}

// Location provide browser Geo IP location informations
type Location struct {
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

// Wrapper function to http.HandleFunc()
func makeHandler(fn func(http.ResponseWriter, *http.Request, *Browser), b bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		browser, err := getBroswerInfo(r, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("%s %s %s ", browser.IP, r.Method, r.URL)
		fn(w, r, browser)
	}
}

// Render the index html page
func renderIndexTemplate(w http.ResponseWriter, b *Browser) {
	err := templates.ExecuteTemplate(w, indexName, b)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error while rendring html template", http.StatusInternalServerError)
	}
}

// https://husobee.github.io/golang/ip-address/2015/12/17/remote-ip-go.html
func getIPAddress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip", "x-forwarded-for", "x-real-ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() {
				// bad address, go to next
				continue
			}
			return ip
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	return ip
}

// Get the client remote port from a string of type http.Request.RemoteAddr
// Port is an integer
func parsePortFromRemoteAddr(r string) (int, error) {
	_, p, err := net.SplitHostPort(r)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	// Cast port from string to integer
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	return port, nil
}

// From the IP of the client, call the free GeoIP database https://freegeoip.app/json/<ip>
// to get GeoIP infos, such as country name, city name or coordinates.
// Return *Location
func getLocationInfo(ip string) (*Location, error) {
	// New http request to https://freegeoip.app/json/<ip>
	req, _ := http.NewRequest("GET", urlGeoIP+ip, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	log.Printf("Calling: %s\n", urlGeoIP+ip)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Unmarshal json response to Location structure
	var location Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}
	return &location, nil

}

// Parse the client *http.Request and store informations in a Browser struct.
func getBroswerInfo(r *http.Request, getLocation bool) (*Browser, error) {
	// Get client IP address
	ip := getIPAddress(r)

	// Get client remote port
	port, err := parsePortFromRemoteAddr(r.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// Get http host
	host := r.Host

	// Get http headers
	headers := r.Header

	// 'Host' header is removed from Request.Header and promoted to Request.Host
	// Manually add it to Browser.Headers
	headers["Host"] = []string{host}

	// Remove infrastructure http headers (like 'X-Forwarded-For' and such)
	for header := range headers {
		matched, _ := regexp.Match(headersToRemoveRegex, []byte(header))
		if matched {
			delete(headers, header)
		}
	}

	// Get client URL
	url := r.URL

	// Get client browser language from 'Accept-Languages' http header
	lang := ""
	if headers["Accept-Language"] != nil {
		lang = strings.Join(headers["Accept-Language"], "")
	}

	// Get client browser User-Agent from 'User-Agent' http header
	userAgent := ""
	if headers["User-Agent"] != nil {
		userAgent = strings.Join(headers["User-Agent"], "")
	}

	// Json string containing client brower http headers
	jsonData := ""
	j, err := json.Marshal(headers)
	if err == nil {
		jsonData = string(j)
	}

	// Yaml string containing client brower http headers
	yamlData := ""
	y, err := yaml.Marshal(headers)
	if err == nil {
		yamlData = string(y)
	}

	// Don't get GeoIP location informations if not used (ex: /ip, /lang, ...)
	if getLocation {
		location, err := getLocationInfo(ip)
		if err != nil {
			log.Fatal(err)
		}
		return &Browser{
			IP:        ip,
			Port:      port,
			Host:      host,
			Headers:   headers,
			URL:       url,
			Lang:      lang,
			UserAgent: userAgent,
			Location:  location,
			JSON:      jsonData,
			YAML:      yamlData,
		}, nil
	}

	return &Browser{
		IP:        ip,
		Port:      port,
		Host:      host,
		Headers:   headers,
		URL:       url,
		Lang:      lang,
		UserAgent: userAgent,
		Location:  nil,
		JSON:      jsonData,
		YAML:      yamlData,
	}, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	renderIndexTemplate(w, b)
}

func ipHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	ip := b.IP
	fmt.Fprintln(w, ip)
}

func portHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	port := b.Port
	fmt.Fprintln(w, port)
}

func hostHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	host := b.Host
	fmt.Fprintln(w, host)
}

func headersHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	headers := b.Headers
	fmt.Fprintln(w, headers)
}

func langHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	lang := b.Headers["Accept-Language"]
	fmt.Fprintln(w, strings.Join(lang, ""))
}

func userAgentHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	ua := b.Headers["User-Agent"]
	fmt.Fprintln(w, strings.Join(ua, ""))
}

func debugHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	fmt.Fprintln(w, r)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintln(w, r.Header)
}

func rawJSONHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	w.Header().Set("Content-type", "application/json")

	fmt.Fprintln(w, string(b.JSON))
}

func rawYAMLHandler(w http.ResponseWriter, r *http.Request, b *Browser) {
	fmt.Fprintln(w, string(b.YAML))
}

func main() {
	css := http.FileServer(http.Dir("./static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", css))

	img := http.FileServer(http.Dir("./static/img"))
	http.Handle("/static/img/", http.StripPrefix("/static/img/", img))

	http.HandleFunc("/", makeHandler(rootHandler, true))
	http.HandleFunc("/ip", makeHandler(ipHandler, false))
	http.HandleFunc("/port", makeHandler(portHandler, false))
	http.HandleFunc("/host", makeHandler(hostHandler, false))
	http.HandleFunc("/lang", makeHandler(langHandler, false))
	http.HandleFunc("/ua", makeHandler(userAgentHandler, false))
	http.HandleFunc("/debug", makeHandler(debugHandler, false))
	http.HandleFunc("/raw/go", makeHandler(headersHandler, false))
	http.HandleFunc("/raw/json", makeHandler(rawJSONHandler, false))
	http.HandleFunc("/raw/yaml", makeHandler(rawYAMLHandler, false))
	log.Println("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
