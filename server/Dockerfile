# Use an official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the server source code into the container
COPY . .

# Build the Go server binary
RUN go build -o server server.go handler.go

# Expose the gRPC port
EXPOSE 6666

# Command to run the server binary
CMD ["./server", "-ip", "0.0.0.0", "-port", "6666"]