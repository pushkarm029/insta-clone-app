.PHONY: frontend-install frontend-test frontend-build backend-test backend-format validate

frontend-install:
	cd frontend && npm ci

frontend-test:
	cd frontend && CI=true npm run test:ci -- --runInBand

frontend-build:
	cd frontend && npm run build

backend-test:
	cd backend && go test ./...

backend-format:
	cd backend && gofmt -w .

validate:
	./scripts/validate-local.sh
