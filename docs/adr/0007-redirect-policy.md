# ADR-0007: Redirect policy (302 for MVP)
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
The redirect endpoint `GET /:shortCode` must redirect to the original URL.
We must choose 301 (permanent) vs 302 (temporary).

## Decision
For MVP, we will use **302 Temporary Redirect**.

## Alternatives considered
- **301 Permanent Redirect**: can be cached aggressively by clients/browsers and makes future changes harder

## Consequences
- 302 keeps flexibility if we need to change original URLs or add rules later.
- Redirect behavior must be documented in API docs and README.
