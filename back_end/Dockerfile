# Use an official Go runtime as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy your Go application source code into the container
COPY . /app

# Copy Go package dependencies
COPY ../go.* ./

# Build your Go application
RUN go build -o server

# Specify the command to run your backend application
CMD ["./server"]