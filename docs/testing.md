# Testing Guide

The project has frontend Jest tests and backend Go package checks.

## Frontend

Run all frontend tests once:

```shell
cd frontend
CI=true npm test -- --watchAll=false
```

Run tests with coverage:

```shell
cd frontend
npm test -- --coverage --watchAll=false
```

Build the frontend:

```shell
cd frontend
npm run build
```

## Backend

Run backend tests:

```shell
cd backend
go test ./...
```

Format backend code:

```shell
cd backend
gofmt -w .
```

## Maintenance Expectation

New maintenance work should either add tests or explain why the change is documentation-only.
