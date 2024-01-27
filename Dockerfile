# Use the official Golang image based on Alpine
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o perene-golang-x .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./perene-golang-x"]
