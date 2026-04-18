# ── Build stage ──────────────────────────────────────────────────────────────
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY oracle/ ./oracle/
RUN go build -o /oracle-bin ./oracle/

# ── Runtime stage ─────────────────────────────────────────────────────────────
FROM alpine:3.19
RUN apk add --no-cache wget
COPY --from=builder /oracle-bin /usr/local/bin/oracle
EXPOSE 7070
CMD ["oracle"]
