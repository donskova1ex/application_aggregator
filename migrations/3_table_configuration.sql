-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS configuration (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) unique,
    organization_uuid VARCHAR(36) REFERENCES organizations(uuid) ON DELETE CASCADE,
    new_client BOOLEAN DEFAULT FALSE,
    loan_available BOOLEAN DEFAULT FALSE,
    max_loan_amount DECIMAL(18, 2) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    is_avaliable BOOLEAN DEFAULT FALSE
    );
CREATE INDEX IF NOT EXISTS idx_config_org_uuid ON configuration(organization_uuid);
CREATE INDEX IF NOT EXISTS idx_config_is_avaliable ON configuration(is_avaliable);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_config_org_uuid;
DROP INDEX IF EXISTS idx_config_is_avaliable;
DROP TABLE IF EXISTS configuration;
-- +goose StatementEnd
