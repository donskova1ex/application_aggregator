package domain

type LoanApplication struct {
	Value       int32  `json:"value" db:"value"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}
