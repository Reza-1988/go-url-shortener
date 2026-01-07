# go-url-shortener

A **production-like, monolithic URL shortener** built with **Go (Fiber)** and **PostgreSQL (GORM)**.  
It provides authenticated APIs for users/admins and a public redirect endpoint for visitors, with minimum analytics (click count).

---

## What this service does (quick overview)
Marketing URLs are often long and error-prone to share. This service:
- lets authenticated users create short links for long URLs
- redirects public visitors from `/:shortCode` to the original URL
- tracks minimum analytics (**click_count**)
- provides admin endpoints to manage users/URLs

### Roles
- **User (authenticated):** create short URLs, list own URLs, view basic stats
- **Visitor (public):** open a short URL and get redirected (no login required)
- **Admin (authenticated, role=admin):** manage users and URLs across the system

---

## Tech stack
- **Language:** Go
- **HTTP framework:** Fiber ([ADR-0001](docs/adr/0001-use-fiber.md))
- **Database:** PostgreSQL ([ADR-0002](docs/adr/0002-use-postgres-gorm.md))
- **ORM:** GORM ([ADR-0002](docs/adr/0002-use-postgres-gorm.md))
- **Migrations:** golang-migrate/migrate
- **Auth:** JWT access tokens (MVP) ([ADR-0003](docs/adr/0003-auth-jwt-access-token.md))
- **Password hashing:** Argon2id ([ADR-0004](docs/adr/0004-password-hashing-argon2.md))
- **Logging:** zerolog ([ADR-0008](docs/adr/0008-logging-zerolog.md))

---

## MVP product rules (locked decisions)
- **Redirect:** `302` by default ([ADR-0007](docs/adr/0007-redirect-policy.md))
- **Short code:** random Base62, length **7**, uniqueness enforced by DB unique index + retry on collision ([ADR-0005](docs/adr/0005-shortcode-strategy.md))
- **URL validation:** allow only `http/https`, must be parseable with a host (optional hardening: block localhost/private IPs)
- **Analytics:** minimum `click_count` (atomic increment during redirect)

---

## API summary (v0.1)
- Base path: `/api/v1`
- Public redirect: `/:shortCode` (no `/api/v1` prefix)

### Auth
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### User URLs (JWT required)
- `POST /api/v1/urls`
- `GET  /api/v1/urls`
- `GET  /api/v1/urls/:id` (optional)

### Admin (JWT + role=admin)
- `GET  /api/v1/admin/users`
- `GET  /api/v1/admin/urls`
- `PATCH /api/v1/admin/urls/:id/disable`

Full contract + examples: **[docs/api.md](docs/api.md)**

---

## Repository standards (Docs-as-Code)
This repository is the **source of truth**:
- Architecture overview: **[docs/overview.md](docs/overview.md)**
- API contract: **[docs/api.md](docs/api.md)**
- Runbook (run/test/debug): **[docs/runbook.md](docs/runbook.md)**
- Decisions (ADR): **[docs/adr/](docs/adr/)**
- Meeting notes: **[docs/meetings/](docs/meetings/)**

**Telegram policy:** Telegram is used only for quick coordination and posting links to PRs/Issues/ADRs/meeting notes.

---

## Quick start (development)
> Requires Docker + Docker Compose.

1) Create env file:
```bash
cp .env.example .env
````

2. Start services:

```bash
docker compose up --build
```

3. Run tests:

```bash
go test ./...
```

More details (migrations, troubleshooting): **[docs/runbook.md](docs/runbook.md)**

---

## Project structure (high level)

* `cmd/server`: application entrypoint
* `internal/http`: Fiber routes/handlers/middleware/DTOs + centralized error handling
* `internal/service`: business logic (use-cases)
* `internal/repository`: database access (Postgres/GORM)
* `migrations/`: SQL migrations

---

## Contribution workflow

* `main` is protected → PR required
* 1 approval required
* CI must pass before merge
* Keep PRs small and scoped; update docs/ADR when introducing decisions

---

## License

See [LICENSE](LICENSE)

