FROM golang:1.15-alpine as builder

RUN apk add git \
  && go get -v "github.com/ghodss/yaml"

ADD . /go/src/whoami

WORKDIR /go/src/whoami

RUN go build -o main

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*

WORKDIR /app
COPY --from=builder /go/src/whoami/main /app
COPY --from=builder /go/src/whoami/templates /app/templates
COPY --from=builder /go/src/whoami/static /app/static

EXPOSE 8080

CMD ["./main"]
