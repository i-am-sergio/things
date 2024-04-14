# Build the binary
FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gateway

# Run the app
FROM alpine:3.19.1
WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
COPY --from=builder /app/.env /app/.env

EXPOSE 8888

CMD ["./gateway"]