# Backend Dockerfile for development with hot reload

FROM golang:1.22-alpine

# Install system dependencies
RUN apk add --no-cache git make bash curl

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install Air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Expose API port
EXPOSE 8080

# Default command (will be overridden by docker-compose)
CMD ["air", "-c", ".air.toml"]
