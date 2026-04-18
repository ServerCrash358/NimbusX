# ── Build stage ──────────────────────────────────────────────────────────────
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY node-agent/ ./node-agent/
# Build without cgo for minimal static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /node-agent-bin ./node-agent/

# ── Runtime stage ─────────────────────────────────────────────────────────────
FROM alpine:3.19
# Note: In a real deployment, the node-agent needs access to the Docker daemon socket
# (e.g., via volume mount -v /var/run/docker.sock:/var/run/docker.sock)
RUN apk add --no-cache ca-certificates
COPY --from=builder /node-agent-bin /usr/local/bin/node-agent
EXPOSE 9091
CMD ["node-agent"]
