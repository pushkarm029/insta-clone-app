#!/usr/bin/env sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)

echo "Running frontend tests"
cd "$ROOT_DIR/frontend"
CI=true npm run test:ci -- --runInBand

echo "Building frontend"
npm run build

echo "Running backend tests"
cd "$ROOT_DIR/backend"
go test ./...

echo "Checking backend formatting"
test -z "$(gofmt -l .)"

echo "Local validation complete"
