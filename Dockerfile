# Use Go official image as a builder
FROM golang:1.19 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o transaction-service

# Use a minimal base image for the final build
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built application from the builder
COPY --from=builder /app/transaction-service .

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./transaction-service"]
