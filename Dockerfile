# Use the official Golang image
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy files to the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# set swagger docs
RUN swag init

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
