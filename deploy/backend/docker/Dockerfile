# Stage 1: Build the Go app
FROM golang:latest AS gobuilder
# FROM golang:alpine AS gobuilder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN swag init -g cmd/gin-realtime/main.go -o doc/openapi
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/gin-realtime ./cmd/gin-realtime

# Stage 2: Create the final container
FROM alpine:latest

# Install ca-certificates
RUN apk update && apk add curl jq ca-certificates && rm -rf /var/cache/apk/*

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=gobuilder /app/bin/gin-realtime .
COPY --from=gobuilder /app/doc/openapi ./doc/openapi
COPY --from=gobuilder /app/.env .

# Expose port 3001 to the outside world
EXPOSE 3001

# Command to run the executable
ENTRYPOINT ["./gin-realtime"]
