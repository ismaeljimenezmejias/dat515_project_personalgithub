# Build stage
FROM golang:1.25-alpine AS builder

# Install git (sometimes needed for Go modules)
RUN apk add --no-cache git

# TODO : Set working directory to /app

# TODO : Copy go mod and go sum* files to the working directory

# Download dependencies
RUN go mod download

# TODO : Copy all source files from the current directory to container´s working directory

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 appgroup && \
    adduser -D -u 1001 -G appgroup appuser

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# TODO : Change ownership of binary to user (appuser)

# TODO : Switch to non-root user (appuser)

# Set environment variables
ENV APP_VERSION=1.0
ENV PORT=8080

# TODO : Expose port

# TODO : Health check

# TODO : Run the binary
