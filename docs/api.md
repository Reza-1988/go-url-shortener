# API Contract — v0.1

## Base URLs & routing

- **API endpoints prefix:** `/api/v1`
- **Public redirect (no API prefix):** `/:shortCode`

All responses are **JSON** unless specified otherwise (the public redirect returns an HTTP redirect).

---

## Authentication

MVP uses **JWT access token**.

### Request header
- `Authorization: Bearer <token>`

### RBAC
- **Admin endpoints** require `role=admin` in JWT claims.

See:
- [ADR-0003](./adr/0003-auth-jwt-access-token.md)

---

## Standard error format

All API errors follow this shape:

```json
{
  "error": {
    "code": "STRING_CODE",
    "message": "Human readable message"
  }
}
```

### Status codes (minimum)

- `400` — invalid input / validation error
- `401` — missing or invalid token
- `403` — forbidden (not admin / not allowed)
- `404` — not found
- `409` — conflict (optional; for unique conflicts if exposed)
- `500` — internal error

See:
- ADR-0006

---

## Common validation rules

### URL validation (MVP)

- allowed schemes: `http`, `https`
- must be parseable and have a valid host
- max length: **2048** (or **4096** if needed)

Optional hardening (if implemented):
- block localhost/private IP targets to reduce SSRF risk

---

## Endpoints

### 1) Auth

#### POST `/api/v1/auth/register`

Create a new user.

**Request**
```json
{
  "email": "user@example.com",
  "password": "StrongPassword123!"
}
```

**Response (201)**
```json
{
  "id": "user_id",
  "email": "user@example.com",
  "role": "user"
}
```

**Errors**
- `400` — invalid email/password format
- `409` — email already exists (optional)

---

#### POST `/api/v1/auth/login`

Login and receive a JWT access token.

**Request**
```json
{
  "email": "user@example.com",
  "password": "StrongPassword123!"
}
```

**Response (200)**
```json
{
  "access_token": "JWT_TOKEN",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

**Errors**
- `400` — invalid input
- `401` — wrong credentials

---

### 2) User URLs (JWT required)

#### POST `/api/v1/urls`

Create a short URL for the authenticated user.

**Request**
```json
{
  "original_url": "https://example.com/landing?utm_source=telegram&utm_campaign=w1"
}
```

**Response (201)**
```json
{
  "id": "url_id",
  "original_url": "https://example.com/landing?utm_source=telegram&utm_campaign=w1",
  "short_code": "aB7xQp1",
  "short_url": "https://<your-domain>/aB7xQp1",
  "click_count": 0,
  "is_disabled": false,
  "created_at": "2026-01-07T17:30:00Z"
}
```

**Notes**
- `short_code` generation: **random Base62 length 7**, uniqueness enforced by **DB unique index**.  
  See: ADR-0005

**Errors**
- `400` — invalid URL
- `401` — missing/invalid token
- `500` — could not generate unique code after retries (rare)

---

#### GET `/api/v1/urls`

List URLs owned by the authenticated user.

**Response (200)**
```json
{
  "items": [
    {
      "id": "url_id",
      "original_url": "https://example.com/landing?utm_source=telegram&utm_campaign=w1",
      "short_code": "aB7xQp1",
      "short_url": "https://<your-domain>/aB7xQp1",
      "click_count": 12,
      "is_disabled": false,
      "created_at": "2026-01-07T17:30:00Z"
    }
  ]
}
```

**Pagination (optional but recommended)**
If implemented:

- Query: `?limit=20&offset=0`
- Response may include: `limit`, `offset`, `total`

**Errors**
- `401` — missing/invalid token

---

#### GET `/api/v1/urls/:id` (optional)

Get details for a URL owned by the authenticated user.

**Response (200)**  
Same shape as the **Create URL** response.

**Errors**
- `401` — missing/invalid token
- `403` — not owner
- `404` — not found

---

### 3) Admin endpoints (JWT + role=admin)

#### GET `/api/v1/admin/users`

List users.

**Response (200)**
```json
{
  "items": [
    {
      "id": "user_id",
      "email": "user@example.com",
      "role": "user",
      "created_at": "2026-01-07T17:00:00Z"
    }
  ]
}
```

**Errors**
- `401` — missing/invalid token
- `403` — not admin

---

#### GET `/api/v1/admin/urls`

List URLs across the system.

**Response (200)**
```json
{
  "items": [
    {
      "id": "url_id",
      "owner_id": "user_id",
      "original_url": "https://example.com",
      "short_code": "aB7xQp1",
      "click_count": 12,
      "is_disabled": false,
      "created_at": "2026-01-07T17:30:00Z"
    }
  ]
}
```

**Errors**
- `401` — missing/invalid token
- `403` — not admin

---

#### PATCH `/api/v1/admin/urls/:id/disable`

Disable a URL (minimum admin action in MVP).

**Request**
```json
{
  "is_disabled": true
}
```

**Response (200)**
```json
{
  "id": "url_id",
  "is_disabled": true
}
```

**Errors**
- `401` — missing/invalid token
- `403` — not admin
- `404` — not found

---

### 4) Public redirect (no auth)

#### GET `/:shortCode`

Redirect to original URL (MVP uses **302**).

**Behavior**
- If exists and not disabled:
    - increment `click_count` atomically
    - respond **302** with `Location: <original_url>`
- If not found or disabled:
    - respond **404**

See:
- ADR-0007

**Response**
- `302` redirect (no JSON body required)
