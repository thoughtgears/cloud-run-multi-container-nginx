# -- Auth Helper Build Stage --
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache upx=4.2.4-r0

ARG SRC_PATH

WORKDIR /go/src/github.com/${SRC_PATH}/proxy/auth_helper

COPY proxy/auth_helper/go.mod ./
COPY proxy/auth_helper/go.sum ./
COPY proxy/auth_helper/main.go ./
COPY proxy/auth_helper/handlers ./handlers
COPY proxy/auth_helper/config ./config
COPY proxy/auth_helper/services ./services

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o auth_helper_go . && \
     upx --best --lzma auth_helper_go

# -- Artifact Stage --
FROM nginx:1.28-alpine AS artifact

ARG SRC_PATH

# Install ca-certificates package
# This package provides the root CA certificates typically found in /etc/ssl/certs/
RUN apk --no-cache add ca-certificates=20241121-r1 supervisor=4.2.5-r5

RUN rm /etc/nginx/nginx.conf

COPY --from=builder /go/src/github.com/${SRC_PATH}/proxy/auth_helper/auth_helper_go /app/auth_helper_go
COPY proxy/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY proxy/nginx.conf /etc/nginx/nginx.conf
COPY proxy/nginx.proxy.conf /etc/nginx/conf.d/proxy.conf

EXPOSE 8080

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]