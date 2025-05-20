# Build stage
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies and generate go.sum
RUN go mod download && \
    go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /app

# Install curl for healthcheck
RUN apk add --no-cache curl

# Copy the binary from builder
COPY --from=builder /app/api .

# Expose the application port
EXPOSE 8081

# Create config directory
RUN mkdir -p /etc/driveby

# Set environment variables
ENV DRIVEBY_SERVER_HOST=0.0.0.0 \
    DRIVEBY_SERVER_PORT=8081 \
    DRIVEBY_SERVER_MODE=release

# Run the application
CMD ["./api"]