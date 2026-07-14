.PHONY: install build test api-build api-test dashboard-build dashboard-typecheck sdk-build

install:
	pnpm install --ignore-scripts

build: api-build dashboard-build sdk-build

test: api-test dashboard-typecheck

api-build:
	cd apps/api && go build ./...

api-test:
	cd apps/api && go test ./...

dashboard-build:
	cd apps/dashboard && pnpm build

dashboard-typecheck:
	cd apps/dashboard && pnpm typecheck

sdk-build:
	cd packages/sdk-typescript && pnpm build
