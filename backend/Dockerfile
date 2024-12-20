# Stage 1: Build the Go application
FROM golang:1.21.3-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

# Stage 2: Create a smaller final image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/myapp /myapp

# Define the entry point for the container
ENTRYPOINT ["/myapp"]

# Expose port (if applicable)
# EXPOSE 8080

# Set default command (if applicable)
# CMD ["your-default-command"]

