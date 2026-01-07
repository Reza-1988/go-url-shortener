# Architecture Overview — go-url-shortener

This project is a **monolithic, production-like URL shortener** built with Go (Fiber) + PostgreSQL (GORM).
It provides authenticated APIs for users/admins and a public redirect endpoint for visitors.

---

## Goals (what we must deliver)
- Authenticated users can **create** short URLs and **list** their own URLs.
- Public visitors can open `GET /:shortCode` and be **redirected** to the original URL.
- Minimum analytics: **click_count** increments on each redirect.
- Admin can **manage** (at least list + disable) users/URLs.
- Engineering quality: Docker (dev/prod), tests, docs-as-code, consistent error handling.

## Non-goals (MVP)
- Custom aliases, expiration, QR codes, detailed analytics (visits table), Redis cache, rate limiting, metrics dashboard.

---

## Roles (who uses the system)
- **User (authenticated):** creates short URLs and views own data.
- **Visitor (public):** opens short URLs; no login needed.
- **Admin (authenticated, role=admin):** manages users and URLs.

---

## Key product rules (locked decisions)
- Redirect: **302** (temporary)  
  See: [ADR-0007](./adr/0007-redirect-policy.md)
- Short code: **random Base62**, fixed length **7**, uniqueness enforced by DB unique index, retry on collision  
  See: [ADR-0005](./adr/0005-shortcode-strategy.md)
- Auth: **JWT access token (MVP)**, RBAC via role claim  
  See: [ADR-0003](./adr/0003-auth-jwt-access-token.md)
- Password hashing: **Argon2id**  
  See: [ADR-0004](./adr/0004-password-hashing-argon2.md)
- Logging: **zerolog**, structured JSON + request_id  
  See: [ADR-0008](./adr/0008-logging-zerolog.md)
- Error format: consistent JSON error structure across endpoints  
  See: [ADR-0006](./adr/0006-error-format.md)

---

## High-level request flow

### A) Create short URL (authenticated)
Client
→ Fiber router
→ Auth middleware (JWT)
→ Handler (HTTP parsing/validation)
→ Service (business rules, short code generation)
→ Repository (DB insert)
→ Response (short URL + metadata)

### B) Public redirect (no auth)
Visitor
→ `GET /:shortCode`
→ Handler
→ Service (lookup + check disabled)
→ Repository:
- fetch original_url
- atomic increment click_count
  → `302 Location: original_url`

---

## Code organization (clean layering)
We keep Fiber-specific code in the HTTP layer and business logic in services:

- **HTTP Layer (`internal/http/...`)**
    - Routing, handlers/controllers, DTOs, middleware, centralized error handler.
    - No business rules beyond basic request validation.

- **Service Layer (`internal/service/...`)**
    - Use-cases / business logic:
        - validate URL rules (service-level)
        - generate short code + retry policy
        - RBAC checks for admin endpoints
        - redirect behavior + analytics increment strategy

- **Repository Layer (`internal/repository/...`)**
    - DB access (Postgres/GORM).
    - Must implement atomic operations (e.g., click_count increment).

This separation keeps logic testable (unit tests can target services without Fiber).

---

## Persistence model (MVP)
### users
- id
- email (unique)
- password_hash
- role: user/admin
- created_at, updated_at

### urls
- id
- owner_id (FK -> users.id)
- original_url
- short_code (unique)
- click_count (default 0)
- is_disabled (default false)
- created_at, updated_at

**DB guarantees**
- Unique index on `urls.short_code`
- Unique index on `users.email`
- Atomic increment on click_count in redirect path

---

## Configuration & secrets
- Use environment variables for configuration (Docker-friendly).
- Never commit `.env` to the repo; commit `.env.example` only.
- JWT secret and DB credentials must come from env.

---

## Observability (logging)
- Structured logs with zerolog (JSON)
- Include `request_id` on every request
- Log method/path/status/latency and user_id when available

---

## Testing strategy
- Unit tests: service-level business rules (no DB, repo mocked).
- Integration tests: core E2E flows (auth → shorten → redirect → admin RBAC).

---

## Docs-as-Code workflow (source of truth)
- Architecture & docs: `docs/`
- Decisions: `docs/adr/`
- Meeting notes: `docs/meetings/`
- Telegram: links-only announcements (PR/Issue/ADR/meeting links)

See also:
- API contract: [docs/api.md](./api.md)
- Runbook: [docs/runbook.md](./runbook.md)

