#FROM golang:1.21-alpine as builder

#WORKDIR /app
#COPY main.go .
#COPY go.mod .

#RUN go mod download
#RUN CGO_ENABLED=0 GOOS=linux go build -o metrics-app

#FROM alpine:3.18
#WORKDIR /app
#COPY --from=builder /app/metrics-app .

#EXPOSE 8080
#CMD ["./metrics-app"]

# Use a Go base image for building the binary
FROM golang:1.20 AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o metrics-app

# Use a minimal image for running the app
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/metrics-app .
CMD ["./metrics-app"]

