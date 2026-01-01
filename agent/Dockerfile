# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /build
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apc-agent ./cmd/apc-agent

# Runtime stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /build/apc-agent .

# Default to cloud mode (no cursor movement), allow insecure HTTP
ENV APC_BIND=0.0.0.0:8787
ENV APC_ALLOW_INSECURE=true
ENV APC_CLOUD_MODE=true

EXPOSE 8787

CMD ["./apc-agent"]
