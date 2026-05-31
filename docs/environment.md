# Environment Variables

Use local environment files for development settings. Do not commit `.env`, `.env.local`, or service account files.

## Root

`docker-compose.yml` reads these values when running both services together:

| Name | Default | Purpose |
| --- | --- | --- |
| `FRONTEND_PORT` | `3000` | Host port for the React development server |
| `BACKEND_PORT` | `8080` | Host port for the Go API |

## Frontend

The frontend uses Create React App, so browser-exposed variables must start with `REACT_APP_`.

| Name | Purpose |
| --- | --- |
| `REACT_APP_API_PROXY_TARGET` | Backend target for local `/api` proxying |
| `REACT_APP_FIREBASE_API_KEY` | Firebase web API key |
| `REACT_APP_FIREBASE_AUTH_DOMAIN` | Firebase auth domain |
| `REACT_APP_FIREBASE_PROJECT_ID` | Firebase project ID |
| `REACT_APP_FIREBASE_STORAGE_BUCKET` | Firebase Storage bucket |
| `REACT_APP_FIREBASE_MESSAGING_SENDER_ID` | Firebase sender ID |
| `REACT_APP_FIREBASE_APP_ID` | Firebase web app ID |
| `REACT_APP_FIREBASE_MEASUREMENT_ID` | Optional Analytics measurement ID |

## Backend

| Name | Purpose |
| --- | --- |
| `PORT` | HTTP port for the Go API |
| `CORS_ALLOWED_ORIGINS` | Comma-separated frontend origins allowed by the API |
| `FIREBASE_CREDENTIALS_FILE` | Path to a Firebase Admin service account JSON file |
| `FIREBASE_CREDENTIALS_JSON` | Firebase Admin service account JSON value |
