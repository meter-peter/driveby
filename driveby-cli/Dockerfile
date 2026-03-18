# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_DATE=unknown
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X github.com/meter-peter/driveby/driveby-cli/internal/cli.version=${VERSION} \
              -X github.com/meter-peter/driveby/driveby-cli/internal/cli.commit=${COMMIT} \
              -X github.com/meter-peter/driveby/driveby-cli/internal/cli.date=${BUILD_DATE}" \
    -o driveby ./cmd/driveby

FROM gcr.io/distroless/static
WORKDIR /
COPY --from=builder /app/driveby /driveby
ENTRYPOINT ["/driveby"]