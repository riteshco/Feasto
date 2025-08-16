# ---- Stage 1: Build Go App & Get Tools ----
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git curl

# Download and install the golang-migrate CLI tool
RUN curl -o migrate.tar.gz -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.tar.gz && \
    mv ./migrate /usr/local/bin/migrate && \
    rm migrate.tar.gz

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

# ---- Stage 2: Final Production Image ----
FROM alpine:latest

RUN apk add --no-cache netcat-openbsd

COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

WORKDIR /app

# Copy the compiled application binary from the builder stage
COPY --from=builder /server .

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY ./database/migration /migrations

ENTRYPOINT ["/entrypoint.sh"]

CMD ["./server"]