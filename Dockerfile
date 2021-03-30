FROM golang:1.14 as builder
WORKDIR /usr/local/src/
COPY . ./
RUN CGO_ENABLED=0 go build -o ./helloworld-server ./server
# download grpc_health_probe for health check
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.6 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
# creating a small image
FROM gcr.io/distroless/static
COPY --from=builder /usr/local/src/helloworld-server /helloworld-server
COPY --from=builder /bin/grpc_health_probe ./grpc_health_probe
CMD ["/helloworld-server"]
