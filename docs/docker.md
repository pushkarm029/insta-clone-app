# Docker Usage

The repository includes Docker files for local development and deployment experiments.

## Compose

From the repository root:

```shell
cp .env.example .env
docker compose up --build
```

The frontend is available at `http://localhost:3000` and proxies API calls to the backend service. The backend listens on `http://localhost:8080`.

## Frontend Image

The frontend image is based on Node and runs the Create React App development server. It is useful for local integration testing, not as an optimized static production image.

## Backend Image

The backend image builds the Go API and starts it inside the container. Runtime Firebase Admin credentials should be passed through environment variables or mounted secret files.

## Cleanup

```shell
docker compose down
docker compose down --volumes
```
