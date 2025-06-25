package domain

type LoanApplication struct {
	Uuid string `json:"uuid" db:"uuid"`

	Value int32 `json:"value" db:"value"`

	Phone string `json:"phone" db:"phone"`

	IncomingOrganizationUuid string `json:"incoming_organization_uuid" db:"incoming_organization_uuid"`
}
