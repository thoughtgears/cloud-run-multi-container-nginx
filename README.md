# Cloud Run Multi-Container NGINX Example

This project demonstrates a multi-container application using Go (Gin), React, and NGINX, orchestrated with Docker Compose.
The main POC will be to run NGINX as a API proxy in Cloud Run to support the Frontend with the Backend API.

## Project Structure

* `apis/users`: Go service for user management (Gin)
* `apis/ping`: Go service for health checks (Gin)
* `frontend`: React app served by NGINX
* `proxy`: NGINX reverse proxy for API routing

## Services

| Service  | Description                         | Port |
|----------|-------------------------------------|------|
| frontend | React app via NGINX                 | 3000 |
| proxy    | NGINX reverse proxy for API routing | 8080 |
| users    | Go user API (Gin)                   | N/A  |
| ping     | Go ping API (Gin)                   | N/A  |

## Development

### Prerequisites

* Docker & Docker Compose
* Node.js (for local frontend dev)
* Go (for local API dev)

### Running with Docker Compose

`docker compose up --build`

* Frontend: http://localhost:3000
* API Gateway: http://localhost:8080

### API Endpoints

* GET /ping — returns pong
* GET /users — user management endpoints (see apis/users)

### Environment Variables

Frontend uses VITE_API_BASE_URL (see frontend/.env.development.local)

### Building and Pushing Docker Images

This is the required `.env` file for building and pushing the images:

```dotenv
DOCKER_REPO=my-full-repo-path
GCP_REGION=service-region
GCP_PROJECT_ID=service-project-id
USERS_SERVICE=user-service-run.app# for preparing the proxy config
PING_SERVICE=ping-service-run.app# for preparing the proxy config
```

```shell
task build:all # build and push all images
task deploy:all # deploy all images to Cloud Run
```