package models

import (
	"time"
)

type Person struct {
	ID        int       `json:"id" gorm:"primary_key; auto_increment; unique"`
	FullName  string    `json:"full_name" gorm:"size:255; not null"`
	Age       int       `json:"age" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt  time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP"`
}
