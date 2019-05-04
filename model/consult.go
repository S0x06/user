package model

type Consult struct {
	Base
	ContactName `json:"contact_name"`
	CompanyName `json:"company_name"`
	Jobs        `json:"jobs"`
	Phone       `json:"phone"`
	Email       `json:"email"`
	Demand      `json:"demand"`
}
