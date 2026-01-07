# ADR-0003: Use JWT access tokens for authentication (MVP)
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
We need API authentication and role-based access (user vs admin).
The project allows JWT/PASETO/JWE/TLS-based approaches.
We want a simple, industry-standard approach within 2 weeks.

## Decision
We will use **JWT access tokens** for authentication in MVP.

### Token strategy (MVP)
- **Access token only** (no refresh token in MVP)
- Tokens are sent using: `Authorization: Bearer <token>`

### Claims (minimum)
- `sub` = user_id
- `role` = user/admin
- `exp` = expiration timestamp

## Alternatives considered
- **PASETO**: safer by design; less common in many ecosystems
- **JWE**: encrypted tokens; more complexity than needed for MVP
- **Sessions/cookies**: not ideal for pure API service in this bootcamp scope

## Consequences
- Protected endpoints require JWT middleware.
- RBAC is enforced using `role` claim checked in middleware for admin routes.
- Token expiration requires re-login once expired (acceptable for MVP).

