# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o driveby ./cmd/driveby

FROM gcr.io/distroless/static
WORKDIR /
COPY --from=builder /app/driveby /driveby
ENTRYPOINT ["/driveby"] 