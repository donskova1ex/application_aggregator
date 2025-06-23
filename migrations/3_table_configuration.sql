-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS configuration (
    organization_id SERIAL REFERENCES organizations(id),
    new_client BOOLEAN DEFAULT FALSE,
    loan_available BOOLEAN DEFAULT FALSE,
    max_loan_amount DECIMAL(18, 2) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (organization_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS configuration;
-- +goose StatementEnd
