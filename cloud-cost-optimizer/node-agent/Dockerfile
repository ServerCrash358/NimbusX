FROM alpine:latest
RUN apk add --no-cache docker-cli ca-certificates
WORKDIR /root/
COPY node-agent .
RUN chmod +x node-agent
EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=3s --start-period=5s \
  CMD wget -qO- http://localhost:8080/health || exit 1

CMD ["./node-agent"]
