package domain

type Organization struct {
	Uuid string `json:"uuid" db:"uuid"`
	Name string `json:"name" db:"name"`
}
