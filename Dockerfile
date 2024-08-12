# Use an official Go image as the base image
FROM golang:1.22.6-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o task-manager ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./task-manager"]
