# Use the official Golang image
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy files to the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


# Command to run the executable
RUN go build -o main main.go

FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/default.env .
COPY --from=builder /app/wait_for.sh .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/test.sh .

RUN apk update  # Update the package index
RUN apk add make  # Install make

# Create a directory for Migrate
RUN mkdir -p /usr/local/bin
# Download the Migrate binary
RUN wget -O /tmp/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz
# Unzip the Migrate binary
RUN tar -xzf /tmp/migrate.tar.gz -C /usr/local/bin/ && rm /tmp/migrate.tar.gz
# Make the binary executable
RUN chmod +x /usr/local/bin/migrate
