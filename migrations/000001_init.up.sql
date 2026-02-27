-- 000001_init.up.sql

-- USERS
CREATE TABLE IF NOT EXISTS users (
                                     id            BIGSERIAL PRIMARY KEY,
                                     email         TEXT NOT NULL UNIQUE,
                                     password_hash TEXT NOT NULL,
                                     role          TEXT NOT NULL DEFAULT 'user',
                                     created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

-- URLS
CREATE TABLE IF NOT EXISTS urls (
                                    id           BIGSERIAL PRIMARY KEY,
                                    owner_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    original_url TEXT NOT NULL,
    short_code   VARCHAR(8) NOT NULL,
    click_count  BIGINT NOT NULL DEFAULT 0,
    is_disabled  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

-- Uniqueness guarantee for short codes
CREATE UNIQUE INDEX IF NOT EXISTS idx_urls_short_code_unique ON urls(short_code);

-- Common query index
CREATE INDEX IF NOT EXISTS idx_urls_owner_id ON urls(owner_id);