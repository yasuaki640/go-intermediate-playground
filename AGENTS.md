# Repository Guidelines

## Project Structure & Module Organization
- `main.go` wires the HTTP server and database connection; keep startup logic here.
- `api/routers.go` declares routes and attaches middleware from `api/middlewares/`.
- HTTP handlers live in `controllers/`; use `controllers/testdata/` for fixture payloads and `controllers/services/` for mocks.
- Business rules belong in `services/`, while MySQL persistence stays inside `repositories/`.
- Declare shared DTOs in `models/` and domain-specific errors in `apperrors/`.
- SQL bootstrap scripts (`createTable.sql`, `insertData.sql`) support local development.

## Environment & Database Setup
- Copy `.env` to your local shell and export `ROOTUSER`, `ROOTPASS`, `DATABASE`, `USERNAME`, and `USERPASS`.
- Start MySQL with `docker compose up -d mysql`; the app expects it on `127.0.0.1:3306`.
- Apply `createTable.sql` then `insertData.sql` using your preferred MySQL client to mirror fixtures used in tests.

## Build, Run & Test Commands
- `go build ./...` verifies the module compiles and catches missing dependencies.
- `go run ./main.go` launches the API on `:8080`; watch logs for middleware output.
- `go test ./...` executes the controller, service, and repository suites; add `-run <Name>` for focused work.

## Coding Style & Naming Conventions
- Format every change with `gofmt` (tabs, 1TBS braces) and `goimports`; CI reviewers expect zero diffs.
- Exported types and functions use PascalCase; private helpers use camelCase.
- Keep handler filenames pluralized (e.g., `article_controller.go`) and suffix interfaces with `er` when they describe capabilities.

## Testing Guidelines
- Follow table-driven tests mirroring `services/service_test.go`; prefer `t.Helper()` for shared assertions.
- Name test functions `Test<Subject>` and mirror directory structure to the code under test.
- Provide deterministic fixtures; reuse `controllers/testdata/` or embed inline JSON to keep tests hermetic.

## Commit & Pull Request Guidelines
- Commit messages follow the existing short, imperative style (e.g., “add logging middleware to router”).
- One logical change per commit; include schema updates whenever SQL changes.
- Pull requests should summarise behaviour, list manual test steps, and link issues or tickets; attach screenshots for HTTP responses when relevant.
