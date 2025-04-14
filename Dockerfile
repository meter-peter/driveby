# Build stage
FROM golang:1.19-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o driveby ./cmd/driveby

# Run stage
FROM alpine:3.16

# Set working directory
WORKDIR /app

# Add ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/driveby .

# Expose the application port
EXPOSE 8080

# Create config directory
RUN mkdir -p /etc/driveby

# Set environment variables
ENV DRIVEBY_SERVER_HOST=0.0.0.0 \
    DRIVEBY_SERVER_PORT=8080 \
    DRIVEBY_SERVER_MODE=release

# Run the application
CMD ["./driveby"]