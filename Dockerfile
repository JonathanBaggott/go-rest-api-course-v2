# First, we specify the base image for the build stage.
FROM golang:1.16 AS builder
# Create a directory called "app" inside the container.
RUN mkdir /app
# Copy the entire current directory (including the source code) into the "/app" directory in the container.
ADD . /app
# Set the working directory to "/app" inside the container.
WORKDIR /app
# Build the Go application inside the container.
# CGO_ENABLED=0 disables CGO (C library interface) to ensure a pure Go build.
# GOOS=linux specifies the target operating system as Linux.
# The compiled binary is named "app" and is located in the "cmd/server" directory relative to the current working directory.
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Next, we specify the base image for the production stage.
FROM alpine:latest AS production
# Copy the built application from the previous build stage (builder) into the current stage.
COPY --from=builder /app .
# Set the default command to run the application when the container starts.
CMD ["./app"]