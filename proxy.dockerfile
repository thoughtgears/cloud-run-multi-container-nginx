FROM nginx:1.28-alpine AS artifact

# Install ca-certificates package
# This package provides the root CA certificates typically found in /etc/ssl/certs/
RUN apk --no-cache add ca-certificates=20241121-r1

RUN rm /etc/nginx/nginx.conf
COPY proxy/nginx.conf /etc/nginx/nginx.conf
COPY proxy/nginx.proxy.conf /etc/nginx/conf.d/proxy.conf

EXPOSE 8080