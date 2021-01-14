# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main


#Second stage build
FROM alpine:latest

WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main .

# This container exposes port 80 to the outside world
EXPOSE 9000

# Run the binary program produced by `go install`
ENTRYPOINT ["./main"] --port 9000