# Stage 1: Build application
FROM golang:1.23 AS builder

# Set GOPROXY for better dependency resolution
ENV GOPROXY=https://proxy.golang.org,direct

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download all dependencies. Dependencies are cached if they have not changed
RUN go mod download

# Copy source code into container
COPY . .

# Build the Go library
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Test that the binary was created successfully
RUN if [ ! -f ./main ]; then echo "main binary not built"; exit 1; fi
RUN chmod +x ./main

# Stage 2: Run application
FROM alpine:latest

# Install necessary utilities (curl) and add dockerize
RUN apk add --no-cache curl \
    && curl -sSL https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar -xz -C /usr/local/bin

# Verify dockerize installation
RUN if [ ! -f /usr/local/bin/dockerize ]; then echo "dockerize installation failed"; exit 1; fi

# Set environment variables
ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASSWORD=secret
ENV DB_NAME=items

# Set the current working directory to root
WORKDIR /root/

# Copy the prebuilt binary file from the builder stage
COPY --from=builder /app/main .

# Verify the binary file exists
RUN if [ ! -x ./main]; then echo "main binary is missing or not executable"; exit 1; fi

# Expose the port the app will run on
EXPOSE 8080

# Command to run exectutable
CMD [ "/usr/local/bin/dockerize", "-wait", "tcp://db:5432", "-timeout", "60s", "./main" ]