# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Doclet is an anonymous, real-time collaborative rich-text editor. Users create documents (assigned a UUID), share the ID, and collaborate instantly with live presence indicators. No authentication required by design.

## Architecture

Three services communicate through NATS message broker:

- **Document Service** (Go, port 8080) — REST API for document CRUD. Persists metadata and Yjs content snapshots to PostgreSQL. Consumes NATS snapshot events to update stored content.
- **Collaboration Service** (Go, port 8090) — Stateless WebSocket hub. Accepts connections at `/ws?document_id=X&client_id=Y`, broadcasts Yjs edits and presence updates to peers, and publishes/subscribes via NATS for cross-instance delivery.
- **Frontend** (React + Vite, port 5173) — Two-page SPA (Home + Editor). Uses TipTap editor with Yjs CRDT for conflict-free editing and y-protocols/awareness for cursor presence. Custom `DocletProvider` manages the WebSocket-to-Yjs bridge.

**Data flow**: User edits → TipTap generates Yjs update → base64-encoded over WebSocket → Collab Service broadcasts to peers + publishes to NATS → debounced snapshots (~1.5s) published to NATS → Document Service persists snapshot to Postgres.

**NATS subjects**: `doclet.documents.{docId}.updates`, `doclet.documents.{docId}.presence`, `doclet.documents.{docId}.snapshots`

**Database**: Single `documents` table (document_id UUID PK, display_name TEXT, content BYTEA for Yjs state, created_at, updated_at). Code-first migrations via Gorm models + Atlas CLI.

## Build & Development Commands

### Start infrastructure (Postgres + NATS)
```
make up                    # docker compose up -d
make down                  # stop and reset local data (removes volumes)
make logs                  # follow docker compose logs
```

### Run services locally
```
make run-document          # Document Service on :8080 (sets default env vars)
make run-collab            # Collab Service on :8090
make run-frontend          # cd frontend && npm install && npm run dev
```

### Hot reload with Air
```
# Install: go install github.com/air-verse/air@latest
export $(grep -v '^#' .env | xargs)
air -c services/document/.air.toml   # document service
air -c services/collab/.air.toml     # collab service
```

### Full-stack Docker
```
docker compose --profile app up -d   # all services + infra
```

### Build
```
go build ./cmd/document
go build ./cmd/collab
cd frontend && npm run build         # outputs to frontend/dist/
```

### Test
```
go test ./...                        # Go tests
```

### Generate database migration
```
go run ./services/document/cmd/atlas > services/document/migrations/<timestamp>_<name>.sql
```

## Key Environment Variables

See `.env.example`. Core variables:
- `DOCLET_DOCUMENT_ADDR` — Document service listen address (default `:8080`)
- `DOCLET_COLLAB_ADDR` — Collab service listen address (default `:8090`)
- `DOCLET_DATABASE_URL` — Postgres connection string
- `DOCLET_NATS_URL` — NATS connection string
- `VITE_DOC_SERVICE_URL` / `VITE_COLLAB_WS_URL` — Frontend build-time service URLs

Frontend also supports runtime config via `/config.json` (see `frontend/config.example.json`).

## Tech Stack

- **Go 1.24** — Backend services (chi router, Gorm ORM, gorilla/websocket, nats.go)
- **React 18 + TypeScript 5.4** — Frontend (TipTap 2.2, Yjs 13.6, Tailwind 3.4, Vite 5.2)
- **PostgreSQL 16** — Document persistence
- **NATS 2.10** — Inter-service messaging (pub/sub)
- **Atlas** — Code-first database migrations from Gorm models

## Document Service API

- `POST /documents` — Create document (optional `displayName` in body)
- `GET /documents?q=&limit=&offset=` — List/search documents
- `GET /documents/{id}` — Get document with content
- `PUT /documents/{id}/title` — Update display name
- `DELETE /documents/{id}` — Delete document
- `GET /healthz` — Health check

## Code Layout

```
cmd/document/main.go          # Document service entrypoint
cmd/collab/main.go             # Collab service entrypoint
services/document/             # Document service (server, store, models, nats, config, database, migrate)
services/collab/               # Collab service (server, hub, nats, config)
frontend/src/pages/            # HomePage.tsx, EditorPage.tsx
frontend/src/editor/           # DocletProvider.ts (WebSocket + Yjs provider)
frontend/src/api.ts            # Document service API client
frontend/src/config.ts         # Runtime config loader
```

## Docker Compose Profiles

- Default (no profile): starts only `postgres` and `nats` (infrastructure)
- `--profile app`: starts all services including document-svc, collab-svc, and frontend
