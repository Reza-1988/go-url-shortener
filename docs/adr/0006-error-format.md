# ADR-0006: Standard API error format and HTTP status codes
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
To keep clients/tests simple, API error responses must be consistent across endpoints.
The project requires standard HTTP status codes and unified error handling.

## Decision
All API errors will use this JSON structure:

```json
{
  "error": {
    "code": "STRING_CODE",
    "message": "Human readable message"
  }
}
```

## Status code policy (minimum)

- `400` Bad Request: validation errors, invalid URL, malformed input
- `401` Unauthorized: missing/invalid token
- `403` Forbidden: valid token but insufficient role/permissions
- `404` Not Found: short code not found / resource not found
- `409` Conflict: uniqueness conflicts (if exposed; otherwise internal retry)
- `500` Internal Server Error: unexpected errors

## Alternatives considered

- Ad-hoc error responses per endpoint: rejected due to inconsistency and harder testing

## Consequences

- We must implement a centralized Fiber error handler (or helper) to enforce this format.
- Services should return domain errors with codes mapped to HTTP status codes in handlers.