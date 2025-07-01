-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loan_applications (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) unique,
    phone VARCHAR(36) NOT NULL,
    value DECIMAL(18,2) NOT NULL CONSTRAINT chk_positive_value CHECK (value >= 1000),
    incoming_organization_uuid VARCHAR(36) REFERENCES organizations(uuid),
    issue_organization_uuid VARCHAR(36) REFERENCES organizations(uuid),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL
    );
CREATE INDEX IF NOT EXISTS idx_loan_app_value ON loan_applications(value);
CREATE INDEX IF NOT EXISTS idx_loan_app_incoming_org ON loan_applications(incoming_organization_uuid);
CREATE INDEX IF NOT EXISTS idx_loan_app_issue_org ON loan_applications(issue_organization_uuid);
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
