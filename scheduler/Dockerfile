# ── Build stage ──────────────────────────────────────────────────────────────
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY scheduler/ ./scheduler/
RUN go build -o /scheduler-bin ./scheduler/

# ── Runtime stage ─────────────────────────────────────────────────────────────
FROM alpine:3.19
RUN apk add --no-cache wget
COPY --from=builder /scheduler-bin /usr/local/bin/scheduler
EXPOSE 9090
CMD ["scheduler"]
