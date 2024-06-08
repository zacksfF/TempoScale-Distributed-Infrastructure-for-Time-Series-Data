#Build image 
# Use the latest Go version as the builder stage
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /API

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . /API

# Run command as described:
# go build will build a 64bit Linux executable binary file named server in the current directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temposcale .

# Use a minimal base image for the final stage
FROM alpine:latest

# Install certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Copy required binary executable into this image.
COPY --from=build-env /API/temposcale .

EXPOSE 8000

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /API/main .

# Command to run the binary
CMD ["./main"]
