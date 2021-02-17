FROM golang:alpine as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY *.go .

RUN CGO_ENABLED=0 go build -ldflags '-d -w -s' -o main

FROM alpine:3

WORKDIR /app

RUN chown -R 65534:65534 /app

COPY --from=builder --chown=65534:65534 /app/main /app
COPY --chown=65534:65534 ./templates /app/templates
COPY --chown=65534:65534 ./static /app/static

EXPOSE 8080

# nobody
USER 65534

CMD ["./main"]
