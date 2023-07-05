# Use an official Golang runtime as the base image
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Use a minimal Alpine image as the base image
FROM alpine:latest

# Install MySQL client and required libraries
RUN apk add --no-cache mysql-client

# Set the working directory
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/app .

# Set the entry point to run the Go application
ENTRYPOINT ["./app"]
