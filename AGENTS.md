# Repository Guidelines

## Project Structure & Module Organization
This repository is currently minimal. At the time of writing it contains:

- `LICENSE`: project license.
- `.gitignore`: git ignore rules.
- `PRD.md` and `ARCHITECTURE.md`: product and system design.
- `TASKS.md`: milestone checklist for implementation.
- `docker-compose.yml`: local Postgres and NATS services.
- `cmd/`: Go service entrypoints (`document`, `collab`).
- `services/`: Go service packages and migrations.
- `frontend/`: React app with TipTap + Yjs.
- `shared/`: cross-service helpers (empty for now).

When adding source code, keep a clear separation between implementation and tests (for example, `src/` for source and `tests/` or `__tests__/` for test files). Place assets (fixtures, sample data) in an `assets/` or `testdata/` directory instead of mixing them into source folders.

## Build, Test, and Development Commands
Local infra is provisioned via Docker Compose. When more tooling is introduced, document it here with short, specific explanations, for example:

- `docker compose up -d`: start Postgres and NATS locally.
- `docker compose ps`: check local service status.
- `docker compose down -v`: stop and reset local data.
- `go test ./...`: build and test Go services.
- `go run ./cmd/document`: run Document Service.
- `go run ./cmd/collab`: run Collaboration Service.
- `cd frontend && npm install && npm run dev`: run the web app.
- `go run ./services/document/cmd/atlas > services/document/migrations/<timestamp>_init.sql`: generate a schema snapshot from Go models.

## Tech Stack (MVP)
- Backend services: Go.
- Frontend: React with TipTap + Yjs.
- Messaging: NATS.
- Database: PostgreSQL with Atlas code-first migrations.

## Coding Style & Naming Conventions
No code style is enforced yet. When the first language is introduced, establish:

- Indentation (for example, 2 spaces or 4 spaces).
- Naming patterns (for example, `camelCase` for variables, `PascalCase` for types).
- Formatting and linting tools (for example, `prettier`, `eslint`, `ruff`).

## Testing Guidelines
Testing framework is not defined. When tests exist, document:

- Test runner (for example, `pytest`, `jest`, `go test`).
- Test file naming conventions (for example, `*_test.py`, `*.spec.ts`).
- Minimum expectations for coverage or critical paths.

## Commit & Pull Request Guidelines
The Git history only includes a single “Initial commit”, so no commit message convention is established yet. If you adopt a standard (for example, Conventional Commits like `feat: add parser`), document it here.

For pull requests, include:

- A short description of the change and rationale.
- Links to related issues if applicable.
- Screenshots or CLI output when the change is user-facing.

## Configuration & Security Notes
If the project adds configuration files (for example, `.env`), document required keys and provide a safe example file (for example, `.env.example`). Avoid committing secrets to version control. Database schema changes must be done through code-first migrations and applied locally before opening a PR.
