# Tasks

## Milestone 0: Product and architecture baseline
### Epic: Requirements and design
- [x] Write PRD and architecture overview.
- [x] Define local dev services (PostgreSQL, NATS).

### Epic: Decisions to unblock build
- [x] Pick rich-text editor: TipTap.
- [x] Pick collaboration model: Yjs CRDT with awareness for presence.
- [x] Pick migration tool for code-first schema changes: Atlas.

## Milestone 1: Repository and local dev setup
### Epic: Repo scaffold
- [x] Add service folders: `frontend/`, `services/document/`, `services/collab/`, `shared/`.
- [x] Add `.env.example` with service ports and DB settings.

### Epic: Local infrastructure
- [x] Add `docker-compose.yml` for PostgreSQL and NATS.
- [x] Add `make` or `task` shortcuts for `up`, `down`, `logs`.

## Milestone 2: Document service MVP
### Epic: Service skeleton
- [x] Initialize service module and config loading.
- [x] Add PostgreSQL connection and health endpoint.

### Epic: Data model and migrations
- [x] Define document schema in code (UUID, display name, Yjs content as BYTEA).
- [x] Generate code-first migrations and apply locally.
- [x] Add migration command in service startup or a dedicated CLI.

### Epic: REST API
- [x] `POST /documents` create document with optional `displayName`.
- [x] `GET /documents/{document_id}` fetch document state.
- [x] `GET /documents?query=` list and search by `displayName`.
- [x] Validate `document_id` and return 404 on missing docs.
- [x] Return list sorted by `updated_at` desc with limit/offset.

### Epic: NATS persistence
- [x] Subscribe to Yjs snapshot events.
- [x] Persist Yjs snapshots with debounced writes from clients.
- [ ] Track and log update versions for replay safety.

## Milestone 3: Collaboration service MVP
### Epic: WebSocket lifecycle
- [x] Accept connections with `document_id`.
- [x] Join/leave tracking per document session.
- [x] Broadcast updates to all session peers.

### Epic: NATS fan-out
- [x] Publish edit operations to NATS per document.
- [x] Subscribe to NATS and forward to local clients.
- [ ] Handle duplicate or out-of-order messages.

## Milestone 4: Frontend MVP
### Epic: App shell
- [x] Create React app with Home and Editor routes.
- [x] Home: list documents and search by `displayName`.
- [x] Home: create document and navigate to editor.

### Epic: Editor integration
- [x] Integrate rich-text editor and toolbar (bold/italic/etc).
- [x] Display `displayName` in header.
- [x] Load initial content from Document Service.
- [x] Initialize Yjs doc from persisted snapshot.

### Epic: Realtime sync
- [x] Open WebSocket session for document.
- [x] Apply remote updates into editor state.
- [x] Emit local edits with client id and version.
- [x] Show collaborator cursors and activity indicators.
- [x] Allow anonymous display name with random two-word fallback.

## Milestone 5: Reliability and testing
### Epic: Automated testing
- [ ] Unit tests for document handlers and validation.
- [ ] Integration tests with Postgres + NATS via docker compose.

### Epic: Observability
- [ ] Add structured logging with `document_id` context.
- [ ] Emit metrics for latency and WebSocket sessions.

### Epic: Hardening
- [ ] Add rate limits for create/join.
- [ ] Add graceful shutdown and reconnect behavior.
