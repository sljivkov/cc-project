package domain

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	SSN       string `json:"ssn"`
}
