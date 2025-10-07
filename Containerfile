# === Builder ===
FROM golang:1.23-bookworm AS builder

RUN apt-get update && \
    apt-get install -y curl gnupg ca-certificates && \
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g sass && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY web/ ./web/
COPY Makefile ./

RUN make build

# === Runtime ===
FROM alpine:3.20

RUN adduser -D -H appuser

WORKDIR /app

RUN mkdir -p /app/internal/database
COPY --from=builder /app/bin/app ./app
COPY --from=builder /app/internal/database/init.sql ./internal/database
COPY --from=builder /app/web/static ./web/static
COPY --from=builder /app/web/templates ./web/templates

RUN chmod +x ./app && chown -R appuser /app

USER appuser

EXPOSE 8080
ENTRYPOINT ["./app"]
