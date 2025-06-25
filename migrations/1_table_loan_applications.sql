-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loan_applications (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) unique,
    phone_number VARCHAR(36) NOT NULL,
    value DECIMAL(18,2) NOT NULL CONSTRAINT chk_positive_value CHECK (value >= 0),
    incoming_organization_id INTEGER REFERENCES organizations(id),
    issue_organization_id INTEGER REFERENCES organizations(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL
);
CREATE INDEX IF NOT EXISTS idx_loan_app_value ON loan_applications(value);
CREATE INDEX IF NOT EXISTS idx_loan_app_incoming_org ON loan_applications(incoming_organization_id);
CREATE INDEX IF NOT EXISTS idx_loan_app_issue_org ON loan_applications(issue_organization_id);
CREATE INDEX IF NOT EXISTS idx_loan_app_created_at ON loan_applications(created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_loan_app_created_at;
DROP INDEX IF EXISTS idx_loan_app_issue_org;
DROP INDEX IF EXISTS idx_loan_app_incoming_org;
DROP INDEX IF EXISTS idx_loan_app_value;
DROP TABLE IF EXISTS loan_applications;
-- +goose StatementEnd
