# Build the binary
FROM golang:1.21.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o executable

# Run the app
FROM alpine:3.19.1
WORKDIR /app
COPY --from=builder /app/executable /app/executable
COPY --from=builder /app/.env /app/.env

EXPOSE 8002

CMD ["./executable"]
