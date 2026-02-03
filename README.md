# Who am I ?

![go](https://img.shields.io/badge/go->=1.24-blue)

A tiny Go web app that display sample information from the visitor (IP, location, browser HTTP headers ...): **[whoami](https://whoami-go.alexasr.fr/)** (Development version: [dev](https://whoami-go-dev-ujargjwndq-ew.a.run.app/)).
It is a rewriting of one of my previous mini project (**[whoami-python](https://github.com/lescactus/whoami-python)**.)

This app is strongly inspired by **[ifconfig.me](http://ifconfig.me)**

## Requirements

* Golang 1.24 or higher

## Getting started

Prebuilt binaries can be downloaded from the GitHub Releases [section](https://github.com/lescactus/whoami-go/releases), or using a Docker image from the Github Container Registry. See [here](https://github.com/lescactus/whoami-go#running-with-docker-rooster)

### Building `whoami-go` :cd:

<details>

#### From source with Go

You need a working [go](https://golang.org/doc/install) toolchain (It has been developped and tested with go 1.24 and should work with go >= 1.24). Refer to the official documentation for more information (or from your Linux/Mac/Windows distribution documentation to install it from your favorite package manager).

```bash
# Clone this repository
git clone https://github.com/lescactus/whoami-go.git && cd whoami-go/

# Build from sources. Use the '-o' flag to change the compiled binary name
go build

# Default compiled binary is whoami-go
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./whoami-go
```

#### From source with docker

If you don't have [go](https://golang.org/doc/install) installed but have docker, run the following command to build inside a docker container:

```bash
# Build from sources inside a docker container. Use the '-o' flag to change the compiled binary name
# Warning: the compiled binary belongs to root:root
docker run --rm -it -v "$PWD":/app -w /app golang:1.24 go build -buildvcs=false

# Default compiled binary is whoami-go
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./whoami-go
```

The server is accessible at http://127.0.0.1:8080

#### With Docker

`whoami-go` comes with a `Dockerfile`. To build the image:

```bash
docker build -t whoami-go .

docker run -d -p 8080:8080 --restart="always" --name whoami-go whoami-go 
```

The server is accessible at http://127.0.0.1:8080

</details>

### Running with Docker :rooster:

```bash
docker run -d -p 8080:8080 --restart="always" --name whoami-go ghcr.io/lescactus/whoami-go
```

The server is accessible at http://127.0.0.1:8080

### Running with Docker Compose :cactus:

```bash
docker compose up
```

The server is accessible at http://127.0.0.1:8080

## Specifications :ocean:

### Endpoints

* `GET /` display informations from the browser: http headers, public IP, geo locations informations (such as country name, city name, etc ...) and a map

* `GET /index` display informations from the browser: http headers, public IP, geo locations informations (such as country name, city name, etc ...) and a map

* `GET /ip` display the public ip of the browser

* `GET /port` display the remote port of the browser

* `GET /lang` display the language of the browser (`Accept-Language`)

* `GET /ua` display the user agent of the browser (`User-Agent`)

* `GET /raw/go` display all the http headers the browser send in a golang fashion format

* `GET /raw/json` display all the http headers the browser send in json

* `GET /raw/yaml` display all the http headers the browser send in yaml

* `GET /metrics` display prometheus metrics if enabled (see configuration below)


## Configuration :deciduous_tree:

`whoami-go` is a 12-factor app using [Viper](https://github.com/spf13/viper) as a configuration manager. It can read configuration either from environment variables or from a key/value `.env` file. This `.env` file represents system environment variables on the machine. This change was made with the ease-of-use with Docker in mind.

### Available variables

* `APP_ADDR`(default value: `:8080`)

* `APP_ENV`(default value: `local`). If set to anything other than `production`, run the `pprof` server at `/debug/pprof/`

* `APP_CONFIG_NAME` (default `.env`)

* `APP_CONFIG_PATH` (default `.`)

* `SERVER_FIBER_PREFORK`(default value: `false`)

* `SERVER_FIBER_READ_TIMEOUT`(default value: `10s`)

* `SERVER_FIBER_WRITE_TIMEOUT`(default value: `10s`)

* `SERVER_FIBER_IDLE_TIMEOUT`(default value: `75s`)

* `SERVER_FIBER_ENABLE_TRUSTED_PROXY_CHECK`(default value: `true`)

* `SERVER_FIBER_PROXY_HEADER`(default value: `X-Forwarded-For`)

* `SERVER_FIBER_TRUSTED_PROXIES`(default value: `169.254.8.129`)

* `SERVER_FIBER_DISABLE_KEEPALIVE`(default value: `false`)

* `VIEWS_TEMPLATE_DIRECTORY`(default value: `./views/template`)

* `VIEWS_TEMPLATE_EXTENSIONS`(default value: `.html`)

* `VIEWS_STATIC_DIRECTORY`(default value: `./views/static`)

* `MIDDLEWARE_RECOVER_ENABLE_STACK_TRACE`(default value: `true`)

* `HTTP_CLIENT_TIMEOUT` (default value: `5s`)

* `GEOLOCATION_API` (default value: `freegeoip`). Define which IP GeoLocation API to use (https://freegeoip.app/, https://ip-api.com/). Valid values are `freegeoip` or `ipapi`

* `LOGGER_TYPE` (default value: `gofiber`). Define which logger to use. Valid values are `gofiber` (default logger) or `zap`

* `LOGGER_ZAP_LOG_LEVEL` (default value: `info`). Only usable when `LOGGER_TYPE=zap`

* `LOGGER_ZAP_DEVELOPMENT_MODE` (default value: `false`). Only usable when `LOGGER_TYPE=zap`

* `LOGGER_ZAP_DISABLE_CALLER` (default value: `true`). Only usable when `LOGGER_TYPE=zap`

* `LOGGER_ZAP_DISABLE_STACK_TRACE` (default value: `true`). Only usable when `LOGGER_TYPE=zap`

* `LOGGER_ZAP_ENCODING` (default value: `json`). Valid values are: `json` or `console`. Only usable when `LOGGER_TYPE=zap`

* `MIDDLEWARE_PROMETHEUS_ENABLED` (default value: `true`)

* `MIDDLEWARE_PROMETHEUS_METRICS_PATH` (default value: `/metrics`) Only usable when `MIDDLEWARE_PROMETHEUS_ENABLED=true`

Screenshots
-----------
![IP info location](https://i.imgur.com/tAXLqaJ.png "IP info location")
***
![Map](https://i.imgur.com/o0j8NZj.png "Map")
***
![Sitemap](https://i.imgur.com/uemt9fm.png "Site map")
