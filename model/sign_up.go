package model

type SignUp struct {
	FullName string `json:"fullName,omitempty" db:"full_name, omitempty"`
	Password string `json:"password,omitempty" db:"password, omitempty"`
	Email    string `json:"email,omitempty" db:"email, omitempty"`
}
