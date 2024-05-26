# Build stage
FROM golang:1.22.3-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Run stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built application from the builder
COPY --from=builder /app/main .
COPY --from=builder /app .

# Install Go tools
RUN apk add --no-cache git go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
