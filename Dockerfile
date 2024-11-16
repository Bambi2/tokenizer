# Start from a small, secure base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY ./deployment/go.mod ./deployment/go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY ./deployment .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/g42 ./cmd/main.go

# Create a minimal production image
FROM alpine:3.20

# It's essential to regularly update the packages within the image to include security patches
RUN apk update && apk upgrade

# Reduce image size
RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# Avoid running code as a root user
RUN adduser -D appuser
USER appuser

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/bin/g42 .

# Run the binary when the container starts
CMD ["./g42"]