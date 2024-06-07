# Use the official Golang image as a build stage
FROM golang:1.20.14-alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git build-base
RUN apk add -U --no-cache ca-certificates && update-ca-certificates

# Create and set working directory
RUN mkdir /server
WORKDIR /server

# Copy go.mod and go.sum and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY ./ ./

# Build the Go application for Linux ARM64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o build/grpc-demo

# Use scratch image as the base for the final image
FROM scratch
LABEL MAINTAINER="Akshit Verma"
LABEL VERSION="0.0.1"

# Copy the binary and certificates from the builder stage
COPY --from=builder /server/build/grpc-demo /go/bin/grpc-demo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set the entrypoint to the binary
ENTRYPOINT ["/go/bin/grpc-demo"]
