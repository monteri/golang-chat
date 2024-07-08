FROM golang:1.22

# Set the current working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application on port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]
