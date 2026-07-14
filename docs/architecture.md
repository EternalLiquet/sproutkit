# Architecture

SproutKit is bootstrapped as a minimal monorepo:

- `apps/api`: Go + Chi REST API scaffold with a single `GET /health` endpoint.
- `apps/dashboard`: Next.js (TypeScript, App Router, Tailwind, shadcn/ui initialized) with a single placeholder page.
- `packages/sdk-typescript`: Buildable TypeScript package exporting an empty `SproutClient` class.
- `packages/api-contracts`: Minimal OpenAPI document for `GET /health`.
- `docker-compose.yml`: Local PostgreSQL only.
