FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the service files
COPY *.go ./

# Build the service
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/api-gateway .

EXPOSE 8090

CMD ["./api-gateway"]
