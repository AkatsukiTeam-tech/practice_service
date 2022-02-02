package dto

import "time"

type Person struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
