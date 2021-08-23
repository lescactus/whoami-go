Who am I ?
==================

A tiny Go web app that display sample information from the visitor (IP, location, browser HTTP headers ...): **[whoami](https://whoami-go.alexasr.tk/)** (Development version: [dev](https://whoami-go-dev-ujargjwndq-ew.a.run.app/)).
It is a rewriting of one of my previous mini project (**[whoami-python](https://github.com/lescactus/whoami-python)**.)

This app is strongly inspired by **[ifconfig.me](http://ifconfig.me)**

#### Specifications

##### Endpoints

* `GET /` display informations from the browser: http headers, public IP, geo locations informations (such as country name, city name, etc ...) and a map

* `GET /index` display informations from the browser: http headers, public IP, geo locations informations (such as country name, city name, etc ...) and a map

* `GET /ip` display the public ip of the browser

* `GET /port` display the remote port of the browser

* `GET /lang` display the language of the browser (`Accept-Language`)

* `GET /ua` display the user agent of the browser (`User-Agent`)

* `GET /raw/go` display all the http headers the browser send in a golang fashion format

* `GET /raw/json` display all the http headers the browser send in json

* `GET /raw/yaml` display all the http headers the browser send in yaml




Use it now
----------

```sh
# Install via go get and run it
$ go get "github.com/lescactus/whoami-go"
$ whoami-go
...

# Git clone and build:
$ git clone https://github.com/lescactus/whoami-go.git
$ go run main.go
...
```

Now point your browser at http://127.0.0.1:8080

### Docker
**whoami** can easily be dockerized and is shipped with a ``Dockerfile``.

By default, the container will expose port 8080, so change this within the ``Dockerfile`` if necessary. When ready, simply use the ``Dockerfile`` to build the image.

```sh
docker build -t whoami .
```
This will create the Docker image.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 80 of the host to port 8080 of the container:

```sh
docker run -d -p 80:8080 --restart="always" --name whoami whoami 
```

Now point your browser at http://127.0.0.1/

### Configuration

`whoami` is a 12-factor app using [Viper](https://github.com/spf13/viper) as a configuration manager. It can read configuration either from environment variables or from a key/value `.env` file. This `.env` file represents system environment variables on the machine. This change was made with the ease-of-use with Docker in mind.

#### Available variables

* `APP_ADDR`(default value: `:8080`)

* `APP_ENV`(default value: `local`)

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


Screenshots
-----------
![IP info location](https://i.imgur.com/tAXLqaJ.png "IP info location")
***
![Map](https://i.imgur.com/o0j8NZj.png "Map")
***
![Sitemap](https://i.imgur.com/uemt9fm.png "Site map")
