FROM golang:1.23-bookworm

# Install Sass
RUN apt-get update && \
    apt-get install -y curl gnupg ca-certificates && \
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g sass && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY web/ ./web/

# Build the Go app
RUN go build -o bin/app ./cmd/app/main.go

# Run the app
CMD ["./bin/app"]
