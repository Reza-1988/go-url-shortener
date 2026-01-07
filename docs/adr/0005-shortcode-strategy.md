# ADR-0005: Short code strategy (Base62 random, length 7)
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
Short links must map `short_code -> original_url`.
Short codes must be unique and safe under concurrent requests.
We want a simple approach that scales and avoids predictability issues.

## Decision
We will generate short codes as:
- **Random Base62**
- **Fixed length: 7 characters**
- Uniqueness enforced by **DB unique index** on `urls.short_code`

### Collision handling
- On DB unique constraint violation during insert, retry code generation up to **5 times**.
- If still failing (extremely unlikely), return a 500 error with a clear error code.

## Alternatives considered
- **Sequential ID + Base62 encoding**: easier uniqueness, but predictable (enumeration risk) unless additional hardening is added
- **Hash(original_url)**: deterministic but may leak information; also collisions require careful handling

## Consequences
- We must add a unique index/constraint on `short_code`.
- Insert logic must handle duplicate errors and retry.
- Length 7 is a balance of usability and collision probability for MVP scale.
