# Runbook — Local Dev & Operations

This runbook explains how to run, test, and troubleshoot the service locally.

---

## Prerequisites
- Go (recommended: **1.22**)
- Docker + Docker Compose
- (Optional) Make

---

## Configuration

We use environment variables. **Do not commit real secrets.**

### Required env vars (example)

Create `.env` from `.env.example`:

- `APP_ENV=dev`
- `APP_PORT=8080`
- `DATABASE_URL=postgres://<user>:<pass>@db:5432/<db>?sslmode=disable`
- `JWT_SECRET=<random_secret>`
- `JWT_EXPIRES_IN_SECONDS=3600`

> The exact keys should match the implementation in `internal/config`.

---

## Run with Docker (recommended)

1) Create local env file:
```bash
cp .env.example .env
````

2) Start services:

```bash
docker compose up --build
```

3) Stop services:

```bash
docker compose down
```

> Notes:
>
> * `.env` must NOT be committed (only `.env.example` is tracked).
> * DB runs on a persistent Docker volume.

## Database migrations (golang-migrate)

- Migrations live in `migrations/` and must be applied before using the API.
- **Important:** `urls.short_code` must have a **unique index**; do not rely only on application logic.

### Primary (recommended): run migrations via Docker Compose

We use the official `migrate/migrate` container to avoid local installation differences.

**Up**
```bash
docker compose run --rm migrate up
````

**Down (one step)**

```bash
docker compose run --rm migrate down 1
```

**Down (all)**

```bash
docker compose run --rm migrate down -all
```

**Verify tables**

```bash
docker compose exec db psql -U postgres -d go_url_shortener -c "\dt"
```

### Optional (later): Make targets

We may add Make targets for convenience:

```bash
make migrate-up
make migrate-down
make migrate-status
```

---

## Run without Docker (optional)

If you want to run directly on your machine:

1) Start Postgres locally (or keep docker db only)
2) Set env vars (or use `.env`)
3) Run:

```bash
go run ./cmd/server
```

---

## Tests

### Unit tests
```bash
go test ./...
```

### Integration tests (recommended)

We aim to have at least one E2E flow:

- register → login → shorten → redirect → verify `click_count`
- admin endpoint access → `403` for normal user

Integration tests may require a running Postgres instance.

---

## Formatting & static checks

### gofmt
```bash
gofmt -w .
```

---

## CI

CI runs:

- `gofmt` check
- `go test ./...`

Workflow: `.github/workflows/ci.yml`

---

## Troubleshooting

### 1) “connection refused” to DB
- Ensure `docker compose up` is running
- Check `DATABASE_URL` host:
    - inside docker network it is usually `db`
    - outside docker it is usually `localhost`

### 2) Migrations fail
- Confirm `DATABASE_URL` is correct
- Confirm migrations are in correct order and named properly
- Ensure the DB is clean or use down/reset carefully

### 3) 401 Unauthorized
- Ensure `Authorization: Bearer <token>` header is set
- Ensure token is not expired
- Ensure `JWT_SECRET` matches what the server uses

### 4) Redirect returns 404
- `shortCode` not found **OR** URL disabled
- confirm the record exists in DB and `is_disabled=false`

---

## Operational notes (MVP)

- Redirect uses **302** (see ADR-0007)
- Logging is structured JSON (zerolog) with `request_id` (see ADR-0008)
- Passwords use Argon2id (ADR-0004)
