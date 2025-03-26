# Start from the official Go image
FROM golang:1.22 AS build

WORKDIR /app
COPY . .

RUN go build -o api-gateway ./cmd/main.go

# Final image
FROM debian:buster
COPY --from=build /app/api-gateway /api-gateway
ENTRYPOINT ["/api-gateway"]
