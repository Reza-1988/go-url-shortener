-- 000001_init.down.sql

DROP INDEX IF EXISTS idx_urls_owner_id;
DROP INDEX IF EXISTS idx_urls_short_code_unique;

DROP TABLE IF EXISTS urls;
DROP TABLE IF EXISTS users;