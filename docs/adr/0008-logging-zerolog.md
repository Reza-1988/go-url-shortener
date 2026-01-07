# ADR-0008: Structured logging with zerolog (+ request_id)
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
We need production-like logging for debugging and traceability.
Logs should be structured (machine-readable) and consistent across requests.

## Decision
We will use **zerolog** for structured JSON logging.

### Minimum logging fields
- `request_id`
- `method`, `path`
- `status`
- `latency_ms`
- `user_id` (when available)
- `error` details (when applicable)

## Alternatives considered
- **zap**: strong option but slightly heavier for setup
- standard `log` package: not structured, less suitable for production-like observability

## Consequences
- Implement request-id middleware and log middleware early.
- Log format will be JSON for easy filtering in containers/CI logs.
