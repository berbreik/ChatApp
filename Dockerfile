# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build deps
RUN apk add --no-cache git ca-certificates

# Cache modules first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build statically (smaller image, no CGO)
ENV CGO_ENABLED=0
RUN go build -o chatapp ./cmd/api

# ---- Runtime stage ----
FROM alpine:3.19

WORKDIR /app

# Add CA certs for HTTPS calls
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/chatapp /app/chatapp

# Expose Gin port
EXPOSE 8080

# Environment variables (override in runtime)
# JWT secret for your app auth
ENV JWT_SECRET=change-me
# Stream Chat credentials
ENV STREAM_API_KEY=your-stream-api-key
ENV STREAM_API_SECRET=your-stream-api-secret
# Database connection (example for Postgres)
ENV DATABASE_URL=postgres://user:pass@host:5432/chatapp?sslmode=disable

# Optional: healthcheck (expects /health endpoint)
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

# Run the server
CMD ["/app/chatapp"]