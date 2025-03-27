# Build Stage
FROM golang:1.23.1-alpine AS builder
WORKDIR /app

# Copy and download dependencies first (optimization for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production Stage
FROM alpine:3.20
WORKDIR /app

# Install only required runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy binary
COPY --from=builder /app/main .

# Copy .env if needed (optional)
COPY --from=builder /app/.env ./

# Use non-root user
RUN adduser -D -u 1000 appuser && chown appuser:appuser /app/main
USER appuser

EXPOSE 5003
CMD ["./main"]