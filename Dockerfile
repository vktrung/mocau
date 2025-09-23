# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Tidy dependencies
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create app directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Create static directory and copy files if they exist
RUN mkdir -p ./static
COPY --from=builder /app/static ./static

# Copy docs if any
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 3000

# Set environment variables
ENV GIN_MODE=release
ENV PORT=3000

# Run the application
CMD ["./main"]
