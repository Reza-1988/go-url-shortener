# ADR-0004: Use Argon2 for password hashing
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
Passwords must be stored securely. Plain-text passwords are forbidden.
We need a modern password hashing algorithm suitable for production-like systems.

## Decision
We will use **Argon2id** for password hashing and verification.

## Alternatives considered
- **bcrypt**: widely used and acceptable; generally slower to tune and less modern than Argon2id

## Consequences
- Only password hashes are stored in DB (`password_hash`).
- Password verification uses constant-time comparison.
- We must standardize parameters (time/memory/parallelism) and keep them configurable if needed.

