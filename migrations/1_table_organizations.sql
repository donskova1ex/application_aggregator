-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organizations (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) unique,
    name VARCHAR(64),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    UNIQUE (name));
    CREATE INDEX IF NOT EXISTS idx_org_uuid ON organizations(uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_org_uuid;
DROP TABLE IF EXISTS organizations;
-- +goose StatementEnd
