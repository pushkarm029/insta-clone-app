# Troubleshooting

## Frontend Tests Do Not Exit

Run tests in one-shot mode:

```shell
cd frontend
CI=true npm test -- --watchAll=false
```

## Create React App Babel Warning

If the frontend prints a warning about `@babel/plugin-proposal-private-property-in-object`, install the package as a development dependency. This works around a known transitive dependency issue in Create React App.

## API Requests Fail Locally

Confirm the backend is running:

```shell
curl http://localhost:8080/api/search/users
```

Then confirm the frontend proxy target:

```shell
REACT_APP_API_PROXY_TARGET=http://localhost:8080 npm start
```

## Firebase Admin Fails To Initialize

The Go backend needs Firebase Admin credentials. Set either `FIREBASE_CREDENTIALS_FILE` or `FIREBASE_CREDENTIALS_JSON`.

## Docker Ports Are Already In Use

Change ports in `.env`:

```shell
FRONTEND_PORT=3001
BACKEND_PORT=8081
```
