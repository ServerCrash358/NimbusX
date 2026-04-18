# ── Build stage ──────────────────────────────────────────────────────────────
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY api/ ./api/
RUN go build -o /api-bin ./api/

# ── Runtime stage ─────────────────────────────────────────────────────────────
FROM alpine:3.19
COPY --from=builder /api-bin /usr/local/bin/api
EXPOSE 8080
CMD ["api"]
