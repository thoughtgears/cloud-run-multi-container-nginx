FROM nginx:1.28-alpine AS artifact

RUN rm /etc/nginx/nginx.conf
COPY proxy/nginx.conf /etc/nginx/nginx.conf
COPY proxy/nginx.proxy.conf /etc/nginx/conf.d/proxy.conf

EXPOSE 8080