# Base builder stage for shared dependencies
FROM golang:1.24-alpine AS base

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Development stage with hot-reloading using Air
FROM base AS development

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# Copy the source code
COPY . .

# Default to development mode (will be overridden in production)
ENV GO_ENV=development

# Expose port
EXPOSE 8080

# Use Air for live-reloading during development
CMD ["air", "-c", ".air.toml"]

# Build stage for production
FROM base AS builder

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api

# Final production stage
FROM alpine:latest AS production

# Set working directory
WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/api .
COPY --from=builder /app/config ./config

# Set environment to production
ENV GO_ENV=production

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./api"] 