# Use an official Go runtime as the base image
FROM golang:1.21.2

# Set the working directory inside the container
WORKDIR /app

# Pre-copy/cache the go.mod and go.sum files to avoid re-downloading dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy your Go application source code into the container (at /app)
COPY . .

ENV PORT=8080

# Build your Go application
RUN go build -o server

# Specify the command to run your backend application
CMD ["./server"]