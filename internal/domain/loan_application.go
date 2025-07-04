package domain

import "time"

type LoanApplication struct {
	Uuid                     string    `json:"uuid" db:"uuid"`
	Value                    int32     `json:"value" db:"value"`
	Phone                    string    `json:"phone" db:"phone"`
	IncomingOrganizationUuid string    `json:"incoming_organization_uuid" db:"incoming_organization_uuid"`
	IssueOrganizationUuid    string    `json:"issue_organization_uuid" db:"issue_organization_uuid"`
	CreatedAt                time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                time.Time `json:"updated_at" db:"updated_at"`
}
