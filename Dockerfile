# Use the official Go image as the base image for building
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod ./
# Download dependencies
RUN go mod download

# Copy the entire project source code
COPY . .

# Build the Go application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o tic-tac-toe main.go

# Use a minimal Alpine image for the final runtime
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/tic-tac-toe .
# Copy static files and templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expose port 8080 for the server
EXPOSE 8080

# Command to run the application
CMD ["./tic-tac-toe"]