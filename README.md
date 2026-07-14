# SproutKit

Minimal monorepo scaffold for SproutKit.

## Prerequisites

- Go 1.24+
- Node.js 22+
- pnpm 10+
- Docker (for local PostgreSQL via Docker Compose)

## Local setup

1. Copy environment defaults:
   - `cp .env.example .env`
2. Install JavaScript dependencies:
   - `make install`
3. Start local PostgreSQL (optional for now):
   - `docker compose up -d`

## Available commands

- `make build` — build API, dashboard, and SDK
- `make test` — run API tests and dashboard typecheck
- `make api-build` — build Go API
- `make api-test` — test Go API
- `make dashboard-build` — build Next.js dashboard
- `make dashboard-typecheck` — typecheck Next.js dashboard
- `make sdk-build` — build TypeScript SDK

## Project layout

- `apps/api` — Go API (`GET /health`)
- `apps/dashboard` — Next.js dashboard placeholder
- `packages/sdk-typescript` — TypeScript SDK package scaffold
- `packages/api-contracts` — minimal OpenAPI contract
