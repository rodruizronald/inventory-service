# First Stage: Build the Golang App

# Use official Golang image
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app for the correct platform
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Second Stage: Create a Small Production Image
FROM alpine:latest

# Set working directory in container
WORKDIR /root/

# Copy the built Go app from previous stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./main"]