# Stage 1: Build the application
FROM golang:1.22.2-alpine AS builder

WORKDIR /build

# Copy go mod and sum files
COPY app/go.mod app/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY app/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /build/main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
