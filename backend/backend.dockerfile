# Use the official Golang image
FROM golang:1.24

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Install Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]