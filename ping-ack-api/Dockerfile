# Use an official Go runtime as a base image
FROM golang:1.21.5

# Set the working directory in the container
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main
ENTRYPOINT ["/app/main"]

