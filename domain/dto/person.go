package dto

import "time"

type Person struct {
	ID        int       `json:"id,omitempty"`
	FullName  string    `json:"fullName,omitempty"`
	Age       int       `json:"age,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdateAt  time.Time `json:"updateAt,omitempty"`
}
