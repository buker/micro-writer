FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/micro-writer

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Download grpc_health_probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.11 && \
wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
chmod +x /bin/grpc_health_probe

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod make init && go mod download 
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod make proto tidy build

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /bin/grpc_health_probe /bin/
COPY --from=builder /go/src/micro-writer/micro-writer /micro-writer
ENTRYPOINT ["/micro-writer"]
CMD []
