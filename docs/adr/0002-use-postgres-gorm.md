# ADR-0002: Use PostgreSQL + GORM for persistence
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
We need persistent storage for users, URLs, and analytics (minimum click count).
The project recommends PostgreSQL and requires an ORM.

## Decision
We will use:
- **PostgreSQL** as the database
- **GORM** as the ORM

## Alternatives considered
- **MongoDB**: allowed, but relational constraints (unique short_code, ownership) are simpler in Postgres
- **Ent**: strong typed ORM but more setup overhead for the 2-week timeline
- **sqlc**: excellent for SQL-first, but requires more manual query/structure work (still acceptable, but team chose GORM)

## Consequences
- Schema will be defined via SQL migrations, while runtime access is via GORM models.
- We will enforce key integrity at DB level (e.g., unique index on `short_code`, unique email).
- Repository layer will wrap GORM and expose interfaces to services.

## Notes
- Minimum tables: `users`, `urls` (+ optional fields for admin/disable).

