FROM golang:1.23-alpine

WORKDIR /app

# Install git for fetching dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files (if they exist)
COPY go.mod go.sum* ./

# Install dependencies
RUN go mod download

# Copy the source code
COPY main.go ./

# Build the application
RUN go build -o loadtester main.go

# Environment variables with defaults
ENV API_HOST=api \
    API_PORT=8080 \
    API_BASE_PATH="" \
    OPENAPI_PATH="/swagger/doc.json" \
    REQUEST_RATE=50 \
    TEST_DURATION=30 \
    REQUEST_TIMEOUT=5

# Run the load test
CMD ["./loadtester"]