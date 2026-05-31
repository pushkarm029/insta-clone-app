# Contributing

Thanks for taking a look at Insta-Clone-App.

## Local Setup

1. Install Go and Node.js.
2. Run `npm ci` from `frontend/`.
3. Run `go mod download` from `backend/`.
4. Create local environment files from the examples.

## Development Checks

Run these before opening a pull request:

```shell
cd frontend
CI=true npm test -- --watchAll=false
npm run build

cd ../backend
go test ./...
gofmt -w .
```

## Commit Style

Use short conventional-style subjects:

- `docs: update setup guide`
- `test: add search page coverage`
- `fix: handle missing Firebase credentials`
- `chore: tidy backend modules`

## Pull Requests

Keep pull requests focused. Include setup notes, screenshots for UI changes, and local verification output.
