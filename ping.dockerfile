# ---- Base Stage ----
FROM node:22-alpine AS base
WORKDIR /usr/src/app
COPY ./apis/ping/package*.json ./

# ---- Dependencies Stage ----
FROM base AS dependencies
RUN npm install --only=production --no-package-lock


# ---- Build/Source Stage ----
# Copy the rest of your application's source code.
FROM base AS source_builder
COPY ./apis/ping/src/ ./src/

# ---- Production Stage ----
FROM node:22-alpine AS artifact
WORKDIR /usr/src/app
ENV NODE_ENV=production

COPY --from=dependencies /usr/src/app/node_modules ./node_modules
COPY --from=source_builder /usr/src/app/src/ ./src
COPY --from=source_builder /usr/src/app/package.json ./package.json

EXPOSE 8080

CMD ["node", "src/index.js"]