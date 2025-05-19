FROM golang:1.23 AS builder

LABEL author="Nazar Parnosov"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apt-get update && apt-get install -y gcc libc6-dev libsqlite3-dev

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN /go/bin/swag init -g cmd/WeatherSubscriptionAPI/main.go


RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o weather-app ./cmd/WeatherSubscriptionAPI

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
  ca-certificates \
  libsqlite3-0 \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/weather-app .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/web ./web
COPY --from=builder /app/docs ./docs

COPY .env .env

EXPOSE 8080

CMD ["./weather-app"]
