package model

import "time"

type User struct {
	UserID    string    `json:"-" db:"user_id, omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name, omitempty"`
	Role      string    `json:"role,omitempty" db:"role, omitempty"`
	Password  string    `json:"password,omitempty" db:"password, omitempty"`
	Email     string    `json:"email,omitempty" db:"email, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
}
