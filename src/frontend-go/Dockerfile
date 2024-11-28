# Stage 1: Build the application
FROM golang:1.18-bullseye AS base
# Set CGO_ENABLED=0 for a fully static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .
# Build the binary with optimizations
RUN go build -o main .

# Stage 2: Create the smallest possible runtime container
FROM gcr.io/distroless/static-debian12 AS runtime
# Set the working directory
WORKDIR /

# Copy the statically compiled binary from the builder stage
COPY --from=base /app/main .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["/main"]
