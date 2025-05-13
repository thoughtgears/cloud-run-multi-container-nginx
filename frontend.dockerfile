FROM node:lts-alpine AS builder

WORKDIR /app

ARG MODE=development

COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npx tsc -b && \
    npx vite build --mode $MODE

FROM nginx:stable-alpine AS artifact

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/dist .

EXPOSE 5000

RUN rm /etc/nginx/nginx.conf
COPY ./frontend/nginx.conf /etc/nginx/nginx.conf
COPY ./frontend/nginx.frontend.conf /etc/nginx/conf.d/frontend.conf