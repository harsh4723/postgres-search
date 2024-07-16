# Start with the official Golang image
FROM golang:1.20

# Set the current working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
