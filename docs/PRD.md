# Product Requirements Document

## Product Overview
Doclet is a lightweight, real-time collaborative rich-text editor. Users create a document, share its `document_id`, and collaborate instantly without authentication. The product prioritizes simplicity, speed, and reliability while remaining easy to deploy and scale.

## Goals
- Enable real-time collaborative editing with rich-text formatting.
- Keep onboarding friction low (no accounts).
- Persist documents reliably so sessions can resume later.
- Support horizontal scaling with low-latency sync.
- Keep the UI limited to Home and Editor pages.

## Non-Goals
- User accounts, permissions, or access control.
- Advanced document management (folders, search, versioning).
- Offline-first or mobile-native clients in the MVP.

## Target Users
- Small teams or individuals needing quick collaboration.
- Educators, PMs, or engineers drafting and editing shared notes.

## Key Features
### 1) Document Creation
- Create a new document with a generated `document_id` (UUID).
- Optional `displayName` shown in the UI (e.g., "Project Plan v1").
- Share `document_id` to invite collaborators.

### 2) Real-Time Collaborative Editing
- Join by entering an existing `document_id`.
- Edits are synchronized in near real-time (<100ms target).
- Rich-text features: bold, italic, underline, headings, lists, links.
- Live presence: cursor positions and user activity are visible to all collaborators.
- Each browser session is an anonymous user with a generated session id.
- Users can optionally set a display name; otherwise a random two-word name is assigned.

### 3) Document Persistence
- Store `document_id`, `displayName`, and `content` in PostgreSQL.
- Last saved state is preserved when all users disconnect.
- Document can be reloaded on next session.

### 4) Real-Time Updates via Messaging
- WebSockets handle live collaboration sessions.
- NATS broadcasts updates across service replicas for scale and resilience.

### 5) Frontend Simplicity
- Two pages only:
  - Home: document list with search and create.
  - Editor: display name, rich-text editor, live presence.
- Editor uses TipTap with Yjs collaboration.
 - Cursor tooltips show collaborator names without taking a full row.

## User Flows
### Home and Search
1. User lands on Home and sees a document list sorted by most recent update.
2. User searches by `displayName`.
3. User selects a document to open or creates a new one.

### Create Document
1. User clicks "Create Document".
2. Backend creates record and returns `document_id` and initial content.
3. Editor opens and shows `displayName`.

### Join Document
1. User enters `document_id`.
2. Client fetches document metadata and content.
3. Client opens WebSocket for live updates.

### Real-Time Collaboration
1. Client emits local edits over WebSocket.
2. Collaboration service publishes change to NATS.
3. Other clients receive and apply the change instantly.

### Save Document State
- Document service consumes NATS events and persists content updates.
- Latest state remains available after users disconnect.

## Functional Requirements
- Create and fetch documents via REST endpoints.
- List documents by `updated_at` desc, with optional `displayName` search.
- WebSocket sessions per `document_id` for live edits.
- Broadcast edits to all collaborators in a session.
- Broadcast presence and cursor updates per session.
- Persist content and metadata in PostgreSQL.
- Support multiple backend replicas with NATS fan-out.

## Non-Functional Requirements
- Latency: propagate edits in under 100ms (p95).
- Scalability: stateless collaboration nodes with NATS.
- Reliability: no loss of last persisted state on disconnect.
- Usability: no authentication or setup friction.
- Data: schema changes use code-first migrations tracked in repo.

## Success Metrics
- p95 edit propagation latency <100ms in typical load.
- 99.9% document save success for edit events.
- Average time to start collaborating <30 seconds.

## Open Questions
- Which CRDT or OT model to use for rich-text operations?
- Persistence cadence: per edit vs. debounced batching.
- Rate limits or guardrails for public document IDs.
