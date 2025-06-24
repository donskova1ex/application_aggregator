-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS configuration (
    id SERIAL PRIMARY KEY,
    organization_id SERIAL REFERENCES organizations(id),
    new_client BOOLEAN DEFAULT FALSE,
    loan_available BOOLEAN DEFAULT FALSE,
    max_loan_amount DECIMAL(18, 2) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    is_avaliable BOOLEAN DEFAULT FALSE,
);
CREATE INDEX IF NOT EXISTS idx_org_id ON configuration(organization_id);
CREATE INDEX IF NOT EXISTS idx_is_avaliable ON configuration(is_avaliable);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_org_id;
DROP INDEX IF EXISTS idx_is_avaliable;
DROP TABLE IF EXISTS configuration;
-- +goose StatementEnd
